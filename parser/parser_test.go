package parser_test

import (
	"testing"

	"github.com/oalders/is/parser"
	"github.com/oalders/is/types"
	"github.com/stretchr/testify/assert"
)

const (
	ssh  = "../testdata/bin/ssh"
	tmux = "../testdata/bin/tmux"
)

//nolint:lll
func TestCLIVersion(t *testing.T) {
	t.Parallel()
	ctx := &types.Context{}

	tests := [][]string{
		{"ansible", "2.14.2", "ansible [core 2.14.2]"},
		{
			"bash",
			"5.2.15",
			`GNU bash, version 5.2.15(1)-release (aarch64-apple-darwin22.1.0)
	Copyright (C) 2022 Free Software Foundation, Inc.
	License GPLv3+: GNU GPL version 3 or later <http://gnu.org/licenses/gpl.html>

	This is free software; you are free to change and redistribute it.
	There is NO WARRANTY, to the extent permitted by law.`,
		},
		{"bat", "0.23.0", "bat 0.23.0 (871abd2)"},
		{"csh", "6.21.00", "tcsh 6.21.00 (Astron) 2019-05-08 (x86_64-apple-darwin) options wide,nls,dl,bye,al,kan,sm,rh,color,filec"},
		{
			"curl", "7.88.1",
			`curl 7.88.1 (x86_64-apple-darwin22.0) libcurl/7.88.1 (SecureTransport) LibreSSL/3.3.6 zlib/1.2.11 nghttp2/1.51.0
Release-Date: 2023-02-20
Protocols: dict file ftp ftps gopher gophers http https imap imaps ldap ldaps mqtt pop3 pop3s rtsp smb smbs smtp smtps telnet tftp
Features: alt-svc AsynchDNS GSS-API HSTS HTTP2 HTTPS-proxy IPv6 Kerberos Largefile libz MultiSSL NTLM NTLM_WB SPNEGO SSL threadsafe UnixSockets`,
		},
		{"dig", "9.10.6", "DiG 9.10.6"},
		{"docker", "20.10.21", "version 20.10.21, build baeda1f"},
		{"fpp", "0.9.2", "fpp version 0.9.2"},
		{"fzf", "0.53.0", "0.53.0 (c4a9ccd)"},
		{"gcc", "14.0.3", "clang version 14.0.3 (clang-1403.0.22.14.1)"},
		{"gh", "2.30.0", "gh version 2.30.0 (2023-05-30)"},
		{"go", "1.20.4", "go version go1.20.4 darwin/amd64"},
		{"jq", "1.6", "jq-1.6"},
		{"less", "633", "less 633 (PCRE2 regular expressions)"},
		{"lua", "5.4.6", "Lua 5.4.6  Copyright (C) 1994-2023 Lua.org, PUC-Rio"},
		{"make", "3.81", `GNU Make 3.81
Copyright (C) 2006  Free Software Foundation, Inc.
This is free software; see the source for copying conditions.
There is NO warranty; not even for MERCHANTABILITY or FITNESS FOR A
PARTICULAR PURPOSE.

This program built for i386-apple-darwin11.3.0`},
		{"md5sum", "9.3", "md5sum (GNU coreutils) 9.3"},
		{"nvim", "v0.10.0-dev-2663+gc1c6c1ee1", `NVIM v0.10.0-dev-2663+gc1c6c1ee1
Build type: RelWithDebInfo
LuaJIT 2.1.1710088188
Run "nvim -V1 -v" for more info`},
		{"node", "v20.2.0", "v20.2.0"},
		{"npx", "9.6.6", "9.6.6"},
		{"ocaml", "5.1.0", "The OCaml toplevel, version 5.1.0"},
		{"opam", "2.1.5", "2.1.5"},
		{"openssl", "3.3.6", "LibreSSL 3.3.6"},
		{"openssl", "1.1.1f", "OpenSSL 1.1.1f  31 Mar 2020"},
		{
			"perl", "v5.36.0",
			`This is perl 5, version 36, subversion 0 (v5.36.0) built for darwin-2level`,
		},
		{"oh-my-posh", "16.9.1", "16.9.1"},
		// the trailing newline is in perltidy's output, so this test should preserve it
		{"perltidy", "v20230701", `This is perltidy, v20230701 

Copyright 2000-2023, Steve Hancock

Perltidy is free software and may be copied under the terms of the GNU
General Public License, which is included in the distribution files.

Complete documentation for perltidy can be found using 'man perltidy'
or on the internet at http://perltidy.sourceforge.net.
`},
		{"perldoc", "v3.2801", "v3.2801, under perl v5.040000 for darwin"},
		{
			"pihole", "v5.17.1",
			`  Pi-hole version is v5.17.1 (Latest: v5.17.1)
  AdminLTE version is v5.20.1 (Latest: v5.20.1)
  FTL version is v5.23 (Latest: v5.23)`,
		},
		{"plenv", "2.3.1-8-gd908472", "plenv 2.3.1-8-gd908472"},
		{"python", "3.11.3", "Python 3.11.3"},
		{"python3", "3.11.3", "Python 3.11.3"},
		{"ripgrep", "13.0.0", "ripgrep 13.0.0"},
		{"rustc", "1.73.0", "rustc 1.73.0 (cc66ad468 2023-10-03)"},
		{"screen", "4.08.00", "Screen version 4.08.00 (GNU) 05-Feb-20"},
		{"sh", "3.2.57", `GNU bash, version 3.2.57(1)-release (x86_64-apple-darwin22)
Copyright (C) 2007 Free Software Foundation, Inc.
`},
		{"sqlite3", "3.46.0", "3.46.0 2024-05-23 13:25:27 96c92aba00c8375bc32fafcdf12429c58bd8aabfcadab6683e35bbb9cdebf19e (64-bit)"},
		{"tar", "3.5.3", "bsdtar 3.5.3 - libarchive 3.5.3 zlib/1.2.11 liblzma/5.0.5 bz2lib/1.0.8"},
		{"tcsh", "6.21.00", "tcsh 6.21.00 (Astron) 2019-05-08 (x86_64-apple-darwin) options wide,nls,dl,bye,al,kan,sm,rh,color,filec"},
		{"trurl", "0.6", "trurl version 0.6 libcurl/7.88.1 [built-with 7.87.0]"},
		{"tmux", "3.3a", "tmux 3.3a"},
		{
			"tree", "v2.1.0",
			`tree v2.1.0 © 1996 - 2022 by Steve Baker, Thomas Moore, Francesc Rocher, Florian Sesser, Kyosuke Tokoro`,
		},
		{"typos", "1.22.7", "typos-cli 1.22.7"},
		{"ubi", "0.0.24", "ubi 0.0.24"},
		{
			"unzip", "6.00",
			`caution:  both -n and -o specified; ignoring -o
UnZip 6.00 of 20 April 2009, by Info-ZIP.  Maintained by C. Spieler.  Send
bug reports using http://www.info-zip.org/zip-bug.html; see README for details.`,
		},
		{
			"vim", "9.0",
			"VIM - Vi IMproved 9.0 (2022 Jun 28, compiled Apr 15 2023 04:26:46)",
		},
		{"zsh", "5.9", "zsh 5.9 (x86_64-apple-darwin22.0)"},
	}

	for _, test := range tests {
		assert.Equal(t, test[1], parser.CLIVersion(
			ctx, test[0], test[2],
		))
	}

	{
		ctx := &types.Context{Debug: true}
		o, err := (parser.CLIOutput(ctx, "../testdata/bin/bad-version"))
		assert.NoError(t, err)
		assert.Equal(t, "X3v", o)
		got := parser.CLIVersion(ctx, "../testdata/bin/bad-version", o)
		assert.Equal(t, "X3v", got)
	}
}

func TestCLIOutput(t *testing.T) {
	t.Parallel()
	{
		ctx := &types.Context{Debug: true}
		o, err := (parser.CLIOutput(ctx, ssh))
		assert.NoError(t, err)
		assert.NotEmpty(t, o)
	}
	{
		ctx := &types.Context{}
		o, err := (parser.CLIOutput(ctx, tmux))
		assert.NoError(t, err)
		assert.NotEmpty(t, o)
	}

	{
		ctx := &types.Context{}
		o, err := (parser.CLIOutput(ctx, "tmuxxx"))
		assert.Error(t, err)
		assert.Empty(t, o)
	}

	{
		ctx := &types.Context{Debug: true}
		o, err := (parser.CLIOutput(ctx, "../testdata/bin/bad-version"))
		assert.NoError(t, err)
		assert.Equal(t, "X3v", o)
	}
}
