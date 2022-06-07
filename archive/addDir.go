package main

import (
	"archive/tar"
	"bytes"
	"fmt"
	"os"
	"path"
)

func addDir(writer *tar.Writer, dirPath string) error {
	dir, err := os.Open(dirPath)
	if err != nil {
		return err
	}
	defer dir.Close()
	fi, err := dir.Stat()
	if err != nil {
		return err
	}
	//fmt.Println(fi.Name())
	header, err := tar.FileInfoHeader(fi, "")
	header.Name = dirPath
	fmt.Println(header.Name)

	err = writer.WriteHeader(header)
	if err != nil {
		return err
	}

	fis, err := dir.Readdir(0)
	if err != nil {
		return err
	}
	for _, fi := range fis {
		if fi.IsDir() {
			err = addDir(writer, path.Join(dirPath, fi.Name()))
		} else {
			fmt.Println("fi is not dir!")
		}
		if err != nil {
			return err
		}

	}
	return nil
}

func main() {
	var buf bytes.Buffer
	tw := tar.NewWriter(&buf)
	err := addDir(tw, "F:\\AudioDump")
	if err != nil {
		panic(err)
	}
}
