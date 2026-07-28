---
title: List Web Analytics site tags
page_id: operation-get-accounts-account-id-rum-site-info-site-tag-list-1f169f9d
path: operations/web-analytics
description: Returns all site tags for an account as an array of site tag strings. This endpoint returns all results without pagination.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/rum/site_info/site_tag/list
operation_ids:
    - web-analytics-list-site-tags
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Web Analytics site tags

`GET /accounts/{account_id}/rum/site_info/site_tag/list`

Operation ID: `web-analytics-list-site-tags`

Returns all site tags for an account as an array of site tag strings. This endpoint returns all results without pagination.

## Definition

```yaml
{"operationId": "web-analytics-list-site-tags", "summary": "List Web Analytics site tags", "description": "Returns all site tags for an account as an array of site tag strings. This endpoint returns all results without pagination.\n", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/rum_identifier"}}, {"name": "all", "in": "query", "description": "When true, includes sites that are disabled or paused.", "schema": {"type": "boolean", "example": false}}], "responses": {"200": {"description": "List of Web Analytics site tags.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/rum_site-tag-list-response"}}}}, "4XX": {"description": "Failure response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/rum_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Web Analytics"], "x-api-token-group": ["Account Settings Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
