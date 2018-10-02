package cmd

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestServe(t *testing.T) {
	cases := []struct {
		command string
		want    string
	}{
		{command: "helloCobra serve", want: "serve called: optint: 0, optstr: default"},
		{command: "helloCobra serve --int 10", want: "serve called: optint: 10, optstr: default"},
		{command: "helloCobra serve --str test", want: "serve called: optint: 0, optstr: test"},
		{command: "helloCobra config", want: "config called"},
		{command: "helloCobra config create", want: "create called"},
	}

	for _, c := range cases {
		buf := new(bytes.Buffer)
		cmd := NewCmdRoot()
		cmd.SetOutput(buf)
		cmdArgs := strings.Split(c.command, " ")
		fmt.Printf("cmdArgs %+v\n", cmdArgs)
		cmd.SetArgs(cmdArgs[1:])
		cmd.Execute()

		get := buf.String()
		if c.want != get {
			t.Errorf("unexpected response: want:%+v, get:%+v", c.want, get)
		}
	}
}
