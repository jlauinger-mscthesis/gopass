package fs

import (
	"context"
	"fmt"
	"os"

	"github.com/gopasspw/gopass/internal/backend"
	"github.com/gopasspw/gopass/internal/out"
	"github.com/gopasspw/gopass/pkg/fsutil"
)

const (
	name = "fs"
)

func init() {
	backend.RegisterStorage(backend.FS, name, &loader{})
}

type loader struct{}

// New implements backend.StorageLoader
func (l loader) New(ctx context.Context, path string) (backend.Storage, error) {
	if err := os.MkdirAll(path, 0700); err != nil {
		return nil, err
	}
	be := New(path)
	out.Debug(ctx, "Using Storage Backend: %s", be.String())
	return be, nil
}

func (l loader) Init(ctx context.Context, path string) (backend.Storage, error) {
	if err := os.MkdirAll(path, 0700); err != nil {
		return nil, err
	}
	return l.New(ctx, path)
}

func (l loader) Handles(path string) error {
	if fsutil.IsDir(path) {
		return nil
	}
	return fmt.Errorf("dir not found")
}

func (l loader) Priority() int {
	return 1
}
func (l loader) String() string {
	return name
}
