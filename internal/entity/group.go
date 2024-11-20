package entity

type Group struct {
    ID      string   `json:"id"`
    Name    string   `json:"name"`
    Members []string `json:"members"` // User ID'lar
}
