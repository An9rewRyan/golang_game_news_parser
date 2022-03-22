package main

const MAX_amount_of_goroutines = 10

//in millisecokonds
const Links_loading_retry_time = 200
const Article_loading_retry_time = 100

var MAX_amount_of_loading_retries = 3 //not a constant beacause of use in error formatting
