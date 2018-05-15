package utils

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path"
	"strconv"
	"time"
)

func CreateFile(file *multipart.FileHeader, saveTo string) (string, error) {

	if CheckExt(file.Filename) {

		src, err := file.Open()
		if err != nil {
			return "", err
		}
		defer src.Close()

		timestamp := strconv.FormatInt(time.Now().Unix(), 10)
		filename := file.Filename + timestamp
		filename = path.Join(saveTo, Sha1(filename)+".jpg")

		pwd, _ := os.Getwd()
		// root, _ := filepath.Abs(filepath.Dir(os.Args[0]))

		pathFile := path.Join(pwd, PUBLIC, filename)

		// Destination
		dst, err := os.Create(pathFile)
		if err != nil {
			return "", err
		}
		defer dst.Close()

		// Copy
		if _, err = io.Copy(dst, src); err != nil {
			return "", err
		}

		return filename, nil
	}

	return "", fmt.Errorf("Not supported!")
}

func DeleteFile(pathfile string) {
	pwd, _ := os.Getwd()
	fmt.Println("remove", path.Join(pwd, pathfile))
	err := os.Remove(path.Join(pwd, pathfile))
	if err != nil {
		fmt.Println(err)
	}
}
