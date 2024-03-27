package readers

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"strings"
	"time"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

type sshDirEntry struct {
	fi fs.FileInfo
}

// Name returns the name of the file (or subdirectory) described by the entry.
// This name is only the final element of the path (the base name), not the entire path.
// For example, Name would return "hello.go" not "home/gopher/hello.go".
func (e sshDirEntry) Name() string {
	return e.fi.Name()
}

// IsDir reports whether the entry describes a directory.
func (e sshDirEntry) IsDir() bool {
	return e.fi.IsDir()
}

// Type returns the type bits for the entry.
// The type bits are a subset of the usual FileMode bits, those returned by the FileMode.Type method.
func (e sshDirEntry) Type() fs.FileMode {
	return e.fi.Mode().Type()
}

// Info returns the FileInfo for the file or subdirectory described by the entry.
// The returned FileInfo may be from the time of the original directory read
// or from the time of the call to Info. If the file has been removed or renamed
// since the directory read, Info may return an error satisfying errors.Is(err, ErrNotExist).
// If the entry denotes a symbolic link, Info reports the information about the link itself,
// not the link's target.
func (e sshDirEntry) Info() (fs.FileInfo, error) {
	return e.fi, nil
}

type SshDirReader struct {
	client *sftp.Client
}

func NewSshDirReader(user, pass string, privateKey []byte, privateKeyPass []byte, privKeyPath string,
	host string, port int,
) (*SshDirReader, error) {
	sshClient, err := newSshClient(user, pass, privateKey, privateKeyPass, privKeyPath, host, port, 0)
	if err != nil {
		return nil, err
	}

	sftpClient, err := sftp.NewClient(sshClient)
	if err != nil {
		return nil, err
	}

	return &SshDirReader{client: sftpClient}, nil
}

func (r *SshDirReader) ReadDir(dirName string) ([]fs.DirEntry, error) {
	info, err := r.client.ReadDir(dirName)
	if err != nil {
		return nil, err
	}

	outDirs := make([]fs.DirEntry, 0)

	for _, i := range info {
		outDirs = append(outDirs, sshDirEntry{fi: i})
	}

	return outDirs, nil
}

func (r *SshDirReader) ReadLink(name string) (string, error) {
	return r.client.ReadLink(name)
}

func (r *SshDirReader) ReadFile(filePath string) (string, error) {
	file, err := r.client.Open(filePath)
	if err != nil {
		return "", err
	}

	data, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	return string(data), nil

}

func newSshClient(user string, pass string, privKey []byte, privKeyPass []byte, privKeyPath string,
	host string, port int, timeout int,
) (*ssh.Client, error) {
	var authMethod []ssh.AuthMethod

	var err error

	if host == "" {
		err = errors.New("host is not specified")

		return nil, err
	}

	if pass == "" && privKey == nil && privKeyPath == "" {
		return nil, errors.New("one of private key, password or path to private key file should be specified")
	}

	if privKeyPath != "" {
		privKey, _ = getPrivKey(privKeyPath)
	}

	// If key specified
	if len(privKey) > 0 {
		signer, err := privKeyHandle(privKey, privKeyPass)
		if err == nil {
			authMethod = append(authMethod, ssh.PublicKeys(signer))
		}
	}

	if pass != "" {
		authMethod = append(authMethod, ssh.Password(pass))
	}

	if len(authMethod) < 1 {
		return nil, errors.New("no authentication methods available")
	}

	config := &ssh.ClientConfig{
		User:            user,
		Auth:            authMethod,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         time.Duration(timeout) * time.Second,
	}

	config.SetDefaults()

	addr := fmt.Sprintf("%s:%d", host, port)

	client, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func getPrivKey(pk string) ([]byte, error) {
	key, err := os.ReadFile(pk)
	if err != nil {
		return nil, err
	}

	return key, nil
}

func privKeyHandle(pk []byte, passphrase []byte) (signer ssh.Signer, err error) {
	signer, err = ssh.ParsePrivateKey(pk)
	if err != nil && strings.Contains(err.Error(), "this private key is passphrase protected") {
		signer, err = ssh.ParsePrivateKeyWithPassphrase(pk, passphrase)
	}

	return
}
