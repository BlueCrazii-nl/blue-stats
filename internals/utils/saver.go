package utils

import (
	"reflect"
	"strings"
	"time"
)

// this code will be simplified in the future with yml entries/reflection
func (s *ActiveUsers) SaveUser(id string, stream string) bool {
	t := strings.SplitAfter(stream, "/")
	streamName := strings.ToUpper(strings.Replace(t[len(t)-1], ".m3u8", "", -1))

	c := reflect.ValueOf(s).Elem()
	typeOf := c.Type()
	for i := 0; i < c.NumField(); i++ {
		f := c.Field(i)
		if typeOf.Field(i).Name == streamName {
			user := f.Interface().(map[string]int64)[id]
			if user == 0 {
				f.Interface().(map[string]int64)[id] = time.Now().Unix()
				return true
			} else {
				if time.Now().Unix()-user < 15 {
					return false
				} else {
					f.Interface().(map[string]int64)[id] = time.Now().Unix()
					return true
				}
			}
		}
	}
	return false
}
