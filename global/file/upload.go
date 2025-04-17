package file

import (
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func UploadFile(c *gin.Context, filename string) (string, error) {

	// Retrieve the uploaded file
	file, err := c.FormFile(filename)
	if err != nil {
		return "", err
	}

	// Save the file
	newFilename := uuid.New().String() + filepath.Ext(file.Filename)

	if err := c.SaveUploadedFile(file, "./upload/"+newFilename); err != nil {
		return "", err
	}

	return newFilename, nil
}
