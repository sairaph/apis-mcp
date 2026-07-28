---
title: Query analytics summary
page_id: operation-post-accounts-account-id-analytics-query-dataset-summary-e56d2848
path: operations/art-analytics
description: Returns aggregate summary stats for a dataset. Includes current-period and previous-period totals for trend comparison.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/analytics/query/{dataset}/summary
operation_ids:
    - art-analytics-query-summary
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Query analytics summary

`POST /accounts/{account_id}/analytics/query/{dataset}/summary`

Operation ID: `art-analytics-query-summary`

Returns aggregate summary stats for a dataset. Includes current-period and previous-period totals for trend comparison.

## Definition

```yaml
{"operationId": "art-analytics-query-summary", "summary": "Query analytics summary", "description": "Returns aggregate summary stats for a dataset. Includes current-period and previous-period totals for trend comparison.\n", "parameters": [{"$ref": "#/components/parameters/art_account_id"}, {"name": "dataset", "in": "path", "description": "Dataset name to query. Examples: `access-logins`, `gateway-http`, `gateway-dns`, `gateway-http`, `shadow-it`.\n", "required": true, "schema": {"type": "string", "example": "access-logins"}}], "requestBody": {"required": true, "content": {"application/json": {"examples": {"access_summary": {"summary": "Access login summary for a week", "value": {"filters": [], "from": "2024-11-01T00:00:00Z", "groupBy": [], "stats": ["attemptsTotal"], "to": "2024-11-08T00:00:00Z"}}}, "schema": {"$ref": "#/components/schemas/art_QuerySummary"}}}}, "responses": {"200": {"description": "Summary query result.", "content": {"application/json": {"examples": {"success": {"summary": "Successful summary result", "value": {"errors": [], "messages": [{"code": 1000, "message": "API in beta: expect breaking changes."}], "result": {"currentTotal": [{"attemptsTotal": 48291}], "previousTotal": [{"attemptsTotal": 41033}]}, "success": true}}}, "schema": {"$ref": "#/components/schemas/art_SummaryResponse"}}}}, "400": {"$ref": "#/components/responses/art_BadRequest"}, "401": {"$ref": "#/components/responses/art_Unauthorized"}, "403": {"$ref": "#/components/responses/art_Forbidden"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["ART Analytics"], "x-api-token-group": ["Zero Trust Read"]}
```
