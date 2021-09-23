package functionality

import (
	"encoding/json"
	"fmt"
	"net/http"

	utils "github.com/Yadiiiig/blue-stats/internals/utils"
)

func (s *Collection) Watcher(w http.ResponseWriter, r *http.Request) {
	var b Body
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil || b.Stream == "" {
		w.WriteHeader(404)
		return
	}

	id, userAgent := utils.GetId(r)

	// Can probably be replaced with CORS
	if !utils.CheckAgent(userAgent) {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, "You are not allowed to perform this action.")
		return
	}

	if !s.Saved.SaveUser(id, b.Stream) {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, "Please stop trying.")
		return
	} else {
		w.WriteHeader(http.StatusOK)
		return
	}
}
