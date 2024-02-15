package readers

import (
	"io/fs"

	"github.com/pkg/sftp"
)

type sshDirEntry struct {
	fi fs.FileInfo
}

// Name returns the name of the file (or subdirectory) described by the entry.
// This name is only the final element of the path (the base name), not the entire path.
// For example, Name would return "hello.go" not "home/gopher/hello.go".
func (e sshDirEntry) Name() string {
	return e.Name()
}

// IsDir reports whether the entry describes a directory.
func (e sshDirEntry) IsDir() bool {
	return e.IsDir()
}

// Type returns the type bits for the entry.
// The type bits are a subset of the usual FileMode bits, those returned by the FileMode.Type method.
func (e sshDirEntry) Type() fs.FileMode {
	return e.Type()
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

func NewSshDirReader(client *sftp.Client) *SshDirReader {
	return &SshDirReader{client: client}
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
