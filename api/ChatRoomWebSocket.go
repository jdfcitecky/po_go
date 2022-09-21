package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"po_go/cache"
	"po_go/entity"
	"po_go/utils"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func ChatRoomSocketHandler(c *gin.Context) {
	logger := utils.Log()
	logger.Info("-----------------------In the socket ", c.Param("id"))
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
		panic(err)
	}
	// defer close function
	defer func() {
		closeSocketErr := ws.Close()
		logger.Info("THE WEBSOCKET IS CLOSE")
		if closeSocketErr != nil {
			panic(err)
		}
	}()
	utils.Cache.Set(fmt.Sprintf("chatroom%s", c.Param("id")), "-----------------THIS IS THE DEFAULT MESSAGE-----------------", cache.DefaultExpiration)
	//define the behavior when websocket is established
	for {
		chatRoomMessage := entity.ChatRoomMessage{}
		// !!!This line has I/O reader so will stuck the for loop
		msgType, msg, _ := ws.ReadMessage()
		logger.Info("-----------------------In the socket ", c.Param("id"), fmt.Sprintf("Message Type: %d, Message: %s\n", msgType, string(msg)))
		json.Unmarshal(msg, &chatRoomMessage)
		if err != nil {
			panic(err)
		}
		msgSend := tryGetMessage(c.Param("id"))
		if msgSend != -1 {
			err = ws.WriteJSON(msgSend)
			if err != nil {
				panic(err)
			}
		}

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
