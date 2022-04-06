const express = require('express')
const app = express()
const connect = require("./connect.js");

// let processed_pages = 0
// let page_failed = false

const waitTillHTMLRendered = async (page, timeout = 30000) => {
  const checkDurationMsecs = 1000;
  const maxChecks = timeout / checkDurationMsecs;
  let lastHTMLSize = 0;
  let checkCounts = 1;
  let countStableSizeIterations = 0;
  const minStableSizeIterations = 3;

  while(checkCounts++ <= maxChecks){
    let html = await page.content();
    let currentHTMLSize = html.length; 

    let bodyHTMLSize = await page.evaluate(() => document.body.innerHTML.length);

    console.log('last: ', lastHTMLSize, ' <> current: ', currentHTMLSize, " body html size: ", bodyHTMLSize);

    if(lastHTMLSize != 0 && currentHTMLSize == lastHTMLSize) 
      countStableSizeIterations++;
    else 
      countStableSizeIterations = 0; //reset the counter

    if(countStableSizeIterations >= minStableSizeIterations) {
      console.log("Page rendered fully..");
      break;
    }

    lastHTMLSize = currentHTMLSize;
    await page.waitFor(checkDurationMsecs);
  }  
};

async function get_js_rendered_page (link) {
    const browser = await connect();
    if (link === "finished loading links, thank u, dear puppeteer server!"){
      // let pages = await browser.pages();
      // for (let page of pages){
      //   await page.close()
      //   console.log(`one page of ${pages.length} closed`)
      // }
      await browser.close()
      console.log("You so sweet...:p")
      return "browser closed"
    }
    console.log("Processing page: "+link)
    page = await browser.newPage();
    await page.setDefaultNavigationTimeout(0); 
    console.log("Waiting for loading...: "+link)
    await page.goto(link, {
      waitUntil: 'load',
    })
    await waitTillHTMLRendered(page)
    console.log("Waiting for html...: "+link)
    await page.waitForSelector("html")
    console.log("Getting html out...: "+link)
    let bodyHTML = await page.evaluate(() => document.documentElement.outerHTML);
    console.log("Html is out...: "+link)
    // await page.close();
    return bodyHTML;
}

app.use(
  express.urlencoded({
    extended: true
  })
)

app.use(express.json())
app.post('/', (req, res) => {
  if (req.body.link === "finished loading links, thank u, dear puppeteer server!"){
    (async ()=>{
      let result = await get_js_rendered_page(req.body.link)
      console.log(result)
    })()
    let message = "You are wellcome, golang news parser!:)"
    console.log(message)
    res.send(message)
  }else{
  console.log("Got request to "+req.body.link)
  let data = get_js_rendered_page(req.body.link)
  data.then(res.send.bind(res)).catch(error => {
    console.log(error+" promise rejected!")
    res.send("Couldnt load page")
  })
  }
})
console.log("server is running!")
app.listen(8000);
