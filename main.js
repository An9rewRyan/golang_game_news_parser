// console.log(link)
(async (link) => {
    const browser = await puppeteer.launch();
    const page = await browser.newPage();

    await page.goto(link);
    let bodyHTML = await page.evaluate(() => document.documentElement.outerHTML);
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
})()
