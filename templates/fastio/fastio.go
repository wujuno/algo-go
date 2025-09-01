
package fastio

import (
	"bufio"
	"os"
)

// NewReader returns a buffered reader bound to stdin.
func NewReader() *bufio.Reader { return bufio.NewReader(os.Stdin) }

// NewWriter returns a buffered writer bound to stdout.
func NewWriter() *bufio.Writer { return bufio.NewWriter(os.Stdout) }
