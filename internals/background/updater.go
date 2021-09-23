package background

import (
	"reflect"
	"time"

	utils "github.com/Yadiiiig/blue-stats/internals/utils"
)

func UpdateViewers(c *utils.Utilities) {
	go func() {
		for {
			s := reflect.ValueOf(c.ActiveUsers).Elem()
			for i := 0; i < s.NumField(); i++ {
				f := s.Field(i)
				if f.Len() != 0 {
					for k, v := range f.Interface().(map[string]int64) {
						if time.Now().Unix()-v > 15 {
							delete(f.Interface().(map[string]int64), k)
						}
					}
				}
			}
			time.Sleep(15 * time.Second)
		}
	}()
}
