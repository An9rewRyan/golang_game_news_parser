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
    // await fs.truncate('/mnt/d/go/parser/golang_game_news_parser/loaded.html', 0, function(){console.log('done')});
    // let written = await fs.writeFile('/mnt/d/go/parser/golang_game_news_parser/loaded.html', bodyHTML, err => {
    //     if (err) {
    //       console.error(err);
    //       return 'Failed!';
    //     }
    // })
    console.log(bodyHTML)
    await browser.close();
    return bodyHTML;
}


app.use(
  express.urlencoded({
    extended: true
  })
)
app.use(express.json())
app.post('/', (req, res) => {
  console.log(req.body)
  let data = get_js_rendered_page(req.body.link)
  data.then(res.send.bind(res)).catch(error => {console.log(error+"promise rejected!")})
})
app.listen(8000);
