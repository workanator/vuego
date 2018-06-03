package server

import (
	"io/ioutil"
	"net/http"
)

func (server *Server) openFile(path string) (http.File, error) {
	// Search for the file in bundle file system
	if server.bundle.Fs != nil {
		if f, err := server.bundle.Fs.Open(path); err != nil || f != nil {
			return f, err
		}
	}

	// Search for the file in built-in file system
	return server.fs.Open(path)
}

func (server *Server) readFileContent(path string) ([]byte, error) {
	// Open the application template
	f, err := server.openFile(path)
	if err != nil {
		return nil, err
	}

	// Read the whole content of the template
	content, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	return content, nil
}
