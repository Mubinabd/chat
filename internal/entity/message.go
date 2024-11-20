package entity

import "time"

type Message struct {
    ID        string    `json:"id"`         
    GroupID   string    `json:"group_id"`   
    Sender    string    `json:"sender"`     
    Content   string    `json:"content"`    
    SentAt    time.Time `json:"sent_at"`    
    UpdatedAt time.Time `json:"updated_at"` 
}
