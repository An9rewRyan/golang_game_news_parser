## Welcome to the golang web scraper 

![](https://img.shields.io/badge/golang-1.17-52a7f7) ![](https://img.shields.io/badge/express_js-ffee03) ![](https://img.shields.io/badge/-selenium-ff69b4) ![](https://img.shields.io/badge/-postgresql-3294f0) ![](https://img.shields.io/badge/-docker-32c7f0) ![](https://img.shields.io/badge/-puppeteer-63b871) ![](https://img.shields.io/badge/-htmlquery-4f75ff)


>  Work in progress... *(i mean it)*

---

### How to launch: 
 1. Download docker compose on your machine
 2. Uncomment marked string in main.go file, for creating tables (comment it after first sucessfull launch)
 3. Download this repo
 4. Run (from the root directory): docker compose build 
 5. Run: docker compose up
 6. Enjoy! ;p

***
**What is used for what:**
 - **Go** - the basis of the parser, sends requests, searches for data *(+ htmlquery)*
 - **Js** - handles javascript on found pages if necessary *(+ puppetteer)*

***

**Other projects using this parser:**
 - **[Gamers Gazette](https://github.com/An9rewRyan/Gamers-Gazette)** - web platform based on this parser
 - **[Frontend on react for this service](https://github.com/An9rewRyan/Gamers-Gazette-frontend-react)** - react based frontend

***

**What resources does the parser use:**
 - **[dtf.ru](https://dtf.ru/)**
 - **[igromania.ru](https://www.igromania.ru/)**
 - **[kanobu.ru](https://kanobu.ru/videogames/)**
 - **[playground.ru](https://www.playground.ru/)**
 - **[stopgame.ru](https://stopgame.ru/)**
 - **[vgtimes.ru](https://vgtimes.ru/)**
