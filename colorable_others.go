// +build !windows

package colorable

import (
	"io"
	"os"
)

func NewColorableStdout() io.Writer {
	return os.Stdout
}
