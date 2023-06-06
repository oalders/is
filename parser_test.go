package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCLIVersion(t *testing.T) {
	ctx := &Context{}

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
		{
			"curl", "7.88.1",
			`curl 7.88.1 (x86_64-apple-darwin22.0) libcurl/7.88.1 (SecureTransport) LibreSSL/3.3.6 zlib/1.2.11 nghttp2/1.51.0
Release-Date: 2023-02-20
Protocols: dict file ftp ftps gopher gophers http https imap imaps ldap ldaps mqtt pop3 pop3s rtsp smb smbs smtp smtps telnet tftp
Features: alt-svc AsynchDNS GSS-API HSTS HTTP2 HTTPS-proxy IPv6 Kerberos Largefile libz MultiSSL NTLM NTLM_WB SPNEGO SSL threadsafe UnixSockets`,
		},
		{"docker", "20.10.21", "version 20.10.21, build baeda1f"},
		{"gcc", "14.0.3", "clang version 14.0.3 (clang-1403.0.22.14.1)"},
		{"gh", "2.30.0", "gh version 2.30.0 (2023-05-30)"},
		{"go", "1.20.4", "go version go1.20.4 darwin/amd64"},
		{"jq", "1.6", "jq-1.6"},
		{"less", "633", "less 633 (PCRE2 regular expressions)"},
		{"lua", "5.4.6", "Lua 5.4.6  Copyright (C) 1994-2023 Lua.org, PUC-Rio"},
		{"make", "3.81", "GNU Make 3.81"},
		{"md5sum", "9.3", "md5sum (GNU coreutils) 9.3"},
		{"node", "v20.2.0", "v20.2.0"},
		{"npx", "9.6.6", "9.6.6"},
		{
			"perl", "v5.36.0",
			`This is perl 5, version 36, subversion 0 (v5.36.0) built for darwin-2level`,
		},
		{"oh-my-posh", "16.9.1", "16.9.1"},
		{"plenv", "2.3.1-8-gd908472", "plenv 2.3.1-8-gd908472"},
		{"python", "3.11.3", "Python 3.11.3"},
		{"python3", "3.11.3", "Python 3.11.3"},
		{"ripgrep", "13.0.0", "ripgrep 13.0.0"},
		{"tar", "3.5.3", "bsdtar 3.5.3 - libarchive 3.5.3 zlib/1.2.11 liblzma/5.0.5 bz2lib/1.0.8"},
		{"trurl", "0.6", "trurl version 0.6 libcurl/7.88.1 [built-with 7.87.0]"},
		{"tmux", "3.3a", "tmux 3.3a"},
		{
			"tree", "v2.1.0",
			`tree v2.1.0 © 1996 - 2022 by Steve Baker, Thomas Moore, Francesc Rocher, Florian Sesser, Kyosuke Tokoro`,
		},
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
		assert.Equal(t, test[1], cliVersion(
			ctx, test[0], test[2],
		))
	}
}

func TestCLIOutput(t *testing.T) {
	ctx := &Context{}
	{
		o, err := (cliOutput(ctx, "ssh"))
		assert.NoError(t, err)
		assert.NotEmpty(t, o)
	}
	{
		o, err := (cliOutput(ctx, "tmux"))
		assert.NoError(t, err)
		assert.NotEmpty(t, o)
	}

	{
		o, err := (cliOutput(ctx, "tmuxxx"))
		assert.Error(t, err)
		assert.Empty(t, o)
	}
}
