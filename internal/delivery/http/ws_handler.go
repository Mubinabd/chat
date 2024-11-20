package http

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/Mubinabd/chat/internal/usecase"
	"github.com/gorilla/websocket"
)

type WebSocketHandler struct {
	FileUC  usecase.FileUseCase
	GroupUC usecase.GroupUseCase
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true }, 
}

var groupConnections = make(map[string][]*websocket.Conn)

func (h *WebSocketHandler) HandleConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Failed to upgrade WebSocket:", err)
		return
	}
	defer conn.Close()

	groupID := r.URL.Query().Get("group_id")
	if groupID == "" {
		log.Println("Missing group_id in query parameters")
		return
	}

	groupConnections[groupID] = append(groupConnections[groupID], conn)
	log.Printf("New connection added to group %s", groupID)

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading WebSocket message:", err)
			break
		}

		for _, memberConn := range groupConnections[groupID] {
			if memberConn == conn {
				continue
			}
			err := memberConn.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				log.Println("Error sending WebSocket message:", err)
			}
		}
	}

	h.removeConnection(groupID, conn)
}

func (h *WebSocketHandler) removeConnection(groupID string, conn *websocket.Conn) {
	connections := groupConnections[groupID]
	for i, c := range connections {
		if c == conn {
			groupConnections[groupID] = append(connections[:i], connections[i+1:]...)
			break
		}
	}
}


func (h *WebSocketHandler) HandleFileUpload(w http.ResponseWriter, r *http.Request) {
    file, header, err := r.FormFile("file")
    if err != nil {
        http.Error(w, "Failed to upload file", http.StatusBadRequest)
        return
    }
    defer file.Close()

    filePath := "/path/to/store/" + header.Filename
    out, err := os.Create(filePath) 
    if err != nil {
        http.Error(w, "Failed to save file", http.StatusInternalServerError)
        return
    }
    defer out.Close()

    _, err = io.Copy(out, file) 
    if err != nil {
        http.Error(w, "Failed to save file", http.StatusInternalServerError)
        return
    }

    w.Write([]byte("File uploaded successfully: " + filePath))
}