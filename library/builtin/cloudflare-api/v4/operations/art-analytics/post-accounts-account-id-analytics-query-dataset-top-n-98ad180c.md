---
title: Query analytics top-N
page_id: operation-post-accounts-account-id-analytics-query-dataset-top-n-c6ec21f2
path: operations/art-analytics
description: Returns the top N results for a dataset by a specified stat. Includes an array of result rows, each containing the requested stats and group-by dimensions.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/analytics/query/{dataset}/top-n
operation_ids:
    - art-analytics-query-top-n
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Query analytics top-N

`POST /accounts/{account_id}/analytics/query/{dataset}/top-n`

Operation ID: `art-analytics-query-top-n`

Returns the top N results for a dataset by a specified stat. Includes an array of result rows, each containing the requested stats and group-by dimensions.

## Definition

```yaml
{"operationId": "art-analytics-query-top-n", "summary": "Query analytics top-N", "description": "Returns the top N results for a dataset by a specified stat. Includes an array of result rows, each containing the requested stats and group-by dimensions.\n", "parameters": [{"$ref": "#/components/parameters/art_account_id"}, {"name": "dataset", "in": "path", "description": "Dataset name to query. Examples: `access-logins`, `gateway-http`, `gateway-dns`, `gateway-http`, `shadow-it`.\n", "required": true, "schema": {"type": "string", "example": "gateway-http"}}], "requestBody": {"required": true, "content": {"application/json": {"examples": {"top_apps_by_bytes": {"summary": "Top 10 shadow IT apps by bytes", "value": {"filters": [], "from": "2024-11-05T00:00:00Z", "groupBy": ["appName", "appCategory"], "n": 10, "orderBy": "bytesTotal", "stats": ["bytesTotal", "requestsTotal"], "to": "2024-11-06T00:00:00Z"}}, "top_countries_access": {"summary": "Top 5 countries by login attempts", "value": {"filters": [{"name": "allowed", "op": "eq", "values": [false]}], "from": "2024-11-01T00:00:00Z", "groupBy": ["country"], "n": 5, "orderBy": "attemptsTotal", "stats": ["attemptsTotal"], "to": "2024-11-08T00:00:00Z"}}}, "schema": {"$ref": "#/components/schemas/art_QueryTopN"}}}}, "responses": {"200": {"description": "Top-N query result.", "content": {"application/json": {"examples": {"success": {"summary": "Successful top-N result", "value": {"errors": [], "messages": [{"code": 1000, "message": "API in beta: expect breaking changes."}], "result": [{"appCategory": "Collaboration", "appName": "Slack", "bytesTotal": 10485760, "requestsTotal": 1024}, {"appCategory": "File Storage", "appName": "Dropbox", "bytesTotal": 5242880, "requestsTotal": 512}], "success": true}}}, "schema": {"$ref": "#/components/schemas/art_TopNResponse"}}}}, "400": {"$ref": "#/components/responses/art_BadRequest"}, "401": {"$ref": "#/components/responses/art_Unauthorized"}, "403": {"$ref": "#/components/responses/art_Forbidden"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["ART Analytics"], "x-api-token-group": ["Zero Trust Read"]}
```
