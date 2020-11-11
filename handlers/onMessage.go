package handlers

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/qbxt/status-benson/constants"
	"github.com/qbxt/status-benson/fullInit"
	"strings"
	"time"
)

func OnMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	if strings.ToLower(m.Content) == "!!status" {
		printStatus(s, m)
		return
	}
}

func printStatus(s *discordgo.Session, m *discordgo.MessageCreate) {
	offlineCount := 0
	totalCount := 0

	e := &discordgo.MessageEmbed{
		Title: "NMT Status",
		Fields: make([]*discordgo.MessageEmbedField, 0),
	}

	for _, ping := range fullInit.Pings { // read-only, this could be safe idk
		totalCount++
		if isOnline(ping.LastSeen) {
			f := &discordgo.MessageEmbedField {
				Name: constants.ApprovedIDs[ping.ID],
				Value: fmt.Sprintf("%s %sms", constants.EmojiGreen, ping.Value),
			}
			e.Fields = append(e.Fields, f)
		} else {
			f := &discordgo.MessageEmbedField {
				Name: constants.ApprovedIDs[ping.ID],
				Value: fmt.Sprintf("%s Last seen %f minutes ago", constants.EmojiRed, time.Since(time.Unix(ping.LastSeen, 0)).Minutes()),
			}
			e.Fields = append(e.Fields, f)
			offlineCount++
		}
	}

	if offlineCount == totalCount {
		e.Color = constants.ColorRed
	} else if offlineCount == 0 {
		e.Color = constants.ColorGreen
	} else {
		e.Color = constants.ColorYellow
	}

	_, _ = s.ChannelMessageSendEmbed(m.ChannelID, e)
}

func isOnline(timestamp int64) bool { // return true if time since last ping is < 60s
	if time.Since(time.Unix(timestamp, 0)) > time.Minute {
		return false
	} else {
		return true
	}
}