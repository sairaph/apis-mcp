---
title: List Page Shield Cookies
page_id: operation-get-zones-zone-id-page-shield-cookies-bd0a4f31
path: operations/page-shield
description: Lists all cookies collected by Page Shield.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/page_shield/cookies
operation_ids:
    - page-shield-list-cookies
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Page Shield Cookies

`GET /zones/{zone_id}/page_shield/cookies`

Operation ID: `page-shield-list-cookies`

Lists all cookies collected by Page Shield.

## Definition

```yaml
{"operationId": "page-shield-list-cookies", "summary": "List Page Shield Cookies", "description": "Lists all cookies collected by Page Shield.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/page-shield_id"}}, {"name": "hosts", "in": "query", "schema": {"description": "Includes cookies that match one or more URL-encoded hostnames separated by commas.\n\nWildcards are supported at the start and end of each hostname to support starts with, ends with\nand contains. If no wildcards are used, results will be filtered by exact match\n", "type": "string", "example": "blog.cloudflare.com,www.example*,*cloudflare.com"}}, {"name": "page", "in": "query", "schema": {"description": "The current page number of the paginated results.\n\nWe additionally support a special value \"all\". When \"all\" is used, the API will return all the cookies\nwith the applied filters in a single page. This feature is best-effort and it may only work for zones with\na low number of cookies\n", "type": "string", "example": 2}}, {"name": "per_page", "in": "query", "schema": {"description": "The number of results per page.", "type": "number", "example": 100, "maximum": 100, "minimum": 1}}, {"name": "order_by", "in": "query", "schema": {"description": "The field used to sort returned cookies.", "type": "string", "example": "first_seen_at", "enum": ["first_seen_at", "last_seen_at"]}}, {"name": "direction", "in": "query", "schema": {"description": "The direction used to sort returned cookies.'", "type": "string", "example": "asc", "enum": ["asc", "desc"]}}, {"name": "page_url", "in": "query", "schema": {"description": "Includes connections that match one or more page URLs (separated by commas) where they were last seen\n\nWildcards are supported at the start and end of each page URL to support starts with, ends with\nand contains. If no wildcards are used, results will be filtered by exact match\n", "type": "string", "example": "example.com/page,*/checkout,example.com/*,*checkout*"}}, {"name": "export", "in": "query", "schema": {"description": "Export the list of cookies as a file, limited to 50000 entries.", "type": "string", "example": "csv", "enum": ["csv"]}}, {"name": "name", "in": "query", "schema": {"description": "Filters the returned cookies that match the specified name.\nWildcards are supported at the start and end to support starts with, ends with\nand contains. e.g. session*\n", "type": "string", "example": "session_id", "maxLength": 1024}}, {"name": "secure", "in": "query", "schema": {"description": "Filters the returned cookies that are set with Secure", "type": "boolean", "example": true}}, {"name": "http_only", "in": "query", "schema": {"description": "Filters the returned cookies that are set with HttpOnly", "type": "boolean", "example": true}}, {"name": "same_site", "in": "query", "schema": {"description": "Filters the returned cookies that match the specified same_site attribute", "type": "string", "example": "strict", "enum": ["lax", "strict", "none"]}}, {"name": "type", "in": "query", "schema": {"description": "Filters the returned cookies that match the specified type attribute", "type": "string", "example": "first_party", "enum": ["first_party", "unknown"]}}, {"name": "path", "in": "query", "schema": {"description": "Filters the returned cookies that match the specified path attribute", "type": "string", "example": "/", "maxLength": 1024}}, {"name": "domain", "in": "query", "schema": {"description": "Filters the returned cookies that match the specified domain attribute", "type": "string", "example": "example.com", "maxLength": 1024}}], "responses": {"200": {"description": "List Page Shield cookies response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/page-shield_list-zone-cookies-response"}}}}, "4XX": {"description": "List Page Shield cookies response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/page-shield_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Page Shield"], "x-api-token-group": ["Page Shield", "Domain Page Shield Read", "Domain Page Shield", "Page Shield Read", "Zone Settings Write", "Zone Settings Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": false}}
```
