package fullInit

import (
	"github.com/bwmarrin/discordgo"
	"github.com/qbxt/status-benson/constants"
	"github.com/qbxt/status-benson/structures"
)

var Bot *discordgo.Session
var Pings map[string]structures.Ping
var IncomingPings chan structures.Ping

func Init() error {
	var err error
	Bot, err = discordgo.New("Bot " + constants.TOKEN)
	Pings = make(map[string]structures.Ping)
	IncomingPings = make(chan structures.Ping)
	return err
}