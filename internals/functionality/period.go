package functionality

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (s *Collection) StatsBetween(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if vars["id"] == "" || vars["start"] == "" || vars["end"] == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Missing parameter")
		return
	}

	start, err := strconv.Atoi(vars["start"])
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid start parameter")
		return
	}

	end, err := strconv.Atoi(vars["end"])
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid end parameter")
		return
	}

	query := fmt.Sprintf(`SELECT SUM(amount) FROM views WHERE category = '%s' AND timeun >= %d AND timeun <= %d`, vars["id"], start, end)
	row, err := s.Connection.Connection.Query(query)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Something went wrong")
	}
	var count int
	for row.Next() {
		row.Scan(&count)
	}

	fmt.Fprintf(w, "%d", count)
}
