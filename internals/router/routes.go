package router

import (
	functionality "github.com/Yadiiiig/blue-stats/internals/functionality"
	"github.com/gorilla/mux"
)

type Mux struct {
	Router *mux.Router
}

func InitRouter() *Mux {
	return &Mux{
		Router: mux.NewRouter().StrictSlash(true),
	}
}

func (mux *Mux) InitRoutes(c *functionality.Collection) {
	mux.Router.HandleFunc("/watcher", c.Watcher).Methods("POST")
	mux.Router.HandleFunc("/views/{id}", c.LiveViewers).Methods("GET")
	mux.Router.HandleFunc("/stats/{id}/{start}/{end}", c.StatsBetween).Methods("GET")
}
