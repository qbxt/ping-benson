package handlers

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/qbxt/status-benson/logger"
)

func OnReady(s *discordgo.Session, r *discordgo.Ready) {
	logger.Info(fmt.Sprintf("%s#%s is fully logged in and ready", s.State.User.Username, s.State.User.Discriminator), nil)
}