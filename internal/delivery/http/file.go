package http

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// @Summary Upload file
// @Description Upload a file to the server and save it
// @Accept multipart/form-data
// @Produce json
// @Tags file
// @Security BearerAuth
// @Param file formData file true "File to be uploaded"
// @Success 200 {object} FileResponse
// @Failure 400 {object} gin.H "Bad request"
// @Failure 500 {object} gin.H "Internal server error"
// @Router /v1/upload [post]
type FileResponse struct {
	Message  string `json:"message"`
	FilePath string `json:"file_path"`
}

func (h *Handlers) SaveFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to upload file"})
		return
	}

	
	filePath := "./uploads/" + file.Filename
	if _, err := os.Stat("./uploads/"); os.IsNotExist(err) {
		if err := os.Mkdir("./uploads/", os.ModePerm); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create directory"})
			return
		}
	}

	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	c.JSON(http.StatusOK, FileResponse{
		Message:  "File uploaded successfully",
		FilePath: filePath,
	})
}
