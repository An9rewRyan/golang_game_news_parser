// const writeText = require("mylib/core.io.file/write-text");
const puppeteer = require("puppeteer");
const path = require("path");
const common = require("./common.js");
const fs = require('fs');

async function launch_chrome () {
    console.log("Launching chrome")

    const launch_options = {
        args: common.minimal_args,
        headless: true,
        devtools: false,
        defaultViewport: {width: 1200, height: 1000},
        userDataDir: common.userDataDir,
        // executablePath: "./node_modules/puppeteer/.local-chromium/linux-970485/chrome-linux",
    };
    const browser = await puppeteer.launch(launch_options);
    const wsEndpoint = browser.wsEndpoint();
    console.log("Endpoint: "+wsEndpoint)
    fs.writeFile("./fnsettings.json", JSON.stringify({wsEndpoint}, null, "  "), err => {
        if (err) {
            console.error("err: "+err)
        }
        // console.log(data)
        })
    };

launch_chrome()