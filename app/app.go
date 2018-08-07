package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	_ "github.com/go-sql-driver/mysql"
)

// Config is a project wide config for connections
type Config struct {
	Database struct {
		Driver string `json:"driver"`
		Host   string `json:"host"`
		Base   string `json:"base"`
		Port   string `json:"port"`
		User   string `json:"user"`
		Pass   string `json:"pass"`
	} `json:"database"`
	Bot struct {
		Token string `json:"token"`
	}
}

// Route is a command routing struct
type Route struct {
	Description string
	Handler     func()
}

// Routes is a pseudo routing map for our command strings
type Routes map[string]Route

func NewRouter() *Router {
	return &Router{
		rules: make(map[string]Handler),
	}
}

func main() {
	config := loadConfiguration("config.json")

	app := App{}
	app.Connect()

	dg, err := discordgo.New(config.Bot.Token)
	checkErr(err)

	dg.AddHandler(ready)
	dg.AddHandler(app.messageCreate)

	err = dg.Open()
	checkErr(err)

	fmt.Println("Bot está online. Aperte CTRL-C para sair")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
}

// Connect application in database
func (app *App) Connect() {
	config := loadConfiguration("config.json")

	dbDriver := config.Database.Driver
	dbHost := config.Database.Host
	dbUser := config.Database.User
	dbPass := config.Database.Pass
	dbName := config.Database.Base

	conn, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+dbHost+")/"+dbName)
	checkErr(err)

	app.Connection = conn
}

func ready(s *discordgo.Session, event *discordgo.Ready) {
	s.UpdateStatus(0, "Ingredientes na panela")
}

type Router struct {
	rules map[string]Handler
}

func (r *Router) Handle(msgName string, handler Handler) {
	r.rules[msgName] = handler
}

func (r *Router) FindHandler(msgName string) (Handler, bool) {
	handler, found := r.rules[msgName]
	return handler, found
}

type Handler func(*Client)
type FindHandler func(string) (Handler, bool)

type Client struct {
	findHandler FindHandler
}

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

	router.Handle("!criar", app.createOrder)
	router.Handle("!finalizar", app.closeOrder)
	router.Handle("!pedir", app.addItem)
	router.Handle("!cancelar", app.removeItem)
	router.Handle("!pedidos", app.listItems)
	router.Handle("!sortear", app.raffle)

	client := NewClient(router.FindHandler)
	command := strings.Split(m.Content, " ")[0]

	if handler, found := client.findHandler(command); found {
		handler(client)
	}

}

func loadConfiguration(file string) Config {
	var config Config
	configFile, err := os.Open(file)
	defer configFile.Close()
	checkErr(err)

	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
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
