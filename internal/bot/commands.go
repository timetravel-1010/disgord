package bot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/timetravel-1010/disgord/internal/bot/handlers"
	"github.com/timetravel-1010/disgord/internal/bot/handlers/components"
)

var (
	integerOptionMinValue          = 1.0
	dmPermission                   = false
	defaultMemberPermissions int64 = discordgo.PermissionManageServer

	commands = []*discordgo.ApplicationCommand{
		handlers.BasicCmd,
		handlers.PermissionOverviewCmd,
		handlers.BasicCmdWithFile,
		handlers.LocalizedCmd,
		handlers.OptionsCmd,
		handlers.SubCommands,
		handlers.ResponsesCmd,
		handlers.FollowUpsCmd,
		handlers.BasicCmd,
		handlers.SelectsCmd,
		handlers.ImageCmd,
	}

	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		handlers.BasicCmd.Name:              handlers.BasicCommandHandler,
		handlers.BasicCmdWithFile.Name:      handlers.BasicCommandWithFileHandler,
		handlers.LocalizedCmd.Name:          handlers.LocalizedCmdHandler,
		handlers.OptionsCmd.Name:            handlers.OptionsCmdHandler,
		handlers.PermissionOverviewCmd.Name: handlers.PermissionsOverviewCmdHandler,
		handlers.SubCommands.Name:           handlers.SubcommandsHandler,
		handlers.ResponsesCmd.Name:          handlers.ResponsesHandler,
		handlers.FollowUpsCmd.Name:          handlers.FollowUpsHandler,
		handlers.ButtonsCmd.Name:            handlers.ButtonsCmdHandler,
		handlers.SelectsCmd.Name:            handlers.SelectsCmdHandler,
		handlers.ImageCmd.Name:              handlers.ImgCmdHandler,
	}

	componentHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"fd_no":              components.FdNoCompHandler,
		"fd_yes":             components.FdYesCompHandler,
		"select":             components.SelectCompHandler,
		"stackoverflow_tags": components.StackOverflowTagsCompHandler,
		"channel_select":     components.StackOverflowTagsCompHandler,
	}
)
