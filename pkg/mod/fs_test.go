package mod

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/joshcarp/sysl-printing/pkg/syslutil"
	"github.com/stretchr/testify/assert"
)

func TestNewFs(t *testing.T) {
	t.Parallel()

	_, backendFs := syslutil.WriteToMemOverlayFs("/")
	fs := NewFs(backendFs)
	assert.Equal(t, backendFs, fs.source)
}

func TestOpenLocalFile(t *testing.T) {
	t.Parallel()

	filename := "deps.sysl"
	_, memfs := syslutil.WriteToMemOverlayFs("../../tests/")
	fs := NewFs(memfs)
	f, err := fs.Open(filename)
	assert.Nil(t, err)
	assert.Equal(t, "deps.sysl", filepath.Base(f.Name()))
}

func TestOpenRemoteFile(t *testing.T) {
	t.Parallel()

	filename := "github.com/joshcarp/sysl-printing/demo/examples/Modules/deps.sysl"
	_, memfs := syslutil.WriteToMemOverlayFs("/")
	fs := NewFs(memfs)
	f, err := fs.Open(filename)
	assert.Nil(t, err)
	assert.Equal(t, "deps.sysl", filepath.Base(f.Name()))
}

func TestOpenRemoteFileFailed(t *testing.T) {
	t.Parallel()

	filename := "github.com/wrong/repo/deps.sysl"
	_, memfs := syslutil.WriteToMemOverlayFs("/")
	fs := NewFs(memfs)
	f, err := fs.Open(filename)
	assert.Nil(t, f)
	assert.Equal(t, fmt.Sprintf("%s not found", filename), err.Error())
}
