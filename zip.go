package main

import (
	"archive/zip"
	"compress/flate"
	"io"
	"log"
	"os"
	"path/filepath"
)

func compress() {
	zipFileName := "archive.zip"

	zipFile, err := os.Create(zipFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer zipFile.Close()

	w := zip.NewWriter(zipFile)
	defer w.Close()

	addFilesToZip(w, "archive", "")
	w.RegisterCompressor(zip.Deflate, func(out io.Writer) (io.WriteCloser, error) {
		return flate.NewWriter(out, flate.BestCompression)
	})
}

func addFilesToZip(zipWriter *zip.Writer, sourcePath, basePath string) {
	sourceInfo, err := os.Stat(sourcePath)
	if err != nil {
		log.Fatal(err)
	}

	header, err := zip.FileInfoHeader(sourceInfo)
	if err != nil {
		log.Fatal(err)
	}
	if basePath != "" {
		header.Name = filepath.Join(basePath, header.Name)
	}

	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		log.Fatal(err)
	}

	if sourceInfo.IsDir() {
		entries, err := os.ReadDir(sourcePath)
		if err != nil {
			log.Fatal(err)
		}

		for _, entry := range entries {
			entryPath := filepath.Join(sourcePath, entry.Name())
			entryBasePath := filepath.Join(basePath, sourceInfo.Name())

			addFilesToZip(zipWriter, entryPath, entryBasePath)
		}
	} else {
		sourceFile, err := os.Open(sourcePath)
		if err != nil {
			log.Fatal(err)
		}
		defer sourceFile.Close()

		_, err = io.Copy(writer, sourceFile)
		if err != nil {
			log.Fatal(err)
		}
	}
}
