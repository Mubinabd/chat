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
