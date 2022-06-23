package cmd

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseEnvVars(t *testing.T) {
	type args struct {
		serverHost string
		serverPort string
	}
	type data struct {
		title string
		args  args
		want  *Env
	}
	d := []data{
		{
			title: "happy flow with host and port",
			args:  args{serverHost: "0.0.0.0", serverPort: "8080"},
			want: &Env{
				ServerHost: "0.0.0.0",
				ServerPort: 8080,
			},
		},
		{
			title: "host empty string set in env",
			args:  args{serverPort: "8080"},
			want: &Env{
				ServerHost: "",
				ServerPort: 8080,
			},
		},
		{
			title: "invalid integer value of port",
			args:  args{serverHost: "0.0.0.0", serverPort: "abcd"},
			want: &Env{
				ServerHost: "0.0.0.0",
				ServerPort: 0,
			},
		},
	}

	for _, tt := range d {
		var e = &Env{}
		os.Setenv("SERVER_HOST", tt.args.serverHost)
		os.Setenv("SERVER_PORT", tt.args.serverPort)
		e.ParseEnvVars()
		os.Unsetenv("SERVER_HOST")
		os.Unsetenv("SERVER_PORT")
		if !assert.Equal(t, tt.want, e) {
			t.Fatalf("%s - %s", tt.title, "env struct mismatched")
		}
	}

	d = []data{
		{
			title: "port not set in env",
			args:  args{serverHost: "0.0.0.0"},
			want: &Env{
				ServerHost: "0.0.0.0",
				ServerPort: 0,
			},
		},
		{
			title: "host not set in env",
			args:  args{serverPort: "8080"},
			want: &Env{
				ServerHost: "",
				ServerPort: 8080,
			},
		},
		{
			title: "host and port both not set in env",
			args:  args{},
			want: &Env{
				ServerHost: "",
				ServerPort: 0,
			},
		},
	}

	for _, tt := range d {
		var e = &Env{}
		if tt.args.serverHost != "" {
			os.Setenv("SERVER_HOST", tt.args.serverHost)
		}
		if tt.args.serverPort != "" {
			os.Setenv("SERVER_PORT", tt.args.serverPort)
		}
		e.ParseEnvVars()
		os.Unsetenv("SERVER_HOST")
		os.Unsetenv("SERVER_PORT")
		if !assert.Equal(t, tt.want, e) {
			t.Fatalf("%s - %s", tt.title, "env struct mismatched")
		}
	}
}
