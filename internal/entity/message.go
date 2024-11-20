package entity

type Message struct {
    ID        string `json:"id"`
    SenderID  string `json:"sender_id"`
    GroupID   string `json:"group_id,omitempty"` // Guruhga tegishli bo'lsa
    Content   string `json:"content"`
    Timestamp int64  `json:"timestamp"`
    FileURL   string `json:"file_url,omitempty"` // Fayl bo'lsa
}
