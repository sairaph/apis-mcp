---
title: List Page Shield scripts
page_id: operation-get-zones-zone-id-page-shield-scripts-7984d4b2
path: operations/page-shield
description: Lists all scripts detected by Page Shield.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/page_shield/scripts
operation_ids:
    - page-shield-list-scripts
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Page Shield scripts

`GET /zones/{zone_id}/page_shield/scripts`

Operation ID: `page-shield-list-scripts`

Lists all scripts detected by Page Shield.

## Definition

```yaml
{"operationId": "page-shield-list-scripts", "summary": "List Page Shield scripts", "description": "Lists all scripts detected by Page Shield.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/page-shield_id"}}, {"name": "exclude_urls", "in": "query", "schema": {"description": "Excludes scripts whose URL contains one of the URL-encoded URLs separated by commas.\n", "type": "string", "example": "blog.cloudflare.com,www.example"}}, {"name": "urls", "in": "query", "schema": {"description": "Includes scripts whose URL contain one or more URL-encoded URLs separated by commas.\n", "type": "string", "example": "blog.cloudflare.com,www.example"}}, {"name": "hosts", "in": "query", "schema": {"description": "Includes scripts that match one or more URL-encoded hostnames separated by commas.\n\nWildcards are supported at the start and end of each hostname to support starts with, ends with\nand contains. If no wildcards are used, results will be filtered by exact match\n", "type": "string", "example": "blog.cloudflare.com,www.example*,*cloudflare.com"}}, {"name": "page", "in": "query", "schema": {"description": "The current page number of the paginated results.\n\nWe additionally support a special value \"all\". When \"all\" is used, the API will return all the scripts\nwith the applied filters in a single page. This feature is best-effort and it may only work for zones with\na low number of scripts\n", "type": "string", "example": 2}}, {"name": "per_page", "in": "query", "schema": {"description": "The number of results per page.", "type": "number", "example": 100, "maximum": 100, "minimum": 1}}, {"name": "order_by", "in": "query", "schema": {"description": "The field used to sort returned scripts.", "type": "string", "example": "first_seen_at", "enum": ["first_seen_at", "last_seen_at"]}}, {"name": "direction", "in": "query", "schema": {"description": "The direction used to sort returned scripts.", "type": "string", "example": "asc", "enum": ["asc", "desc"]}}, {"name": "prioritize_malicious", "in": "query", "schema": {"description": "When true, malicious scripts appear first in the returned scripts.", "type": "boolean", "example": true}}, {"name": "exclude_cdn_cgi", "in": "query", "schema": {"description": "When true, excludes scripts seen in a `/cdn-cgi` path from the returned scripts. The default value is true.", "type": "boolean", "example": true, "default": true}}, {"name": "exclude_duplicates", "in": "query", "schema": {"description": "When true, excludes duplicate scripts. We consider a script duplicate of another if their javascript\ncontent matches and they share the same url host and zone hostname. In such case, we return the most\nrecent script for the URL host and zone hostname combination.\n", "type": "boolean", "example": true, "default": true}}, {"name": "status", "in": "query", "schema": {"description": "Filters the returned scripts using a comma-separated list of scripts statuses. Accepted values: `active`, `infrequent`, and `inactive`. The default value is `active`.", "type": "string", "example": "active,inactive"}}, {"name": "page_url", "in": "query", "schema": {"description": "Includes scripts that match one or more page URLs (separated by commas) where they were last seen\n\nWildcards are supported at the start and end of each page URL to support starts with, ends with\nand contains. If no wildcards are used, results will be filtered by exact match\n", "type": "string", "example": "example.com/page,*/checkout,example.com/*,*checkout*"}}, {"name": "export", "in": "query", "schema": {"description": "Export the list of scripts as a file, limited to 50000 entries.", "type": "string", "example": "csv", "enum": ["csv"]}}], "responses": {"200": {"description": "List Page Shield scripts response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/page-shield_list-zone-scripts-response"}}}}, "4XX": {"description": "List Page Shield scripts response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/page-shield_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Page Shield"], "x-api-token-group": ["Page Shield", "Domain Page Shield Read", "Domain Page Shield", "Page Shield Read", "Zone Settings Write", "Zone Settings Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": true}}
```
