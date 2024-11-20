package entity

import "time"

type Group struct {
    ID          string    `json:"id"`          
    Name        string    `json:"name"`        
    Description string    `json:"description"` 
    CreatedAt   time.Time `json:"created_at"`   
    UpdatedAt   time.Time `json:"updated_at"`   
    Messages    []Message `json:"messages"`     
}
