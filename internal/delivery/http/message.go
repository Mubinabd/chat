package http

import (
	"net/http"

	"github.com/Mubinabd/chat/internal/entity"
	"github.com/gin-gonic/gin"
)

// AddMessageToGroup godoc
// @Summary Add a message to a group
// @Description Adds a new message to the specified group
// @Accept json
// @Produce json
// @Tags group
// @Security BearerAuth
// @Param message body entity.MessageReq true "Message information"
// @Success 200 {object} entity.Void "Message added successfully"
// @Failure 400 {object} string "Bad request"
// @Failure 500 {object} string "Internal server error"
// @Router /v1/groups/messages [post]
func (h *Handlers) AddMessageToGroup(c *gin.Context) {
	var req entity.MessageReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Message": "Invalid input"})
		return
	}

	_, err := h.Message.AddMessageToGroup(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, entity.Void{})
}

// @Summary Upload and Save File
// @Description Upload a file to the server, save it, and store metadata
// @Accept multipart/form-data
// @Produce json
// @Tags file
// @Security BearerAuth
// @Param file formData file true "File to be uploaded"
// @Success 200 {object} entity.FileResponse
// @Failure 400 {object} gin.H "Bad request"
// @Failure 500 {object} gin.H "Internal server error"
// @Router /v1/files/upload [post]
func (h *Handlers) SaveFile(c *gin.Context) {
	// Retrieve the file from the request
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to retrieve file"})
		return
	}

	// Specify the file paths
	srcFilePath := file.Filename
	destFilePath := "./uploads/" + file.Filename

	// Save the uploaded file locally
	if err := c.SaveUploadedFile(file, srcFilePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save uploaded file"})
		return
	}

	// Prepare the repository request
	req := &entity.FileSave{
		FileName: srcFilePath,
		FilePath: destFilePath,
	}

	// Save the file using the repository
	response, err := h.File.SaveFile(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond with success
	c.JSON(http.StatusOK, response)
}
