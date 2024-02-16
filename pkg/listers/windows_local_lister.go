package listers

import "gitlab.mindsw.io/migrate-core-libs/process-list/pkg"

type WindowsLocalProcessLister struct{}

func NewWindowsLocalProcessLister() *WindowsLocalProcessLister {
	return &WindowsLocalProcessLister{}
}

func (l *WindowsLocalProcessLister) ListProcess() ([]pkg.Process, error) {
	panic("not implemented yet") // TODO: Implement
}
