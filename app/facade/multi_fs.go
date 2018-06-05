package facade

import (
	"net/http"
	"sync"
)

// MultiFS allows to attach multiple file systems to application bundle. File systems are being searched for file
// in the order they are attached. MultiFS embeds sync.RWLock for protection from possible data races.
type MultiFS struct {
	sync.RWMutex

	attached []http.FileSystem
}

// Attach file system fs.
func (mfs *MultiFS) Attach(fs http.FileSystem) {
	if fs != nil {
		mfs.Lock()
		defer mfs.Unlock()

		// Create the slice if nit and add the file system
		if mfs.attached == nil {
			mfs.attached = make([]http.FileSystem, 0, 1)
		}

		mfs.attached = append(mfs.attached, fs)
	}
}

// Implement http.MultiFS interface. Open searches for the file with name on all file systems
// attached in order of they were added.
func (mfs *MultiFS) Open(name string) (http.File, error) {
	mfs.RLock()
	defer mfs.RUnlock()

	// Try to open file on attached file systems and return the first successful
	for _, fs := range mfs.attached {
		if f, err := fs.Open(name); err != nil || f != nil {
			return f, err
		}
	}

	return nil, nil
}
