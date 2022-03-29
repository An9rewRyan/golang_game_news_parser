package config

import (
	"os"

	"github.com/tebeka/selenium"
)

const MAX_amount_of_goroutines = 5

//in millisecokonds
const Links_loading_retry_time = 250
const Article_loading_retry_time = 1000

//
const Db_conn_str string = "user=postgres password=1234 dbname=gg host = db port = 5432 sslmode=disable"

// in minutes
const Sleep_time = 10

//
var MAX_amount_of_loading_retries = 10 //not a constant beacause of use in error formatting

const Selenium_path = "selenium/selenium-server.jar"
const Gecko_driver_path = "selenium/geckodriver"
const Chrome_driver_path = "selenium/chromedriver"
const Selenium_port = 8080

var Selenium_opts = []selenium.ServiceOption{
	selenium.StartFrameBuffer(),             // Start an X frame buffer for the browser to run in.
	selenium.GeckoDriver(Gecko_driver_path), // Specify the path to GeckoDriver in order to use Firefox.
	selenium.Output(os.Stderr),
}
