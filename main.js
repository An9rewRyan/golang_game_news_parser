const puppeteer = require("puppeteer");
const fs = require("fs");
const express = require('express')
const app = express()

async function get_js_rendered_page (link) {
    const browser = await puppeteer.launch({
      slowMo: 150,
    });
    const page = await browser.newPage();

    await page.goto(link);
    let bodyHTML = await page.evaluate(() => document.documentElement.outerHTML);
    await page.screenshot({ path: 'example.png' });
    console.log(bodyHTML);
    await fs.truncate('/mnt/d/go/parser/golang_game_news_parser/loaded.html', 0, function(){console.log('done')});
    await fs.writeFile('/mnt/d/go/parser/golang_game_news_parser/loaded.html', bodyHTML, err => {
        if (err) {
          console.error(err);
          return 'Failed!';
        }
    })
    console.log("Cool!")
    await browser.close();
    return 'Processed!';
}


app.use(
  express.urlencoded({
    extended: true
  })
)
app.use(express.json())
app.post('/', (req, res) => {
  console.log(req.body)
})
app.listen(8000);
// const http = require('http');
// const requestListener = function (req, res) {
//   console.log(req.body)
//   console.log.apply(req.)
//   res.writeHead(200);
//   // res.end(get_js_rendered_page(req.body));
// }
// const server = http.createServer(requestListener);
// server.listen(8080);
// // get_js_rendered_page("https://kanobu.ru/videogames//")