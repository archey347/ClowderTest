package main

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

func downloadZip(base url.URL, name string) error {
	//https://open.kattis.com/problems/4thought/file/statement/samples.zip
	path := base.String() + "/problems/" + name + "/file/statement/samples.zip"

	resp, err := http.Get(path)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath.FromSlash(".cache/download.zip"))
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)

	return err
}

func unzip(name string) error {
	r, err := zip.OpenReader(filepath.FromSlash(".cache/download.zip"))

	if err != nil {
		return err
	}

	defer r.Close()

	dest := filepath.Join(".cache", name)

	for _, f := range r.File {

		fPath := filepath.Join(dest, f.Name)

		if !strings.HasPrefix(fPath, filepath.Clean(dest)+string(os.PathSeparator)) {
			return fmt.Errorf("%s: illegal file path", fPath)
		}

		if f.FileInfo().IsDir() {
			// Make Folder
			os.MkdirAll(fPath, os.ModePerm)
			continue
		}

		// Make File
		if err = os.MkdirAll(filepath.Dir(fPath), os.ModePerm); err != nil {
			return err
		}

		outFile, err := os.OpenFile(fPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}

		rc, err := f.Open()
		if err != nil {
			return err
		}

		_, err = io.Copy(outFile, rc)

		// Close the file without defer to close before next iteration of loop
		outFile.Close()
		rc.Close()

		if err != nil {
			return err
		}
	}
	return nil
}

func getTests(base url.URL, name string, cache bool) ([]test, error) {

	if _, err := os.Stat(".cache"); os.IsNotExist(err) {
		err = os.Mkdir(".cache", os.ModePerm)

		if err != nil {
			return nil, err
		}
	}

	// Check if one already exists in .tests
	testDir := filepath.FromSlash(".cache/" + name)

	download := false

	if _, err := os.Stat(testDir); os.IsNotExist(err) {
		download = true
	} else {
		if !cache {
			err := os.RemoveAll(testDir)
			if err != nil {
				return nil, err
			}
			download = true
		}
	}

	if download {
		fmt.Print("Downloading... ")
		err := downloadZip(base, name)

		if err != nil {
			fmt.Print("FAILED!")
			return nil, err
		}

		fmt.Println("done.")

		fmt.Print("Extracting... ")
		err = unzip(name)

		if err != nil {
			fmt.Println("FAILED!")
			return nil, err
		}

		fmt.Println("done.")
	}

	// Delete if no cache

	return nil, nil
}
