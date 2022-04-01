package config

import (
	"sync"
)

const MAX_amount_of_goroutines = 5

//in millisecokonds
const Links_loading_retry_time = 250
const Article_loading_retry_time = 1000

//pq
const Db_conn_str string = "user=postgres password=1234 dbname=gg host = db port = 5432 sslmode=disable"

// const Db_conn_str string = "postgres://postgres:1234@db:5432/gg"

// in minutes
const Sleep_time = 10

//
var MAX_amount_of_loading_retries = 15 //not a constant beacause of use in error formatting
var Wg_main sync.WaitGroup

var Exited_links = make([]string, 10) //made for handling fucked up goroutines
