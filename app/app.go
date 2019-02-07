package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func main() {

	discord, err := discordgo.New("Bot " + os.Getenv("DG_TOKEN"))
	if err != nil {
		fmt.Println("Failed to create discord session", err)
	}

	discord.AddHandler(ready)
	discord.AddHandler(messageCreate)

	err = discord.Open()
	if err != nil {
		fmt.Println("Unable to establish connection", err)
	}

	fmt.Println("Listening...")
	lock := make(chan int)
	<-lock
}

func ready(s *discordgo.Session, event *discordgo.Ready) {
	s.UpdateStatus(0, "Ingredientes na panela")
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	router := NewRouter()
	router.Handle("!criar", OpenList)
	// router.Handle("!finalizar", CloseList(s, m))
	// router.Handle("!pedir", AddItem(s, m))
	// router.Handle("!cancelar", RemoveItem(s, m))
	// router.Handle("!pedidos", ItemsByList(s, m))
	// router.Handle("!sortear", raffle(s, m))

	bc := NewBotCommand(router.FindHandler, s, m)
	command := strings.Split(m.Content, " ")[0]

	if handler, found := bc.findHandler(command); found {
		handler(bc)
	}
}
