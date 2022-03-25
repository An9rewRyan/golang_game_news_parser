## Welcome to the golang web scraper 

![](https://img.shields.io/badge/golang-1.17-52a7f7) ![](https://img.shields.io/badge/express_js-ffee03) ![](https://img.shields.io/badge/-selenium-ff69b4) ![](https://img.shields.io/badge/-postgresql-3294f0) ![](https://img.shields.io/badge/-docker-32c7f0) ![](https://img.shields.io/badge/-puppeteer-63b871) ![](https://img.shields.io/badge/-htmlquery-4f75ff)


>  Work in progress... *(i mean it)*

---

### How to launch: 
 1. Start node server and installing puppeteer
 2. Launch local psql server and changing dbConnStr in config.go by your setting, create article table by article struct
 3. Launch parser.exe

***
**What is used for what:**
 - **Go** - the basis of the parser, sends requests, searches for data *(+ htmlquery)*
 - **Js** - handles javascript on found pages if necessary *(+ puppetteer)*

***

**Other projects using this parser:**
 - **[Gamers Gazette](https://github.com/authoraytee/gamers_gazette)** - web platform based on this parser
 - Maybe else (...)

---

Contact me if you have some questions or suggestions via:
 - Telegram: **[@Michael_J_Goldberg](https://t.me/Michael_J_Goldberg)**
 - Vk - **[vk.com/mj_the_reviewer](https://vk.com/mj_the_reviewer)**
 - Discord - **[YUUJIRO HANMA#6379](https://discordapp.com/users/389483338865311745/)**

***

**What resources does the parser use:**
 - **[dtf.ru](https://dtf.ru/)**
 - **[igromania.ru](https://www.igromania.ru/)**
 - **[kanobu.ru](https://kanobu.ru/videogames/)**
 - **[playground.ru](https://www.playground.ru/)**
 - **[stopgame.ru](https://stopgame.ru/)**
 - **[vgtimes.ru](https://vgtimes.ru/)**
