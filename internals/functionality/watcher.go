package functionality

import (
	"encoding/json"
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

	defer r.Body.Close()

	id := utils.GetId(r)

	s.Lock()
	if !s.Saved.SaveUser(id, b.Stream) {
		w.WriteHeader(http.StatusForbidden)
		s.Unlock()
		return
	} else {
		w.WriteHeader(http.StatusOK)
		s.Unlock()
		return
	}
}
