package http

import (
	"database/sql"
	"net/http"

	"github.com/Mubinabd/chat/internal/entity"
	"github.com/Mubinabd/chat/internal/storage/repository"
	"github.com/gin-gonic/gin"
)

// CreateGroup godoc
// @Summary Create a new group
// @Description Create a new group with the provided name and members
// @Accept json
// @Produce json
// @Tags group
// @Security BearerAuth
// @Param group body entity.CreateGroupReq true "Group information"
// @Success 200 {object} entity.Void "Group created successfully"
// @Failure 400 {object} string "Bad request"
// @Failure 500 {object} string "Internal server error"
// @Router /v1/groups [post]
func (h *Handlers) CreateGroup(c *gin.Context) {
	var groupReq entity.CreateGroupReq
	if err := c.ShouldBindJSON(&groupReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Message": "Invalid input"})
		return
	}

	groupRepo := repository.NewGroupRepo(&sql.DB{})
	_, err := groupRepo.CreateGroup(&groupReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, entity.Void{})
}

// GetGroupByID godoc
// @Summary Get group by ID
// @Description Retrieve group details by ID
// @Accept json
// @Produce json
// @Tags group
// @Security BearerAuth
// @Param id path string true "Group ID"
// @Success 200 {object} entity.Group "Group details retrieved successfully"
// @Failure 400 {object} string "Bad request"
// @Failure 404 {object} string "Group not found"
// @Failure 500 {object} string "Internal server error"
// @Router /v1/groups/{id} [get]
func (h *Handlers) GetGroupByID(c *gin.Context) {
	groupID := c.Param("id")

	group, err := h.Group.GetGroupByID(c, &entity.GetGroupByIDReq{ID: groupID})
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"message": "Group not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, group)
}

// AddMemberToGroup godoc
// @Summary Add a member to a group
// @Description Add a new member to the specified group
// @Accept json
// @Produce json
// @Tags group
// @Security BearerAuth
// @Param request body entity.AddFileToGroupReq true "Add member to group request"
// @Success 200 {object} entity.Void "Member added successfully"
// @Failure 400 {object} string "Bad request"
// @Failure 404 {object} string "Group not found"
// @Failure 500 {object} string "Internal server error"
// @Router /v1/groups/{groupID}/add-member [post]
func (h *Handlers) AddMemberToGroup(c *gin.Context) {
	var req entity.AddFileToGroupReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Message": "Invalid input"})
		return
	}

	_, err := h.Group.AddMemberToGroup(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, entity.Void{})
}

// DeleteGroup godoc
// @Summary Delete a group
// @Description Mark a group as deleted by setting its deleted_at timestamp
// @Accept json
// @Produce json
// @Tags group
// @Security BearerAuth
// @Param id path string true "Group ID"
// @Success 200 {object} entity.Void "Group deleted successfully"
// @Failure 400 {object} string "Bad request"
// @Failure 500 {object} string "Internal server error"
// @Router /v1/groups/{id} [delete]
func (h *Handlers) DeleteGroup(c *gin.Context) {
	groupID := c.Param("id")

	_, err := h.Group.DeleteGroup(c, &entity.DeleteGroupReq{ID: groupID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, entity.Void{})
}

// UpdateGroup godoc
// @Summary Update a group
// @Description Update the group details such as name
// @Accept json
// @Produce json
// @Tags group
// @Security BearerAuth
// @Param group body entity.UpdateGroupReq true "Group information to update"
// @Success 200 {object} entity.Void "Group updated successfully"
// @Failure 400 {object} string "Bad request"
// @Failure 500 {object} string "Internal server error"
// @Router /v1/groups/{id} [put]
func (h *Handlers) UpdateGroup(c *gin.Context) {
	var req entity.UpdateGroupReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Message": "Invalid input"})
		return
	}

	_, err := h.Group.UpdateGroup(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, entity.Void{})
}

// GetAllGroups godoc
// @Summary Get all groups
// @Description Retrieve a list of all groups with pagination
// @Accept json
// @Produce json
// @Tags group
// @Security BearerAuth
// @Param limit query int false "Limit of groups to retrieve"
// @Param offset query int false "Offset for pagination"
// @Success 200 {object} entity.ListGroupsRes "List of groups retrieved successfully"
// @Failure 500 {object} string "Internal server error"
// @Router /v1/groups [get]
func (h *Handlers) GetAllGroups(c *gin.Context) {
	var req entity.ListGroupsReq
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Message": "Invalid query parameters"})
		return
	}

	res, err := h.Group.GetAllGroups(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
