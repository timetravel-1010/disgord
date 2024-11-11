package bot

import (
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	"github.com/timetravel-1010/disgord/config"
)

type Bot struct {
	session *discordgo.Session
	config  *config.Config
}

func NewBot() *Bot {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	s, err := discordgo.New("Bot " + cfg.Token)
	if err != nil {
		log.Fatalf("Invalid bot parameters: %v", err)
	}

	err = s.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}

	return &Bot{
		session: s,
		config:  cfg,
	}
}

func (b *Bot) Start() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop
}

func (b *Bot) RegisterCommands() {
	b.session.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as: %v#%v", b.session.State.User.Username, s.State.User.Discriminator)
	})

	b.session.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		switch i.Type {
		case discordgo.InteractionApplicationCommand:
			if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
				h(s, i)
			}
		case discordgo.InteractionMessageComponent:
			if h, ok := componentHandlers[i.MessageComponentData().CustomID]; ok {
				h(s, i)
			}
		}
	})

	log.Println("Adding commands...")

	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands))
	for i, v := range commands {
		cmd, err := b.session.ApplicationCommandCreate(b.session.State.User.ID, "", v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
		registeredCommands[i] = cmd
	}
}

func (b *Bot) RemoveCommands() {
	log.Println("Removing commands...")
	// We need to fetch the commands, since deleting requires the command ID.
	// We are doing this from the returned commands on line 375, because using
	// this will delete all the commands, which might not be desirable, so we
	// are deleting only the commands that we added.
	registeredCommands, err := b.session.ApplicationCommands(b.session.State.User.ID, "")
	if err != nil {
		log.Fatalf("Could not fetch registered commands: %v", err)
	}

	for _, v := range registeredCommands {
		err := b.session.ApplicationCommandDelete(b.session.State.User.ID, "", v.ID)
		if err != nil {
			log.Panicf("Cannot delete '%v' command: %v", v.Name, err)
		}
	}
}

func (b *Bot) Stop() {
	err := b.session.Close()
	if err != nil {
		log.Fatalf("Error closing the websocket: %s", err)
	}

	log.Println("Gracefully shutting down.")
}
