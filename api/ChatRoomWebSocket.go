package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"po_go/cache"
	"po_go/chatroomhub"
	"po_go/entity"
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
	logger.Info("-----------------------In the socket ", c.Param("chatroomid"))
	var chatRoomRegistData ChatRoomRegistData
	err := c.ShouldBindUri(&chatRoomRegistData)
	logger.Info("-----------------------In the socket bind Regist data", chatRoomRegistData)
	if err != nil {
		logger.Info("-----------------------In the socket bind Error", err)
		// panic(err)
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
		panic(err)
	}
	// regist this websocket to hub
	chatroomhub.Hub.MessageRegistChan <- chatroomhub.RegistMessage{ChatRoomID: chatRoomRegistData.ChatRoomID, SenderID: chatRoomRegistData.SenderID, Websocket: ws}
	// defer close function
	defer func() {
		chatroomhub.Hub.MessageUnRegistChan <- chatroomhub.RegistMessage{ChatRoomID: chatRoomRegistData.ChatRoomID, SenderID: chatRoomRegistData.SenderID, Websocket: ws}
		closeSocketErr := ws.Close()
		logger.Info("THE WEBSOCKET IS CLOSE")
		if closeSocketErr != nil {
			logger.Info("-----------------------In the socket close Error", err)
			panic(err)
		}
	}()
	utils.Cache.Set(fmt.Sprintf("chatroom%s", c.Param("id")), "-----------------THIS IS THE DEFAULT MESSAGE-----------------", cache.DefaultExpiration)
	//define the behavior when websocket is established
	for {
		chatRoomMessage := entity.ChatRoomMessage{}
		// !!!This line has I/O reader so will stuck the for loop
		msgType, msg, err := ws.ReadMessage()
		if err != nil {
			logger.Info("-----------------------In the socket read Error", err)
			panic(err)
		}
		logger.Info("-----------------------In the socket ", c.Param("id"), fmt.Sprintf("Message Type: %d, Message: %s\n", msgType, string(msg)))
		json.Unmarshal(msg, &chatRoomMessage)
		chatroomhub.Hub.MessageComingChan <- chatRoomMessage
		// msgSend := tryGetMessage(c.Param("id"))
		// if msgSend != -1 {
		// 	err = ws.WriteJSON(msgSend)
		// 	if err != nil {
		// 		panic(err)
		// 	}
		// }

	}
}

func setMessage(chatRoomID string, message entity.ChatRoomMessage) {
	utils.Cache.Set(fmt.Sprintf("chatroom%s", chatRoomID), message, cache.DefaultExpiration)
}

func tryGetMessage(chatRoomID string) interface{} {
	messages, found := utils.Cache.Get(fmt.Sprintf("chatroom%s", chatRoomID))
	if found == false {
		return -1
	}
	return messages
}
