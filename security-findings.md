# Security Findings

Identified by automated security review on 2026-04-23. Fix sequentially in discrete PRs.

---

## Critical

### 1. Shell injection in `there.go:40`
`"command -v " + name` is passed to `sh -c`, allowing arbitrary command execution if the argument is externally controlled.
- **Fix:** Replace with `exec.LookPath(name)`.
- **Status:** Open

---

## Important

### 2. `MustCompile` panic in `parser/parser.go:136`
User-supplied CLI name is interpolated directly into a regex and passed to `regexp.MustCompile`. An invalid regex (e.g. a name containing `(`) crashes the process.
- **Fix:** Use `regexp.Compile` + `regexp.QuoteMeta`.
- **Status:** Open

### 3. `ne` operator logic bug in `compare/compare.go:220`
`ctx.Success` is set to a regex match result inside the `Ne` case, producing wrong exit codes for `ne` comparisons.
- **Fix:** Remove the side-effecting `ctx.Success` mutation from the `Ne` case.
- **Status:** Open

### 4. Missing `cmd.Wait()` in `command/command.go:37` and `parser/parser.go:50`
Zombie processes accumulate when `is` is invoked in a tight loop because `cmd.Wait()` is never called after `cmd.Start()`.
- **Fix:** Add `defer cmd.Wait()` after each successful `cmd.Start()`.
- **Status:** Open

### 5. Integer overflow in `age/age.go:36`
`value * unitMultiplier` can overflow on 32-bit systems for large inputs.
- **Fix:** Add a bounds check or cap after `strconv.Atoi`.
- **Status:** Fixed

### 6. `ParseInt` base 0 in `battery.go:39`
Base 0 accidentally accepts hex (`0x...`) and octal (`0...`) input from the user.
- **Fix:** Use base 10 (`strconv.ParseInt(r.Val, 10, strconv.IntSize)`).
- **Status:** Open

---

## Minor

### 7. `--debug` logs env var values
Debug output includes the actual value of environment variables, which can expose secrets in CI logs.
- **Fix:** Document that `--debug` should not be used with sensitive vars, or redact values in `VarCmd` debug output.
- **Status:** Open

### 8. `is known summary var` dumps entire environment
No filtering or redaction. Should be documented prominently in help text.
- **Fix:** Add warning to CLI help text; consider a `--include-pattern` filter flag.
- **Status:** Open

### 9. `is fso age` accepts arbitrary paths
Acts as a file-existence oracle in restricted shell environments via exit code.
- **Fix:** Document the behaviour; add a path allowlist if restricted deployment is intended.
- **Status:** Open

### 10. Silently discarded `io.ReadAll` errors in `parser/parser.go`
Errors from reading subprocess stdout/stderr are discarded with `_`, masking failures.
- **Fix:** Propagate errors or log them when debug mode is active.
- **Status:** Open
