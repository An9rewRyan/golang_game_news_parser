package config

import (
	"sync"
)

const MAX_amount_of_goroutines = 5

//in millisecokonds
const Links_loading_retry_time = 250
const Article_loading_retry_time = 1000

//pq
// const Db_conn_str string = "user=postgres password=1234 dbname=gg host = db port = 5432 sslmode=disable"
const Db_conn_str string = "postgres://ifbubrbxajiqqu:78f6ddf168308eb1afe067ecdc0ea4b6d4a5c325a014fe1bce1995066040a7b3@ec2-52-18-116-67.eu-west-1.compute.amazonaws.com:5432/dhq8ntfel79e3"

// const Db_conn_str string = "postgres://postgres:1234@db:5432/gg"

// in minutes
const Sleep_time = 10

//
var MAX_amount_of_loading_retries = 15 //not a constant beacause of use in error formatting
var Wg_main sync.WaitGroup

var Exited_links = make([]string, 10) //made for handling fucked up goroutines
