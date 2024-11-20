package entity

type Group struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Users   []string  `json:"users"`  
	Messages  []Message `json:"messages"` 
	CreatedAt string    `json:"created_at"`
}

type CreateGroupReq struct {
	Name        string `json:"name" validate:"required"`
	UserID string `json:"user_id" validate:"required"`
}

type AddFileToGroupReq struct {
	GroupID string `json:"group_id"`
	FileID  string `json:"file_id"`
}

type GetGroupByIDReq struct {
	ID string `json:"id"`
}
type ListGroupsReq struct {
	Limit  int32  `json:"limit"`
	Offset int32  `json:"offset"`
	Search string `json:"search,omitempty"`
}

type ListGroupsRes struct {
	Groups []*Group `json:"groups"`
	Count  int32   `json:"count"`
}

type UpdateGroupReq struct {
	ID          string `json:"id" validate:"required"`
	Name        string `json:"name,omitempty"`
}

type DeleteGroupReq struct {
	ID string `json:"id" validate:"required"`
}
