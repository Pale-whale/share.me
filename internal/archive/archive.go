package archive

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var archives []string

func RemoveAll() {
	for _, arc := range archives {
		os.Remove(arc)
	}
}

func Tar(source, target string) error {
	tarfile, err := ioutil.TempFile("/tmp", "tar")
	if err != nil {
		return err
	}
	defer os.Remove(tarfile.Name())

	tarball := tar.NewWriter(tarfile)
	defer tarball.Close()

	info, err := os.Stat(source)
	if err != nil {
		return nil
	}

	var baseDir string
	if info.IsDir() {
		baseDir = filepath.Base(source)
	}

	err = filepath.Walk(source,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			header, err := tar.FileInfoHeader(info, info.Name())
			if err != nil {
				return err
			}

			if baseDir != "" {
				header.Name = filepath.Join(baseDir, strings.TrimPrefix(path, source))
			}

			if err := tarball.WriteHeader(header); err != nil {
				return err
			}

			if info.IsDir() {
				return nil
			}

			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()
			_, err = io.Copy(tarball, file)
			return err
		})

	if err != nil {
		return nil
	}

	filename := filepath.Base(source)
	target = filepath.Join(target, fmt.Sprintf("%s.tar.gz", filename))
	endfile, err := os.Create(target)
	if err != nil {
		return err
	}
	archives = append(archives, endfile.Name())

	archiver := gzip.NewWriter(endfile)
	archiver.Name = filename
	defer archiver.Close()

	_, err = io.Copy(archiver, tarfile)
	return err
}
