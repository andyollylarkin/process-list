package listers

import (
	"github.com/andyollylarkin/process-list/pkg"
	"github.com/andyollylarkin/process-list/pkg/executors"
	"github.com/andyollylarkin/process-list/pkg/listers/internal"
)

type configWin struct {
	smbExePath string
	user       string
	pass       string
	host       string
	port       int
}

type WindowsSmbProcessLister struct {
	cfg configWin
}

func NewWindowsSmbProcessLister(smbExePath, user, pass, host string, port int) *WindowsSmbProcessLister {
	return &WindowsSmbProcessLister{
		cfg: configWin{
			smbExePath: smbExePath,
			user:       user,
			pass:       pass,
			host:       host,
			port:       port,
		},
	}
}

func (l *WindowsSmbProcessLister) ListProcess() ([]pkg.Process, error) {
	exec := executors.NewSmbExecutor(l.cfg.smbExePath, l.cfg.user, l.cfg.pass, l.cfg.host, l.cfg.port)

	return internal.ParseWindows(exec)
}
