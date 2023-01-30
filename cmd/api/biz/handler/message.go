package handler

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/json"
	"github.com/hertz-contrib/websocket"
)

var upgrader = websocket.HertzUpgrader{} // use default options

type ServerMessage struct {
	FromUserId int64  `json:"from_user_id"`
	MsgContent string `json:"msg_content"`
}

type ClientMessage struct {
	UserId     int64  `json:"user_id"`
	ToUserId   int64  `json:"to_user_id"`
	MsgContent string `json:"msg_content"`
}

// Message 收发消息api
func Message(_ context.Context, c *app.RequestContext) {
	err := upgrader.Upgrade(c, func(conn *websocket.Conn) {
		for {
			mt, msg, err := conn.ReadMessage()
			if err != nil {
				hlog.Error(err)
				break
			}
			var clientMessage ClientMessage
			err = json.Unmarshal(msg, &clientMessage)
			if err != nil {
				hlog.Error(err)
				return
			}
			hlog.Infof("recv: %+v", clientMessage)

			serverMessage := ServerMessage{
				FromUserId: clientMessage.UserId,
				MsgContent: clientMessage.MsgContent,
			}
			msgRes, err := json.Marshal(serverMessage)
			if err != nil {
				hlog.Error(err)
				return
			}

			err = conn.WriteMessage(mt, msgRes)
			if err != nil {
				hlog.Error("write:", err)
				break
			}
		}
	})
	if err != nil {
		hlog.Error("upgrade:", err)
	}
}
