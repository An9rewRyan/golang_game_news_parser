const { json } = require("body-parser");
const puppeteer = require("puppeteer");
const exec = require('child_process').exec;
// const delay = require("mylib/promise/delay");
const fs = require('fs').promises;
// const readText = require("mylib/core.io.file/read-text");
const common = require("./common");

// let browser = null;
// const delay = ms => new Promise(res => setTimeout(res, ms));

async function launch () {
  console.log("spawning!")
  const child = exec('node chrome_launcher.js',
  (error, stdout, stderr) => {
      console.log(`stdout: ${stdout}`);
      console.log(`stderr: ${stderr}`);
      if (error !== null) {
          console.log(`exec error: ${error}`);
      }
  });
  await new Promise(resolve => setTimeout(resolve, 4000)) //lets wait 10 seconds, just in case
}

async function getSettings() {
    let data = await fs.readFile("./fnsettings.json", 'utf8' , (err, data) => {
        if (err) {
          console.error("err: "+err)
          return null
        }
        console.log("data: "+data)
        data = JSON.parse(data)
        console.log("rdata: "+data)
        return data
    })
    return data;
}


async function connect () {
//   if (browser) return browser;
  // await launch()
  let settings = await getSettings()
  settings = JSON.parse(settings)
//   console.log(settings)
  console.log(typeof(settings), settings.wsEndpoint)
  settings = settings.wsEndpoint
  if (!settings) {
    await launch();
    settings = await getSettings();
    settings = JSON.parse(settings)
    settings = settings.wsEndpoint
  }
  console.log("Settings: "+settings)
  try {
    console.log("connecting to browser")
    browser = await puppeteer.connect({browserWSEndpoint: settings});
  } catch (e) {
    console.log("Error!: "+e)
    const err = e.error || e;
    if (err.code === "ECONNREFUSED") {
      console.log("connection refused");
      await launch();
      settings = await getSettings();
      settings = JSON.parse(settings)
      settings = settings.wsEndpoint
      try{
      browser = await puppeteer.connect({browserWSEndpoint: settings});
      }catch (e) {
      console.log("Error: "+e)
      }
    }
    // console.log("Errorn: "+e)
  }
  console.log("browser connected!")
  return browser;
}


module.exports = connect;