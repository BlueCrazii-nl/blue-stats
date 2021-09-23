package utils

import "database/sql"

type Utilities struct {
	Connection  *Connection
	ActiveUsers *ActiveUsers
}

type ActiveUsers struct {
	F1ENG      map[string]int64
	F1NL       map[string]int64
	FESTIVALS  map[string]int64
	AFTERPARTY map[string]int64
}

type Connection struct {
	Connection *sql.DB
}

type User struct {
	LastSeen int64
}
