package cmd

import (
	"bytes"
	"github.com/mattn/go-shellwords"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// CmdTestCase describes a test case that works with releases.
type CmdTestCase struct {
	name      string
	cmd       string
	golden    string
	wantError bool
}

func RunTestCmd(t *testing.T, tests []CmdTestCase) {
	t.Helper()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer resetEnv()()

			t.Log("running cmd: ", tt.cmd)
			_, output, err := executeActionCommandC(tt.cmd)
			if (err != nil) != tt.wantError {
				t.Errorf("expected error, got '%v'", err)
			}
			if tt.golden != "" {
				abs, _ := filepath.Abs(tt.golden)
				slash := filepath.FromSlash(abs)
				AssertGoldenString(t, output, slash)
			}
		})
	}
}

func executeActionCommandC(cmd string) (*cobra.Command, string, error) {
	args, err := shellwords.Parse(cmd)
	if err != nil {
		return nil, "", err
	}

	buf := new(bytes.Buffer)
	command := NewRootCmd(buf)

	command.SetOut(buf)
	command.SetArgs(args)

	c, err := command.ExecuteC()

	return c, buf.String(), err
}

func resetEnv() func() {
	origEnv := os.Environ()
	return func() {
		os.Clearenv()
		for _, pair := range origEnv {
			kv := strings.SplitN(pair, "=", 2)
			os.Setenv(kv[0], kv[1])
		}
	}
}
