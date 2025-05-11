# Changes

- Replace some debugging output via:
    - Add "is known summary battery"
    - Add "is known summary battery --json"
    - Add "is known summary os"
    - Add "is known summary os --json"
- Add "is there [binary-name] --verbose"
- Add "is there [binary-name] --verbose --json"
- Wrap "which -a"
    - Add "is there [binary-name] --all"
    - Add "is there [binary-name] --all --json"

## 0.7.0 - 2025-04-18

- Add "is battery" and "is known battery" subcommands
- Add version parsing for golangci-lint
- Add version parsing for gopls
- Suppress "is user sudoer" error message when "sudo" is not installed

## 0.6.1 - 2025-01-27

- Allow for empty string comparisons via "eq" and "ne" in "is var". e.g.
  - is var FOO eq ""
  - is var FOO ne ""

## 0.6.0 - 2025-01-20

- Add var subcommand
- Add sequoia to macos code names

## 0.5.5 - 2024-10-29

- Add version parsing for hugo

## 0.5.4 - 2024-08-16

- Run "command -v" via "sh -c" (GH#37) (Olaf Alders)

## 0.5.3 - 2024-06-22

- Parse versions for: dig, perldoc, fpp, fzf, screen, sqlite3 and typos (GH#35)
  (Olaf Alders)
- Improve completion documentation

## 0.5.2 - 2024-06-18

- Add some simple command line completion (GH#20) (Olaf Alders)

## 0.5.1 - 2024-06-18

- Add NetBSD and FreeBSD to build targets (GH#33) (Jason A. Crome)

## 0.5.0 - 2024-06-08

- Add fso subcommand (GH#32) (Olaf Alders)
- Fix neovim nightly version parsing (GH#28)

## 0.4.3 - 2024-03-05

- Add OCaml toolchain (GH#27) (Rawley)

## 0.4.2 - 2023-11-27

- Fix cli version parsing for "rustc"
- Add "sonoma" to macOS codenames

## 0.4.1 - 2023-09-28

- Ensure stringy comparison of "in" via optimistic compare

## 0.4.0 - 2023-09-27

- Add "in" for matching on items in a comma-delimited list

## 0.3.0 - 2023-09-24

- Add --major, --minor and --patch version segment constraints

## 0.2.0 - 2023-09-15

- Add "is command output"

## 0.1.2 - 2023-09-08

- Improve docs
- Re-organize internals

## 0.1.1 - 2023-08-16

- Add better fallback version parsing
- Add "is user sudoer"

## 0.1.0 - 2023-07-05

- Add h|hour|hours to cli age units

## 0.0.9 - 2023-07-04

- Add "cli age"

## 0.0.8 - 2023-06-23

- Parse openssl versions
- Add a linux release for arm 7

## 0.0.7 - 2023-06-15

- Remove "os arch" which was badly named and not necessarily correct

## 0.0.6 - 2023-06-14

- "command-version" is now "cli version"
- Add more os attributes and attribute comparisons

## 0.0.5 - 2023-05-29

- Silence error output when a command is not found

## 0.0.4 - 2023-05-26

- Add a --version flag

## 0.0.3 - 2023-05-20

- Tweak goreleaser config

## 0.0.2 - 2023-05-20

- Test goreleaser GitHub action

## 0.0.1 - 2023-05-20

- First release upon an unsuspecting world.
