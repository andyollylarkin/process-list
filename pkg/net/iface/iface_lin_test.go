//go:build linux

package iface

import (
	"bytes"
	"testing"

	"io/fs"

	"github.com/andyollylarkin/process-list/pkg"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

const routeFileContent = `Iface   Destination     Gateway         Flags   RefCnt  Use     Metric  Mask            MTU     Window  IRTT
ens4    00000000        FEBFA8C0        0003    0       0       101     00000000        0       0       0
ens3    0064A8C0        00000000        0001    0       0       100     00FFFFFF        0       0       0
ens4    00BFA8C0        00000000        0001    0       0       101     00FFFFFF        0       0       0`

type fakeFile struct {
	buf *bytes.Buffer
}

func (f *fakeFile) Stat() (fs.FileInfo, error) {
	panic("TODO: Implement")
}

func (f *fakeFile) Read(p0 []byte) (int, error) {
	return f.buf.Read(p0)
}

func (f *fakeFile) Close() error {
	return nil
}

func TestNetSettingsObserver_RoutingTable(t *testing.T) {
	reader := pkg.NewMockDirReader(gomock.NewController(t))

	fakeFile := fakeFile{bytes.NewBuffer([]byte(routeFileContent))}

	reader.EXPECT().Open("/proc/net/route").Return(fakeFile, nil).Times(1)

	nso := NewNetSettingsObserver(reader)
	routes, err := nso.RoutingTable()

	require.NoError(t, err)

	require.Len(t, routes, 3)
}
