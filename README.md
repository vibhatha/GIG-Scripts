# Documentation of Scripts
Include Code Snippets that can be used to crawl and feed data into GIG system
## Configurations:
    ApiUrl              string      GIG Server URL 
    NERServerUrl        string      Entity Recognition Service
    NormalizeServer     string      Entity Name Normalization Service
## Crawlers:
* [Crawler Documentation](wikipedia/README.md)
* [PDF Crawler Documentation](../GIG-SDK/pdf_crawler/README.md)
* [Wiki API Crawler Documentation](wikipedia/wiki_api_crawler/README.md)
* [Wiki Web Crawler Documentation](wikipedia/wiki_web_crawler/README.md)
## Entity Handlers:
    1. AddEntitiesAsLinks -> Add list of related entities to a given entity
    2. AddEntityAsAttribute -> Add entity as an attribute to a given entity
    3. AddEntityAsLink -> Add entity as an link to a given entity
    4. CreateEntities -> Create a list of new entities and save to GIG
    5. CreateEntity -> Create a new entity and save to GIG
    6. UploadImage -> Upload image to GIG server

## Importers:
* [eTender Documentation](tenders/etender/README.md)
## Parsers:
    1. ParsePdf -> return the string content of a given PDF file