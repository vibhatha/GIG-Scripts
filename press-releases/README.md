1. Given a website Url crawl for all the images . 
2. Download and Parse the image files and extract the text content.
3. Use Stanford NER library to identify Named Entities in extracted text
4. Save to GIG API

## How to Run:
    1. set category var according the source category. eg. (Tenders, Gazettes, etc.)
    2. go run press_release_crawler.go "https://www.dgi.gov.lk/news/press-releases-sri-lanka"