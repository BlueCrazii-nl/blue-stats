package background

import (
	"fmt"
	"log"
	"reflect"
	"time"

	utils "github.com/Yadiiiig/blue-stats/internals/utils"
)

func Save(c *utils.Utilities) {
	go func() {
		for {
			s := reflect.ValueOf(c.ActiveUsers).Elem()
			typeOfT := s.Type()
			for i := 0; i < s.NumField(); i++ {
				f := s.Field(i)
				if f.Len() != 0 {
					query := fmt.Sprintf(`INSERT INTO views (timeun, amount, category) VALUES (%d, %d, '%s')`, time.Now().Unix(), f.Len(), typeOfT.Field(i).Name)
					_, err := c.Connection.Connection.Exec(query)
					if err != nil {
						log.Println(err)
					}
				}
			}
			time.Sleep(15 * time.Second)
		}
	}()
}
