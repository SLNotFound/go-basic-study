package main

import (
	"archive/tar"
	"os"
	"path/filepath"
)

type StreamArchive struct {
	Writer tar.Writer
}

func (s *StreamArchive) StreamFile(relPath string, fi os.FileInfo, data []byte) error {
	if fi.IsDir() {
		fh, err := tar.FileInfoHeader(fi, "")
		if err != nil {
			return err
		}
		fh.Name = relPath + "/"
		if err = s.Writer.WriteHeader(fh); err != nil {
			return err
		}
	} else {
		target := ""
		if fi.Mode()&os.ModeSymlink != 0 {
			target = string(data)
		}
		fh, err := tar.FileInfoHeader(fi, target)
		if err != nil {
			return err
		}
		fh.Name = filepath.Join(relPath, fi.Name())
		if err = s.Writer.WriteHeader(fh); err != nil {
			return err
		}
		if _, err = s.Writer.Write(data); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	var streamArchive StreamArchive
	var fi os.FileInfo
	var data []byte
	if err := streamArchive.StreamFile("F:\\AliDriver", fi, data); err != nil {
		panic(err)
	}
}
