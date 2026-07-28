---
title: WebSearchRequest
page_id: schema-websearchrequest-a9420fec
path: schemas
source: https://docs.z.ai/openapi.json
source_type: openapi
imported_from: https://docs.z.ai/openapi.json
---

# WebSearchRequest

```yaml
{"type": "object", "properties": {"search_engine": {"type": "string", "description": "The search engine code to call.\n search-prime: Z.AI Premium Version Search Engine", "example": "search-prime", "default": "search-prime", "enum": ["search-prime"]}, "search_query": {"type": "string", "description": "The content to be searched."}, "count": {"type": "integer", "description": "The number of results to return\nFillable range: `1-50`, maximum `50` results per single search\nDefault is `10`\nSupported search engines: \n`search_pro_jina`.", "minimum": 1, "maximum": 50}, "search_domain_filter": {"type": "string", "description": "Used to limit the scope of search results and only return content from specified whitelist domains.\nWhitelist: Directly enter the domain name (e.g., `www.example.com`)\nSupported search engines: \n`search_pro_jina`"}, "search_recency_filter": {"type": "string", "description": "Search for webpages within a specified time range.\nDefault is `noLimit`\nFillable values:\n`oneDay`: within one day\n`oneWeek`: within one week\n`oneMonth`: within one month\n`oneYear`: within one year\n`noLimit`: no limit (default)\nSupported search engines: \n`search_pro_jina`", "enum": ["oneDay", "oneWeek", "oneMonth", "oneYear", "noLimit"]}, "request_id": {"type": "string", "description": "Passed by the user side, needs to be unique; used to distinguish each request, 6–64 characters. If not provided by the user side, the platform will generate one by default.", "minLength": 6, "maxLength": 64}, "user_id": {"type": "string", "description": "Unique ID for the end user, 6–128 characters. Avoid using sensitive information.", "minLength": 6, "maxLength": 128}}, "required": ["search_engine", "search_query"]}
```
