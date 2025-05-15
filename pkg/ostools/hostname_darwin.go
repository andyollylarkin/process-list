//go:build darwin

package ostools

import (
	"os/exec"
	"strings"

	"github.com/andyollylarkin/process-list/pkg"
)

func Hostname(reader pkg.DirReader) (string, error) {
	out, err := exec.Command("hostname").Output()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(out)), nil
}
