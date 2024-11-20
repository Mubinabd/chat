package entity

type Message struct {
	SenderID  string `json:"sender_id"`
	Content   string `json:"content"`
	Timestamp string `json:"timestamp"`
}

type MessageReq struct {
	GroupID string `json:"group_id"`
	Message string `json:"message"`
}