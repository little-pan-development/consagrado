package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"regexp"
	"strconv"
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

// Cart is our itens wrapper
type Cart struct {
	ID          uint
	Description string
	Item        []Item
}

// Item is any order made
type Item struct {
	ID            uint
	Description   string
	DiscordUserId string
}

// Route is a command routing struct
type Route struct {
	Description string
	Handler     func()
}

// Routes is a pseudo routing map for our command strings
type Routes map[string]Route

func main() {
	config := loadConfiguration("config.json")

	dg, err := discordgo.New(config.Bot.Token)
	checkErr(err)

	dg.AddHandler(ready)
	dg.AddHandler(messageCreate)

	err = dg.Open()
	checkErr(err)

	fmt.Println("Bot está online. Aperte CTRL-C para sair")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()

}

func ready(s *discordgo.Session, event *discordgo.Ready) {
	s.UpdateStatus(0, "Ingredientes na panela")
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	db := dbConn()
	defer db.Close()

	if m.Author.ID == s.State.User.ID {
		return
	}

	// routes := loadRoutes()
	// handleRoute(m.Content, routes)

	// Cria um carrinho
	if strings.HasPrefix(m.Content, "!criar") {

		splitRegexp := regexp.MustCompile("[\n| ]")
		split := splitRegexp.Split(m.Content, 2)

		if len(split) == 1 {
			_, err := s.ChannelMessageSend(m.ChannelID, "Digite uma descrição para seu carrinho!")
			checkErr(err)
			return
		}

		rows, err := db.Query("SELECT COUNT(*) FROM cart WHERE status = 1")
		checkErr(err)

		if checkCount(rows) > 0 {
			_, err := s.ChannelMessageSend(m.ChannelID, "Existe um carrinho em aberto!")
			checkErr(err)
			return
		}

		stmt, err := db.Prepare("INSERT cart SET description = ?, status = ?")
		checkErr(err)

		res, err := stmt.Exec(split[1], 1)
		checkErr(err)

		id, err := res.LastInsertId()
		checkErr(err)

		s.UpdateStatus(0, "Faça seu pedido..")

		idToString := strconv.FormatInt(int64(id), 10)

		s.ChannelMessageSend(m.ChannelID, "Carrinho `#"+idToString+" "+split[1]+"` criado com sucesso!")
	}

	// Finaliza carrinho
	if strings.HasPrefix(m.Content, "!finalizar") {

		stmt, err := db.Prepare("update cart set status = ? where status = ?")
		checkErr(err)

		res, err := stmt.Exec(0, 1)
		checkErr(err)

		affect, err := res.RowsAffected()
		checkErr(err)

		fmt.Println(affect)

		s.UpdateStatus(0, "Ingredientes na panela.")
		s.ChannelMessageSend(m.ChannelID, "@here **Pedidos finalizados!**")
	}

	// Insere pedido no carrinho
	if strings.HasPrefix(m.Content, "!pedir") {

		splitRegexp := regexp.MustCompile("[\n| ]")
		split := splitRegexp.Split(m.Content, 2)

		if len(split) == 1 {
			_, err := s.ChannelMessageSend(m.ChannelID, m.Author.Mention()+", digite seu pedido. Por exemplo, `!pedir Lentilha da vó` :heart:")
			checkErr(err)
			return
		}

		var cart Cart
		row := db.QueryRow("SELECT id, description FROM cart WHERE status = 1")
		err := row.Scan(&cart.ID, &cart.Description)

		switch err {
		case sql.ErrNoRows:
			s.ChannelMessageSend(m.ChannelID, m.Author.Mention()+", antes de pedirem, utilize `!criar nome do carrinho` para **criar um novo carrinho**.")
			return
		default:
			checkErr(err)
		}

		fmt.Printf("%v", cart.ID)

		rows, err := db.Query("SELECT COUNT(*) FROM item WHERE discord_user_id = ? AND cart_id = ?", m.Author.ID, cart.ID)
		checkErr(err)

		if checkCount(rows) > 0 {
			_, err := s.ChannelMessageSend(m.ChannelID, m.Author.Mention()+" você já realizou seu pedido. Para **cancelar** digite `!cancelar`")
			checkErr(err)
			return
		}

		stmt, err := db.Prepare("INSERT item SET description = ?, cart_id = ?, discord_user_id = ?")
		checkErr(err)

		_, err = stmt.Exec(split[1], cart.ID, m.Author.ID)
		checkErr(err)

		s.ChannelMessageSend(m.ChannelID, m.Author.Mention()+" seu **pedido foi realizado** com sucesso.")
	}

	// Retira pedido do carrinho
	if strings.HasPrefix(m.Content, "!cancelar") {
		var item Item
		row := db.QueryRow("select i.id from cart c inner join item i on c.id = i.cart_id where c.status = 1 and i.discord_user_id = ?", m.Author.ID)
		err := row.Scan(&item.ID)

		fmt.Printf("%v", item.ID)
		// select i.id from cart c inner join item i on c.id = i.cart_id where c.status = 1 and i.discord_user_id = "186909290475290624";
		stmt, err := db.Prepare("delete from item where id = ?")
		checkErr(err)

		res, err := stmt.Exec(item.ID)
		checkErr(err)

		affect, err := res.RowsAffected()
		checkErr(err)

		fmt.Println(affect)

		s.ChannelMessageSend(m.ChannelID, m.Author.Mention()+" seu pedido foi **cancelado** com sucesso!")
	}

	// Lista todos os pedidos
	if strings.HasPrefix(m.Content, "!pedidos") {

		var cart Cart
		row := db.QueryRow("SELECT id, description FROM cart WHERE status = 1")
		err := row.Scan(&cart.ID, &cart.Description)

		rows, err := db.Query("SELECT description, discord_user_id FROM item WHERE cart_id = ?", cart.ID)

		embed := &discordgo.MessageEmbed{}

		embed.Title = "Pedidos até o momento:"
		embed.Description = "**--** :hamburger: **--**"
		embed.Color = 0xff0000

		embed.Author = &discordgo.MessageEmbedAuthor{}
		embed.Author.Name = "Palmirinha!"
		embed.Author.URL = "https://www.facebook.com/vovopalmirinha/"
		embed.Author.IconURL = "https://i.imgur.com/QTDVdLK.jpg"

		embed.Fields = []*discordgo.MessageEmbedField{}

		for rows.Next() {
			var item Item
			err = rows.Scan(&item.Description, &item.DiscordUserId)
			checkErr(err)

			var user, _ = s.User(item.DiscordUserId)

			embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
				Name:   "\n\n**" + user.Username + "**",
				Value:  item.Description,
				Inline: false,
			})
		}

		s.ChannelMessageSendEmbed(m.ChannelID, embed)
	}

	// Sortear um dos donos de pedidos abertos para pedir
	if strings.HasPrefix(m.Content, "!sortear") {

		var discordUserID string
		row := db.QueryRow("SELECT i.discord_user_id FROM cart c JOIN item i ON i.cart_id = c.id WHERE c.status = 1 ORDER BY RAND() LIMIT 1")
		err := row.Scan(&discordUserID)

		if err != nil {
			fmt.Println(err)
		}

		var user, _ = s.User(discordUserID)

		embed := &discordgo.MessageEmbed{}

		embed.Title = "Parabéns! Hoje é com..."
		embed.Description = user.Mention() + " contamos com você!"
		embed.Color = 0xff0000

		embed.Author = &discordgo.MessageEmbedAuthor{}
		embed.Author.Name = "Palmirinha!"
		embed.Author.URL = "https://www.facebook.com/vovopalmirinha/"
		embed.Author.IconURL = "https://i.imgur.com/QTDVdLK.jpg"

		s.ChannelMessageSendEmbed(m.ChannelID, embed)
	}
}

func dbConn() (db *sql.DB) {

	config := loadConfiguration("config.json")

	dbDriver := config.Database.Driver
	dbHost := config.Database.Host
	dbUser := config.Database.User
	dbPass := config.Database.Pass
	dbName := config.Database.Base

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+dbHost+")/"+dbName)
	checkErr(err)

	return db
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

func loadRoutes() Routes {
	var routes Routes

	// routes["!pedidos"] = Route{
	// 	Description: "Listar todos pedidos do último carrinho aberto",
	// 	Handler:     listCartContent,
	// }
	//
	// routes["!pedir"] = Route{
	// 	Description: "Listar todos pedidos do último carrinho aberto",
	// 	Handler:     listCartContent,
	// }

	return routes
}

func handleRoute(content string, routes Routes) {
	parts := strings.SplitN(strings.TrimLeft(content, " "), " ", 1)

	if len(parts[0]) > 1 {

	}

}

func listCartContent() {

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
