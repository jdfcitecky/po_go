package chatroomhub

import (
	"po_go/entity"
	"po_go/utils"
	"strconv"

	"github.com/gorilla/websocket"
)

type ChatRoom struct {
	ChatRoomID    int
	WebSocketList map[string]*websocket.Conn
}

type ChatRoomHub struct {
	ChatRoomList        map[string]ChatRoom
	MessageComingChan   chan entity.ChatRoomMessage
	MessageRegistChan   chan RegistMessage
	MessageUnRegistChan chan RegistMessage
}

type RegistMessage struct {
	ChatRoomID int
	SenderID   int
	Websocket  *websocket.Conn
}

var Hub ChatRoomHub

func init() {
	// just init the hub
	Hub.ChatRoomList = map[string]ChatRoom{}
	// remembewr to make the channel
	Hub.MessageComingChan = make(chan entity.ChatRoomMessage)
	Hub.MessageRegistChan = make(chan RegistMessage)
	Hub.MessageUnRegistChan = make(chan RegistMessage)
}

func ListenMessageAndFoward() {
	logger := utils.Log()
	logger.Info("Chat room Hub working")
	for {
		select {
		case newMessage := <-Hub.MessageComingChan:
			chatRoomList := Hub.ChatRoomList[strconv.Itoa(newMessage.ChatRoomID)].WebSocketList
			// write the message to all socket
			for _, ws := range chatRoomList {
				err := ws.WriteJSON(newMessage)
				if err != nil {
					logger.Info(err)
				}
			}
		case newMessage := <-Hub.MessageRegistChan:
			RegistToHub(newMessage)
			// logger.Info("********************Chat room Hub registed", Hub.ChatRoomList)
		case newMessage := <-Hub.MessageUnRegistChan:
			UnRegistToHub(newMessage)
			// logger.Info("********************Chat room Hub Unregisted", Hub.ChatRoomList)
		}
	}
}

func RegistToHub(defaultMessage RegistMessage) {
	chatRoomId := strconv.Itoa(defaultMessage.ChatRoomID)
	senderId := strconv.Itoa(defaultMessage.SenderID)
	// if the chat room has been registed
	if chatRoom, ok := Hub.ChatRoomList[chatRoomId]; ok {
		chatRoom.WebSocketList[senderId] = defaultMessage.Websocket
		Hub.ChatRoomList[chatRoomId] = chatRoom
	} else {
		// if the chat room is not registed, new a chat room with the web socket and append to chat room list
		chatRoom := ChatRoom{ChatRoomID: defaultMessage.ChatRoomID, WebSocketList: map[string]*websocket.Conn{}}
		chatRoom.ChatRoomID = defaultMessage.ChatRoomID
		chatRoom.WebSocketList[senderId] = defaultMessage.Websocket
		Hub.ChatRoomList[chatRoomId] = chatRoom
	}
}

func UnRegistToHub(defaultMessage RegistMessage) {
	chatRoomId := strconv.Itoa(defaultMessage.ChatRoomID)
	senderId := strconv.Itoa(defaultMessage.SenderID)
	// if the chat room has been registed
	if chatRoom, ok := Hub.ChatRoomList[chatRoomId]; ok {
		delete(chatRoom.WebSocketList, senderId)
		Hub.ChatRoomList[chatRoomId] = chatRoom
		// if all member leave caht room, remove it from chat room list
		if len(chatRoom.WebSocketList) == 0 {
			delete(Hub.ChatRoomList, chatRoomId)
		}
	}

}
