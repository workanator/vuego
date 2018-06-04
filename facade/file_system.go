package facade

import (
	"net/http"
	"sync"
)

// FileSystem allows for attach multiple file systems to application. File systems are being searched for file
// in the order they are attached. FileSystems embeds sync.RWLock for protection from possible data races.
type FileSystem struct {
	sync.RWMutex

	attached []http.FileSystem
}

// Attach file system fs.
func (fileSys *FileSystem) Attach(fs http.FileSystem) {
	if fs != nil {
		fileSys.Lock()
		defer fileSys.Unlock()

		// Create the slice if nit and add the file system
		if fileSys.attached == nil {
			fileSys.attached = make([]http.FileSystem, 0, 1)
		}

		fileSys.attached = append(fileSys.attached, fs)
	}
}

// Implement http.FileSystem interface. Open searches for the file with name on all file systems
// attached in order of they were added.
func (fileSys *FileSystem) Open(name string) (http.File, error) {
	fileSys.RLock()
	defer fileSys.RUnlock()

	// Try to open file on attached file systems and return the first successful
	for _, fs := range fileSys.attached {
		if f, err := fs.Open(name); err != nil || f != nil {
			return f, err
		}
	}

	return nil, nil
}
