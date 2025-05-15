//go:build linux

package ostools

import (
	"strings"

	"github.com/andyollylarkin/process-list/pkg"
)

func Hostname(reader pkg.DirReader) (string, error) {
	h, err := reader.ReadFile("/etc/hostname")
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(h), nil
}
