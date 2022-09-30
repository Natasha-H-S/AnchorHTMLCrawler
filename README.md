<h1 align="center">WebCrawler</h1>

<p align="center">
Parses HTML to find anchor tags with hrefs and displays them in a list.

</p>

## Introduction

This application parses HTML and finds anchor tags with `href` attributes using html parser and displays the results. 


## How to run

```bash
go build . 


URL="https://en.wikipedia.org/wiki/Ada_Lovelace" go run . 
```
## Environment Variables

| Name | Type | Default Value | Mandatory | Description |
|------|------|---------------|-----------|-------------|
| URL  | string | N/A         |   Yes     |  URL to webcrawl |

