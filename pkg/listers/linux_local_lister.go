package listers

import (
	"github.com/andyollylarkin/process-list/pkg"
	"github.com/andyollylarkin/process-list/pkg/listers/internal"
	"github.com/andyollylarkin/process-list/pkg/readers"
)

type LinuxLocalProcessLister struct{}

func NewLinuxLocalProcessLister() *LinuxLocalProcessLister {
	return &LinuxLocalProcessLister{}
}

func (l *LinuxLocalProcessLister) ListProcess(fn pkg.DoneLookupFunc) ([]pkg.Process, error) {
	localReader := readers.NewLocalDirReader()
	return internal.ParseLinux(localReader, fn)
}
