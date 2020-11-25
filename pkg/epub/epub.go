package epub

import (
	"archive/zip"
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"

	"go.uber.org/zap"
)

const mimetypePath = "mimetype"
const epubMimetype = "application/epub+zip"
const containerPath = "META-INF/container.xml"

var (
	ErrFileNotFound = errors.New("epub: no '%s' found in container")
	ErrNoMimetype   = errors.New("epub: no mimetype found in container")
	// ErrNoRootfile occurs when there are no rootfile entries found in
	// container.xml.
	ErrNoRootfile = errors.New("epub: no rootfile found in container")

	// ErrBadRootfile occurs when container.xml references a rootfile that does
	// not exist in the zip.
	ErrBadRootfile = errors.New("epub: container references non-existent rootfile")

	// ErrNoItemref occurrs when a content.opf contains a spine without any
	// itemref entries.
	ErrNoItemref = errors.New("epub: no itemrefs found in spine")

	// ErrBadItemref occurs when an itemref entry in content.opf references an
	// item that does not exist in the manifest.
	ErrBadItemref = errors.New("epub: itemref references non-existent item")

	// ErrBadManifest occurs when a manifest in content.opf references an item
	// that does not exist in the zip.
	ErrBadManifest = errors.New("epub: manifest references non-existent item")
)

type EpubReader struct {
	Name  string
	Files map[string]*zip.File
}

type EpubReaderCloser struct {
	EpubReader
	f *os.File
}

// Logger
var Logger *zap.SugaredLogger

func init() {
	logger, _ := zap.NewDevelopment()
	Logger = logger.Sugar()
}

func OpenReader(filename string) (*EpubReaderCloser, error) {
	zipFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	zipStat, err := zipFile.Stat()
	if err != nil {
		zipFile.Close()
		return nil, err
	}

	zipReader, err := zip.NewReader(zipFile, zipStat.Size())
	if err != nil {
		zipFile.Close()
		return nil, err
	}

	reader := new(EpubReaderCloser)
	reader.Name = filename
	reader.f = zipFile

	if err = reader.init(zipReader); err != nil {
		return nil, err
	}

	return reader, nil
}

func (epubReader *EpubReader) init(zipReader *zip.Reader) error {
	epubReader.Files = make(map[string]*zip.File)
	for _, f := range zipReader.File {
		epubReader.Files[f.Name] = f
	}

	if mimetype, err := epubReader.readFile(mimetypePath); err != nil {
		Logger.Infof("file %s is not an epub (no mimetype)", epubReader.Name)
		return err
	} else if mimetype != epubMimetype {
		Logger.Infof("file %s is not an epub (invalid mimetype)", epubReader.Name)
		return ErrNoMimetype

	}

	if container, ok := epubReader.Files[containerPath]; ok {
		fmt.Printf(container.Name)
		Logger.Info("file is an epub")
	} else {
		Logger.Error("file is not an epub")
	}

	return nil
}

func (epubReader *EpubReader) readFile(name string) (string, error) {
	file, ok := epubReader.Files[name]
	if !ok {
		return "", fmt.Errorf("epub: no '%s' found in container", name)
	}

	reader, err := file.Open()
	if err != nil {
		return "", err
	}
	defer reader.Close()

	var buffer bytes.Buffer
	_, err = io.Copy(&buffer, reader)
	if err != nil {
		return "", err
	}

	return buffer.String(), nil
}
