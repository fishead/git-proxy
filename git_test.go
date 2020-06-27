package main

import "testing"

func Test_getProtocol(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "ssh://user@host.xz:22/path/to/repo.git/",
			args: args{s: "ssh://user@host.xz:22/path/to/repo.git/"},
			want: ProtocolSSH,
		},
		{
			name: "ssh://host.xz:22/path/to/repo.git/",
			args: args{s: "ssh://host.xz:22/path/to/repo.git/"},
			want: ProtocolSSH,
		},
		{
			name: "ssh://user@host.xz/path/to/repo.git/",
			args: args{s: "ssh://user@host.xz/path/to/repo.git/"},
			want: ProtocolSSH,
		},
		{
			name: "ssh://host.xz/path/to/repo.git/",
			args: args{s: "ssh://host.xz/path/to/repo.git/"},
			want: ProtocolSSH,
		},
		{
			name: "user@host.xz:path/to/repo.git/",
			args: args{s: "user@host.xz:path/to/repo.git/"},
			want: ProtocolSSH,
		},
		{
			name: "host.xz:path/to/repo.git/",
			args: args{s: "host.xz:path/to/repo.git/"},
			want: ProtocolSSH,
		},
		{
			name: "ssh://user@host.xz:22/~user/path/to/repo.git/",
			args: args{s: "ssh://user@host.xz:22/~user/path/to/repo.git/"},
			want: ProtocolSSH,
		},
		{
			name: "ssh://host.xz:22/~user/path/to/repo.git/",
			args: args{s: "ssh://host.xz:22/~user/path/to/repo.git/"},
			want: ProtocolSSH,
		},
		{
			name: "ssh://user@host.xz/~user/path/to/repo.git/",
			args: args{s: "ssh://user@host.xz/~user/path/to/repo.git/"},
			want: ProtocolSSH,
		},
		{
			name: "ssh://user@host.xz:22/~/path/to/repo.git/",
			args: args{s: "ssh://user@host.xz:22/~/path/to/repo.git/"},
			want: ProtocolSSH,
		},
		{
			name: "ssh://host.xz/~user/path/to/repo.git/",
			args: args{s: "ssh://host.xz/~user/path/to/repo.git/"},
			want: ProtocolSSH,
		},
		{
			name: "ssh://host.xz:22/~/path/to/repo.git/",
			args: args{s: "ssh://host.xz:22/~/path/to/repo.git/"},
			want: ProtocolSSH,
		},
		{
			name: "ssh://user@host.xz/~/path/to/repo.git/",
			args: args{s: "ssh://user@host.xz/~/path/to/repo.git/"},
			want: ProtocolSSH,
		},
		{
			name: "user@host.xz:22/~user/path/to/repo.git/",
			args: args{s: "user@host.xz:22/~user/path/to/repo.git/"},
			want: ProtocolSSH,
		},
		{
			name: "host.xz:22/~user/path/to/repo.git/",
			args: args{s: "host.xz:22/~user/path/to/repo.git/"},
			want: ProtocolSSH,
		},
		{
			name: "user@host.xz/~user/path/to/repo.git/",
			args: args{s: "user@host.xz/~user/path/to/repo.git/"},
			want: ProtocolSSH,
		},
		{
			name: "user@host.xz:22/~/path/to/repo.git/",
			args: args{s: "user@host.xz:22/~/path/to/repo.git/"},
			want: ProtocolSSH,
		},
		{
			name: "host.xz/~user/path/to/repo.git/",
			args: args{s: "host.xz/~user/path/to/repo.git/"},
			want: ProtocolSSH,
		},
		{
			name: "host.xz:22/~/path/to/repo.git/",
			args: args{s: "host.xz:22/~/path/to/repo.git/"},
			want: ProtocolSSH,
		},
		{
			name: "user@host.xz/~/path/to/repo.git/",
			args: args{s: "user@host.xz/~/path/to/repo.git/"},
			want: ProtocolSSH,
		},
		{
			name: "git://host.xz:22/path/to/repo.git/",
			args: args{s: "git://host.xz:22/path/to/repo.git/"},
			want: ProtocolGit,
		},
		{
			name: "git://host.xz/path/to/repo.git/",
			args: args{s: "git://host.xz/path/to/repo.git/"},
			want: ProtocolGit,
		},
		{
			name: "git://host.xz:22/~user/path/to/repo.git/",
			args: args{s: "git://host.xz:22/~user/path/to/repo.git/"},
			want: ProtocolGit,
		},
		{
			name: "git://host.xz/~user/path/to/repo.git/",
			args: args{s: "git://host.xz/~user/path/to/repo.git/"},
			want: ProtocolGit,
		},
		{
			name: "git://host.xz:22/~/path/to/repo.git/",
			args: args{s: "git://host.xz:22/~/path/to/repo.git/"},
			want: ProtocolGit,
		},
		{
			name: "http://host.xz:80/path/to/repo.git/",
			args: args{s: "http://host.xz:80/path/to/repo.git/"},
			want: ProtocolHTTP,
		},
		{
			name: "http://host.xz/path/to/repo.git/",
			args: args{s: "http://host.xz/path/to/repo.git/"},
			want: ProtocolHTTP,
		},
		{
			name: "https://host.xz:443/path/to/repo.git/",
			args: args{s: "https://host.xz:443/path/to/repo.git/"},
			want: ProtocolHTTPS,
		},
		{
			name: "https://host.xz/path/to/repo.git/",
			args: args{s: "https://host.xz/path/to/repo.git/"},
			want: ProtocolHTTPS,
		},
		{
			name: "ftp://host.xz:20/path/to/repo.git/",
			args: args{s: "ftp://host.xz:20/path/to/repo.git/"},
			want: ProtocolFTP,
		},
		{
			name: "ftp://host.x/path/to/repo.git/",
			args: args{s: "ftp://host.x/path/to/repo.git/"},
			want: ProtocolFTP,
		},
		{
			name: "ftps://host.xz:989/path/to/repo.git/",
			args: args{s: "ftps://host.xz:989/path/to/repo.git/"},
			want: ProtocolFTPS,
		},
		{
			name: "ftps://host.xz/path/to/repo.git/",
			args: args{s: "ftps://host.xz/path/to/repo.git/"},
			want: ProtocolFTPS,
		},
		{
			name: "/path/to/repo.git/",
			args: args{s: "/path/to/repo.git/"},
			want: ProtocolFile,
		},
		{
			name: "file:///path/to/repo.git/",
			args: args{s: "file:///path/to/repo.git/"},
			want: ProtocolFile,
		},
		{
			name: "./path/to/repo.git/",
			args: args{s: "./path/to/repo.git/"},
			want: ProtocolFile,
		},
		{
			name: "file://./path/to/repo.git/",
			args: args{s: "file://./path/to/repo.git/"},
			want: ProtocolFile,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getProtocol(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("getProtocol() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getProtocol() = %v, want %v", got, tt.want)
			}
		})
	}
}
