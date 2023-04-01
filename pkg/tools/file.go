package tools

import (
	"archive/tar"
	"bufio"
	"compress/gzip"
	"fmt"
	"github.com/jedib0t/go-pretty/v6/progress"
	"github.com/sirupsen/logrus"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

type writeCounter struct {
	track *progress.Tracker
}

func (w *writeCounter) Write(p []byte) (int, error) {
	n := len(p)
	w.track.Increment(int64(n))
	if w.track.Value() >= w.track.Total {
		w.track.MarkAsDone()
	}
	return n, nil
}

func WriteBufferedFile(filename string, src io.ReadCloser, size int64, track *progress.Tracker) {
	defer src.Close()
	file, err := os.Create(filename)
	if err != nil {
		logrus.Fatalf("create file %s error: %+v", filename, err)
	}
	fileWriter := bufio.NewWriter(file)
	defer file.Close()

	wc := &writeCounter{
		track: track,
	}
	_, err = io.Copy(fileWriter, io.TeeReader(src, wc))
	if err != nil {
		logrus.Fatalf("write file %s error: %+v", filename, err)
	}
	fileWriter.Flush()
}

func WriteFile(filename string, content []byte) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			logrus.Fatalf("close file %s error: %+v", filename, err)
		}
	}(file)
	if err != nil {
		logrus.Fatalf("create file %s error: %+v", filename, err)
	}

	_, err = file.Write(content)
	if err != nil {
		logrus.Fatalf("write file %s error: %+v", filename, err)
	}
}

func IsPathExist(path string) (res bool) {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func MkdirPath(path string) {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		logrus.Fatalf("crete path %s error: %+v", path, err)
	}
}

func RemovePath(path string) error {
	err := os.RemoveAll(path)
	return err
}

func TarDir(srcDir, destFile string) {
	if IsPathExist(destFile) {
		logrus.Debugf("delete target file: %s", destFile)
		err := RemovePath(destFile)
		if err != nil {
			logrus.Fatalf("delete target file %s error: %+v", destFile, err)
		}
	}
	fw, err := os.Create(fmt.Sprintf("%s", destFile))
	if err != nil {
		logrus.Fatalf("tar task failed: %+v", err)
	}
	defer fw.Close()

	gw := gzip.NewWriter(fw)
	defer gw.Close()

	tw := tar.NewWriter(gw)
	defer tw.Close()

	filepath.Walk(srcDir, func(fileName string, fi fs.FileInfo, err error) error {
		fileName = strings.ReplaceAll(fileName, "\\", "/")
		if err != nil {
			return err
		}

		hdr, err := tar.FileInfoHeader(fi, "")
		if err != nil {
			return err
		}

		hdr.Name = strings.TrimPrefix(fileName, fmt.Sprintf("%s/", srcDir))

		if err = tw.WriteHeader(hdr); err != nil {
			return err
		}

		if !fi.Mode().IsRegular() {
			return nil
		}

		fr, err := os.Open(fileName)
		defer fr.Close()

		if err != nil {
			return err
		}

		n, err := io.Copy(tw, fr)
		if err != nil {
			return err
		}

		logrus.Debugf("tar %s, size: %d", fileName, n)
		return err
	})
}
