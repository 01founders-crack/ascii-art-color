package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func TestPrintColored(t *testing.T) {
	tests := []struct {
		name       string
		text       string
		textColor  string
		wantOutput string
	}{
		{
			name:       "Print in red",
			text:       "Hello",
			textColor:  "Red",
			wantOutput: "\033[31mHello\033[0m",
		},
		{
			name:       "Print in blue",
			text:       "World",
			textColor:  "Blue",
			wantOutput: "\033[34mWorld\033[0m",
		},
		// Add more test cases as needed
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			os.Args = []string{"testProgram", "--color=" + test.textColor, test.text}
			output := captureOutput(func() {
				main()
			})
			if output != test.wantOutput+"\n" {
				t.Errorf("Expected output: %s, got: %s", test.wantOutput, output)
			}
		})
	}
}

func captureOutput(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old

	var buf strings.Builder
	io.Copy(&buf, r)

	return buf.String()
}
