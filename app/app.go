package main

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
	_ "github.com/go-sql-driver/mysql"
)

// Route is a command routing struct
type Route struct {
	Description string
	Handler     func()
}

// Routes is a pseudo routing map for our command strings
type Routes map[string]Route

// NewRouter ...
func NewRouter() *Router {
	return &Router{
		rules: make(map[string]Handler),
	}
}

func main() {

	dg, err := discordgo.New("Bot " + os.Getenv("DG_TOKEN"))
	if err != nil {
		fmt.Println("Failed to create discord session", err)
	}

	dg.AddHandler(ready)
	dg.AddHandler(app.messageCreate)

	err = dg.Open()
	if err != nil {
		fmt.Println("Unable to establish connection", err)
	}

	defer dg.Close()

	<-make(chan struct{})
}

func ready(s *discordgo.Session, event *discordgo.Ready) {
	s.UpdateStatus(0, "Ingredientes na panela")
}

// Router ...
type Router struct {
	rules map[string]Handler
}

// Handle ...
func (r *Router) Handle(msgName string, handler Handler) {
	r.rules[msgName] = handler
}

// FindHandler ...
func (r *Router) FindHandler(msgName string) (Handler, bool) {
	handler, found := r.rules[msgName]
	return handler, found
}

// Handler ...
type Handler func(*Client)

// FindHandler ...
type FindHandler func(string) (Handler, bool)

// Client ...
type Client struct {
	findHandler FindHandler
}

// NewClient ...
func NewClient(findHandler FindHandler) *Client {
	return &Client{
		findHandler: findHandler,
	}
}

func (app *App) messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		return
	}

	app.Session = s
	app.Message = m

	router := NewRouter()

	router.Handle("!criar", app.OpenList)
	router.Handle("!finalizar", app.CloseList)
	router.Handle("!pedir", app.AddItem)
	router.Handle("!cancelar", app.RemoveItem)
	router.Handle("!pedidos", app.ItemsByList)
	router.Handle("!sortear", app.raffle)

	client := NewClient(router.FindHandler)
	command := strings.Split(m.Content, " ")[0]

	if handler, found := client.findHandler(command); found {
		handler(client)
	}
}

func checkCount(rows *sql.Rows) (count int) {
	for rows.Next() {
		err := rows.Scan(&count)
		checkErr(err)
	}
	return count
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
