package main

import (
	"log"
	"net/http"
	"os"

	background "github.com/Yadiiiig/blue-stats/internals/background"
	functionality "github.com/Yadiiiig/blue-stats/internals/functionality"
	router "github.com/Yadiiiig/blue-stats/internals/router"
	utils "github.com/Yadiiiig/blue-stats/internals/utils"
)

func main() {
	conn, err := utils.ConnectDatabase("stats.sqlite3")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	users := utils.ActiveUsers{
		F1ENG:      make(map[string]int64),
		F1NL:       make(map[string]int64),
		FESTIVALS:  make(map[string]int64),
		AFTERPARTY: make(map[string]int64),
	}

	connection := utils.Connection{
		Connection: conn,
	}

	c := utils.Utilities{
		Connection:  &connection,
		ActiveUsers: &users,
	}

	collection := functionality.NewCollection(&c)

	background.Save(&c)
	background.UpdateViewers(&c)

	mux := router.InitRouter()
	mux.InitRoutes(collection)

	defer conn.Close()
	log.Fatal(http.ListenAndServe(":8080", mux.Router))
}
