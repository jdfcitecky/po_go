package api

import (
	"encoding/json"
	"net/http"
	"po_go/chatroomhub"
	"po_go/entity"
	"po_go/service"
	"po_go/utils"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type ChatRoomRegistData struct {
	ChatRoomID int `uri:"chatroomid"`
	SenderID   int `uri:"senderid"`
}

func ChatRoomSocketHandler(c *gin.Context) {
	logger := utils.Log()
	var chatRoomRegistData ChatRoomRegistData
	err := c.ShouldBindUri(&chatRoomRegistData)
	if err != nil {
		logger.Info("-----------------------In the socket bind Error", err)
		// panic(err)
		c.Abort()
	}
	// upgrade this connection to websocket
	upGrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		logger.Info("-----------------------In the socket upgrade Error", err)
		// panic(err)
		c.Abort()
	}
	// regist this websocket to hub
	chatroomhub.Hub.MessageRegistChan <- chatroomhub.RegistMessage{ChatRoomID: chatRoomRegistData.ChatRoomID, SenderID: chatRoomRegistData.SenderID, Websocket: ws}
	// defer close function
	defer func() {
		chatroomhub.Hub.MessageUnRegistChan <- chatroomhub.RegistMessage{ChatRoomID: chatRoomRegistData.ChatRoomID, SenderID: chatRoomRegistData.SenderID, Websocket: ws}
		closeSocketErr := ws.Close()
		if closeSocketErr != nil {
			logger.Info("-----------------------In the socket close Error", err)
			// panic(err)
		}
	}()
	//define the behavior when websocket is established
	for {
		chatRoomMessages := new(service.ChatRoomMessage)
		chatRoomMessage := entity.ChatRoomMessage{}
		// !!!This line has I/O reader so will stuck the for loop
		_, msg, err := ws.ReadMessage()
		if err != nil {
			// panic(err)
			break
		}
		// logger.Info("-----------------------In the socket ", c.Param("id"), fmt.Sprintf("Message Type: %d, Message: %s\n", msgType, string(msg)))
		// this is for hub notify
		json.Unmarshal(msg, &chatRoomMessage)
		chatroomhub.Hub.MessageComingChan <- chatRoomMessage
		// this is for sql insert
		json.Unmarshal(msg, &chatRoomMessages)
		chatRoomMessages.Insert()

	}
}
