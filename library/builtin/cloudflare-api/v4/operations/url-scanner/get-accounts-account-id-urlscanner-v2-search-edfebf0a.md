---
title: Search URL scans
page_id: operation-get-accounts-account-id-urlscanner-v2-search-2a30661e
path: operations/url-scanner
description: 'Use a subset of ElasticSearch Query syntax to filter scans. Some example queries:<br/> <br/>- ''path:"/bundles/jquery.js"'': Searches for scans who requested resources with the given path.<br/>- ''page.asn:AS24940 AND hash:xxx'': Websites hosted in AS24940 where a resource with the given hash was downloaded.<br/>- ''page.domain:microsoft* AND verdicts.malicious:true AND NOT page.domain:microsoft.com'': malicious scans whose hostname starts with "microsoft".<br/>- ''apikey:me AND date:[2025-01 TO 2025-02]'': my scans from 2025 January to 2025 February.'
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/urlscanner/v2/search
operation_ids:
    - urlscanner-search-scans-v2
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Search URL scans

`GET /accounts/{account_id}/urlscanner/v2/search`

Operation ID: `urlscanner-search-scans-v2`

Use a subset of ElasticSearch Query syntax to filter scans. Some example queries:<br/> <br/>- 'path:"/bundles/jquery.js"': Searches for scans who requested resources with the given path.<br/>- 'page.asn:AS24940 AND hash:xxx': Websites hosted in AS24940 where a resource with the given hash was downloaded.<br/>- 'page.domain:microsoft* AND verdicts.malicious:true AND NOT page.domain:microsoft.com': malicious scans whose hostname starts with "microsoft".<br/>- 'apikey:me AND date:[2025-01 TO 2025-02]': my scans from 2025 January to 2025 February.

## Definition

```yaml
{"operationId": "urlscanner-search-scans-v2", "summary": "Search URL scans", "description": "Use a subset of ElasticSearch Query syntax to filter scans. Some example queries:<br/> <br/>- 'path:\"/bundles/jquery.js\"': Searches for scans who requested resources with the given path.<br/>- 'page.asn:AS24940 AND hash:xxx': Websites hosted in AS24940 where a resource with the given hash was downloaded.<br/>- 'page.domain:microsoft* AND verdicts.malicious:true AND NOT page.domain:microsoft.com': malicious scans whose hostname starts with \"microsoft\".<br/>- 'apikey:me AND date:[2025-01 TO 2025-02]': my scans from 2025 January to 2025 February.", "parameters": [{"name": "account_id", "in": "path", "description": "Account ID.", "required": true, "schema": {"description": "Account ID.", "type": "string"}}, {"name": "size", "in": "query", "description": "Limit the number of objects in the response.", "schema": {"description": "Limit the number of objects in the response.", "type": "integer", "example": 100}}, {"name": "q", "in": "query", "description": "Filter scans", "schema": {"description": "Filter scans", "type": "string"}}], "responses": {"200": {"description": "Search results", "content": {"application/json": {"schema": {"type": "object", "properties": {"results": {"type": "array", "items": {"properties": {"_id": {"type": "string", "example": "9626f773-9ffb-4cfb-89d3-30b120fc8011"}, "page": {"type": "object", "properties": {"asn": {"type": "string", "example": "AS15133"}, "country": {"type": "string", "example": "US"}, "ip": {"type": "string", "example": "93.184.215.14"}, "url": {"type": "string", "example": "https://example.com"}}, "required": ["country", "ip", "url", "asn"]}, "result": {"type": "string", "example": "https://radar.clouflare.com/scan/9626f773-9ffb-4cfb-89d3-30b120fc8011"}, "stats": {"type": "object", "properties": {"dataLength": {"type": "number", "example": 2512}, "requests": {"type": "number", "example": 2}, "uniqCountries": {"type": "number", "example": 1}, "uniqIPs": {"type": "number", "example": 1}}, "required": ["uniqIPs", "uniqCountries", "dataLength", "requests"]}, "task": {"type": "object", "properties": {"time": {"type": "string", "example": "2024-09-30T23:54:02.881000+00:00"}, "url": {"type": "string", "example": "https://example.com"}, "uuid": {"type": "string", "example": "9626f773-9ffb-4cfb-89d3-30b120fc8011"}, "visibility": {"type": "string", "example": "public"}}, "required": ["visibility", "time", "uuid", "url"]}, "verdicts": {"type": "object", "properties": {"malicious": {"type": "boolean"}}, "required": ["malicious"]}}, "required": ["task", "stats", "page", "verdicts", "_id", "result"], "type": "object"}}}, "required": ["results"]}}}}, "400": {"description": "Invalid input.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"properties": {"detail": {"type": "string"}, "status": {"description": "Status code.", "type": "integer", "example": 400}, "title": {"type": "string", "example": "Invalid url"}}, "required": ["title", "detail", "status"], "type": "object"}}, "message": {"type": "string"}, "status": {"description": "Status code.", "type": "integer", "example": 400}}, "required": ["message", "status", "errors"]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["URL Scanner"], "x-api-token-group": ["URL Scanner Write", "URL Scanner Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "url-scanner.scans", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
