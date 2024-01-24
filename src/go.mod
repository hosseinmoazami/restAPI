module hossein.dev

go 1.21.5

replace hossein.dev/backend => ./backend

replace hossein.dev/util => ./util

replace hossein.dev/db => ./db

require hossein.dev/backend v0.0.0-00010101000000-000000000000

require (
	github.com/gorilla/mux v1.8.1 // indirect
	github.com/mattn/go-sqlite3 v1.14.19 // indirect
)
