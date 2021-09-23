package functionality

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/gorilla/mux"
)

func (s *Collection) LiveViewers(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if vars["id"] == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Missing id parameter")
		return
	}

	c := reflect.ValueOf(s.Saved).Elem()
	typeOf := c.Type()
	for i := 0; i < c.NumField(); i++ {
		f := c.Field(i)
		if typeOf.Field(i).Name == strings.ToUpper(vars["id"]) {
			fmt.Fprintf(w, "%d", f.Len())
			return
		}
	}

	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprintf(w, "Id: %s is incorrect.", vars["id"])
	return
}
