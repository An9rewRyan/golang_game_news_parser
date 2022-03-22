package main

const MAX_amount_of_goroutines = 5

//in millisecokonds
const Links_loading_retry_time = 250
const Article_loading_retry_time = 1000

var MAX_amount_of_loading_retries = 10 //not a constant beacause of use in error formatting
