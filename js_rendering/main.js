const express = require('express')
const app = express()
const connect = require("./connect.js");

async function get_js_rendered_page (link) {
    const browser = await connect();
    console.log("Processing page: "+link)
    page = await browser.newPage();
    await page.setDefaultNavigationTimeout(0); 
    console.log("Waiting for loading...: "+link)
    await page.goto(link, {
      waitUntil: 'load',
    })
    console.log("Waiting for html...: "+link)
    await page.waitForSelector("html")
    console.log("Getting html out...: "+link)
    let bodyHTML = await page.evaluate(() => document.documentElement.outerHTML);
    console.log("Html is out...: "+link)
    await page.close();
    return bodyHTML;
}

app.use(
  express.urlencoded({
    extended: true
  })
)
app.use(express.json())
app.post('/', (req, res) => {
  console.log("Got request to "+req.body.link)
  let data = get_js_rendered_page(req.body.link)
  data.then(res.send.bind(res)).catch(error => {
    console.log(error+" promise rejected!")
    res.send("Couldnt load page")
  })
})
console.log("server is running!")
app.listen(8000);
