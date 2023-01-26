package main

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"log"
	"os"
	"path"
)

func compressFile(filename string, output string) error {
	file, err := os.Create(output)
	if err != nil {
		return err
	}

	writer, err := gzip.NewWriterLevel(file, gzip.BestCompression)
	if err != nil {
		return err
	}

	tw := tar.NewWriter(writer)

	fileToCompress, err := os.Open(filename)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(fileToCompress)
	if err != nil {
		return err
	}

	hdr := &tar.Header{
		Name: path.Base(filename),
		Mode: 0600,
		Size: int64(len(body)),
	}
	err = tw.WriteHeader(hdr)
	if err != nil {
		return err
	}

	_, err = tw.Write(body)
	if err != nil {
		return err
	}
	log.Print("Compressed " + filename + " to " + output)
	if err := tw.Close(); err != nil {
		return err
	}
	if err := writer.Close(); err != nil {
		return err
	}
	if err := file.Close(); err != nil {
		return err
	}
	if err := fileToCompress.Close(); err != nil {
		return err
	}
	return nil
}

func decompressFile(filename string, output string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}

	reader, err := gzip.NewReader(file)
	if err != nil {
		return err
	}

	tr := tar.NewReader(reader)

	hdr, err := tr.Next()
	if err != nil {
		return err
	}

	fileToDecompress, err := os.Create(output)
	if err != nil {
		return err
	}

	_, err = io.Copy(fileToDecompress, tr)
	if err != nil {
		return err
	}
	log.Print("Decompressed " + hdr.Name + " to " + output)
	if err := fileToDecompress.Close(); err != nil {
		return err
	}
	if err := reader.Close(); err != nil {
		return err
	}
	if err := file.Close(); err != nil {
		return err
	}
	return nil
}
