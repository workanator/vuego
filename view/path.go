package view

import (
	"io/ioutil"
	"net/http"
)

// Built-in template paths.
const (
	TemplatePage Path = "html/page.html"
	TemplateView      = "html/view.html"
)

// Path to file in application bundle's file system with template content.
type Path string

// Test template is empty.
func (path Path) IsEmpty() bool {
	return string(path) == ""
}

// Load template from file system fs.
func (path Path) Load(fs http.FileSystem) ([]byte, error) {
	// Try open the file
	if file, err := fs.Open(string(path)); err != nil || file == nil {
		// File is not found or error happened
		return nil, err
	} else {
		// Read and return the content
		return ioutil.ReadAll(file)
	}
}
