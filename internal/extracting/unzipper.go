package extracting

import (
	"archive/zip"
	"github.com/dasalgadoc/e-Invoicing-Processor/kit/domain/errors"
	"io"
	"os"
	"path"
)

const unzipper = "extracting/unzipper.go"

func unzipFileToDestination(originFolder, fileName, destination string) *errors.ProjectError {
	fileToExtract := path.Join(originFolder, fileName)
	file, err := zip.OpenReader(fileToExtract)
	if err != nil {
		outputPath := destination + "/" + fileName
		err := copyFile(fileToExtract, outputPath)
		if err != nil {
			return errors.NewProjectError(unzipper, errors.OSError, err.Error())
		}

		return nil
	}
	defer file.Close()

	if err := os.MkdirAll(destination, os.ModePerm); err != nil {
		return errors.NewProjectError(unzipper, errors.OSError, err.Error())
	}

	for _, f := range file.File {
		entryFile, err := f.Open()
		if err != nil {
			return errors.NewProjectError(unzipper, errors.OSError, err.Error())
		}
		defer entryFile.Close()

		outputPath := destination + "/" + f.Name
		outputFile, err := os.Create(outputPath)
		if err != nil {
			return errors.NewProjectError(unzipper, errors.OSError, err.Error())
		}
		defer outputFile.Close()

		_, err = io.Copy(outputFile, entryFile)
		if err != nil {
			return errors.NewProjectError(unzipper, errors.OSError, err.Error())
		}
	}

	return nil
}

func copyFile(source, destination string) error {
	entryFile, err := os.Open(source)
	if err != nil {
		return errors.NewProjectError(unzipper, errors.OSError, err.Error())
	}
	defer entryFile.Close()

	outputFile, err := os.Create(destination)
	if err != nil {
		return errors.NewProjectError(unzipper, errors.OSError, err.Error())
	}
	defer outputFile.Close()

	_, err = io.Copy(outputFile, entryFile)
	if err != nil {
		return errors.NewProjectError(unzipper, errors.OSError, err.Error())
	}

	return nil
}
