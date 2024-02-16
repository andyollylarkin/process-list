package listers

import (
	"gitlab.mindsw.io/migrate-core-libs/process-list/pkg"
	"gitlab.mindsw.io/migrate-core-libs/process-list/pkg/listers/internal"
	"gitlab.mindsw.io/migrate-core-libs/process-list/pkg/readers"
)

type LinuxLocalProcessLister struct{}

func NewLinuxLocalProcessLister() *LinuxLocalProcessLister {
	return &LinuxLocalProcessLister{}
}

func (l *LinuxLocalProcessLister) ListProcess() ([]pkg.Process, error) {
	localReader := readers.NewLocalDirReader()
	return internal.ParseLinux(localReader)
}
