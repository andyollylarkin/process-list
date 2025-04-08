package listers

import (
	"github.com/andyollylarkin/process-list/pkg"
	"github.com/andyollylarkin/process-list/pkg/listers/internal"
	"github.com/andyollylarkin/process-list/pkg/readers"
)

type config struct {
	user           string
	pass           string
	privateKey     []byte
	privateKeyPass []byte
	privKeyPath    string
	host           string
	port           int
}

type LinuxSshProcessLister struct {
	cfg config
}

func NewLinuxSshProcessLister(user string, pass string, privateKey []byte, privateKeyPass []byte,
	privKeyPath string, host string, port int,
) *LinuxSshProcessLister {
	cfg := config{
		user:           user,
		pass:           pass,
		privateKey:     privateKey,
		privateKeyPass: privateKeyPass,
		privKeyPath:    privKeyPath,
		host:           host,
		port:           port,
	}
	return &LinuxSshProcessLister{
		cfg: cfg,
	}
}

func (l *LinuxSshProcessLister) ListProcess() ([]pkg.Process, error) {
	sshReader, err := readers.NewSshDirReader(l.cfg.user, l.cfg.pass, l.cfg.privateKey,
		l.cfg.privateKeyPass, l.cfg.privKeyPath, l.cfg.host, l.cfg.port)
	if err != nil {
		return nil, err
	}

	return internal.ParseLinux(sshReader)
}
