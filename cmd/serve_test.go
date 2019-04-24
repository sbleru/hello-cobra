package cmd

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/mattn/go-shellwords"
)

func TestServe(t *testing.T) {
	cases := []struct {
		command string
		want    string
		errWant string
	}{
		{command: "helloCobra serve", want: "", errWant: "Parameter error: Optstr is required"},
		{command: "helloCobra serve --int 10", want: "", errWant: "Parameter error: Optstr is required"},
		{command: "helloCobra config", want: "config called"},
		{command: "helloCobra config create", want: "create called"},
		{command: "helloCobra serve --str \"test1 test2\"", want: "serve called: optint: 0, optstr: test1 test2"},
	}

	for _, c := range cases {
		buf := new(bytes.Buffer)
		cmd := NewCmdRoot()
		cmd.SetOutput(buf)
		cmdArgs, err := shellwords.Parse(c.command)
		if err != nil {
			t.Fatalf("args parse error: %+v\n", err)
		}
		fmt.Printf("cmdArgs %+v\n", cmdArgs)
		cmd.SetArgs(cmdArgs[1:])
		if err := cmd.Execute(); err != nil {
			if c.errWant != err.Error() {
				t.Errorf("unexpected error response: errWant:%+v, get:%+v", c.errWant, err.Error())
			}
		}

		get := buf.String()
		if c.want != get {
			t.Errorf("unexpected response: want:%+v, get:%+v", c.want, get)
		}
	}
}
