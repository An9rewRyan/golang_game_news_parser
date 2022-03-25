const puppeteer = require("puppeteer");
const fs = require("fs");
const express = require('express')
const app = express()

//this huge thing has necessary args for increasing performance 
const minimal_args = [
  '--autoplay-policy=user-gesture-required',
  '--disable-background-networking',
  '--disable-background-timer-throttling',
  '--disable-backgrounding-occluded-windows',
  '--disable-breakpad',
  '--disable-client-side-phishing-detection',
  '--disable-component-update',
  '--disable-default-apps',
  '--disable-dev-shm-usage',
  '--disable-domain-reliability',
  '--disable-extensions',
  '--disable-features=AudioServiceOutOfProcess',
  '--disable-hang-monitor',
  '--disable-ipc-flooding-protection',
  '--disable-notifications',
  '--disable-offer-store-unmasked-wallet-cards',
  '--disable-popup-blocking',
  '--disable-print-preview',
  '--disable-prompt-on-repost',
  '--disable-renderer-backgrounding',
  '--disable-setuid-sandbox',
  '--disable-speech-api',
  '--disable-sync',
  '--hide-scrollbars',
  '--ignore-gpu-blacklist',
  '--metrics-recording-only',
  '--mute-audio',
  '--no-default-browser-check',
  '--no-first-run',
  '--no-pings',
  '--no-sandbox',
  '--no-zygote',
  '--password-store=basic',
  '--use-gl=swiftshader',
  '--use-mock-keychain',
];

const blocked_domains = [
  'googlesyndication.com',
  'adservice.google.com',
];

async function get_js_rendered_page (link) {
    console.log(link)
    const browser = await puppeteer.launch({
      userDataDir: './cache',
      args: minimal_args
    });
    page = await browser.newPage();
    await page.setDefaultNavigationTimeout(0); 
    await page.goto(link, {
      waitUntil: 'load',
    });

    let bodyHTML = await page.evaluate(() => document.documentElement.outerHTML);
    
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
  console.log("Got request to "+req.body.link)
  let data = get_js_rendered_page(req.body.link)
  data.then(res.send.bind(res)).catch(error => {console.log(error+" promise rejected!")})
})
console.log("server is running!")
app.listen(8000);
