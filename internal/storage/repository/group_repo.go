package repository

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/Mubinabd/chat/internal/entity"
)

type GroupRepo struct {
	db *sql.DB
}

func NewGroupRepo(db *sql.DB) *GroupRepo {
	return &GroupRepo{
		db: db,
	}
}

func (r *GroupRepo) CreateGroup(group *entity.CreateGroupReq) (*entity.Void, error) {
	var id string
	query := `INSERT INTO groups (name, user_id) VALUES ($1, $2) RETURNING id`
	err := r.db.QueryRow(query, group.Name, group.UserID).Scan(&id)
	if err != nil {
		return nil, err
	}
	return &entity.Void{}, nil
}

func (r *GroupRepo) GetGroupByID(req *entity.GetGroupByIDReq) (*entity.Group, error) {
	res := &entity.Group{}

	query := `SELECT id, name, users,messages FROM groups WHERE id = $1`
	err := r.db.QueryRow(query, req.ID).
		Scan(
			&res.ID,
			&res.Name,
			&res.Users,
			&res.Messages,
		)
	if err == sql.ErrNoRows {
		return nil, err
	} else if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *GroupRepo) AddMemberToGroup(req *entity.AddFileToGroupReq) (*entity.Void, error) {
	var groupExists bool
	query := `SELECT EXISTS (SELECT 1 FROM groups WHERE id = $1)`
	err := r.db.QueryRow(query, req.GroupID).Scan(&groupExists)
	if err != nil {
		return nil, fmt.Errorf("error checking group existence: %w", err)
	}
	if !groupExists {
		return nil, fmt.Errorf("group with ID %s not found", req.GroupID)
	}

	updateQuery := `
		UPDATE groups
		SET files = array_append(files, $1)
		WHERE id = $2
	`
	_, err = r.db.Exec(updateQuery, req.FileID, req.GroupID)
	if err != nil {
		return nil, fmt.Errorf("error adding file to group: %w", err)
	}

	return &entity.Void{}, nil
}

func (r *GroupRepo) DeleteGroup(req *entity.DeleteGroupReq) (*entity.Void, error) {
	res := &entity.Void{}

	query := `UPDATE groups SET deleted_at = EXTRACT(EPOCH FROM NOW()) WHERE id = $1`
	_, err := r.db.Exec(query, req.ID)
	if err != nil {
		return nil, err
	}

	return res, nil
}
func (r *GroupRepo) UpdateGroup(req *entity.UpdateGroupReq) (*entity.Void, error) {
	res := &entity.Void{}

	query := `UPDATE groups SET updated_at = NOW()`

	var arg []interface{}
	var conditions []string

	if req.Name != "" && req.Name != "string" {
		arg = append(arg, req.Name)
		conditions = append(conditions, fmt.Sprintf("name = $%d", len(arg)))
	}

	if len(conditions) > 0 {
		query += ", " + strings.Join(conditions, ", ")
	}

	query += fmt.Sprintf(" WHERE id = $%d", len(arg)+1)
	arg = append(arg, req.ID)
	_, err := r.db.Exec(query, arg...)
	if err != nil {
		return nil, err
	}

	return res, nil
}


func (r *GroupRepo) GetAllGroups(req *entity.ListGroupsReq) (*entity.ListGroupsRes, error) {
	res := &entity.ListGroupsRes{}

	query := `SELECT 
				id, 
				name, 
				users,
				messages 
			FROM 
				groups 
			WHERE 
				deleted_at=0 `

	var args []interface{}


	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", len(args)+1, len(args)+2)
	args = append(args, req.Limit, req.Offset)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var group entity.Group
		err := rows.Scan(
			&group.ID,
			&group.Name,
			&group.Users,
			&group.Messages,
		)
		if err != nil {
			return nil, err
		}
		res.Groups = append(res.Groups, &group)
	}

	return res, nil
}
