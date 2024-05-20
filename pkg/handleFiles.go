package pkg

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
)

func UploadFile(filenames []*multipart.FileHeader, filePath string) (string, error) {
	cwd, err := os.Getwd()

	if err != nil {
		fmt.Println("Error:", err)
	}

	var uploadedFileNames []string
	for _, fileHeader := range filenames {

		src, err := fileHeader.Open()
		if err != nil {
			return "", err
		}
		defer src.Close()

		newFilename := uuid.New().String() + "_" + time.Now().Format("20060102150405") + "_" + fileHeader.Filename
		dstPath := cwd + filePath + newFilename

		dst, err := os.Create(dstPath)
		if err != nil {
			return "", err
		}
		defer dst.Close()

		if _, err = io.Copy(dst, src); err != nil {
			return "", err
		}

		uploadedFileNames = append(uploadedFileNames, newFilename)
	}
	return strings.Join(uploadedFileNames, ";"), nil
}

func DeleteFile(filenamesString *string, filePath string) error {
	cwd, err := os.Getwd()

	if err != nil {
		fmt.Println("Error:", err)
	}

	filenames := strings.Split(*filenamesString, ";")

	for _, filename := range filenames {

		filePath := cwd + filePath + filename

		if err := os.Remove(filePath); err != nil {
			return err
		}

	}

	return nil
}
