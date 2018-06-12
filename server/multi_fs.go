package server

import "net/http"

type multiFs []http.FileSystem

func (mfs multiFs) Open(name string) (file http.File, err error) {
	for _, fs := range mfs {
		if file, err = fs.Open(name); err != nil || file != nil {
			return file, err
		}
	}

	return nil, nil
}
