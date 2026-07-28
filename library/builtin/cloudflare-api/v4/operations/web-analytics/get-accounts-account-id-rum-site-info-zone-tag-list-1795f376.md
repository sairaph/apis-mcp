---
title: List Web Analytics zone tags
page_id: operation-get-accounts-account-id-rum-site-info-zone-tag-list-3a3e77ea
path: operations/web-analytics
description: Returns all zone tags associated with Web Analytics sites for an account. These can be used to filter or exclude zones in the zone selection input. This endpoint returns all results without pagination.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/rum/site_info/zone_tag/list
operation_ids:
    - web-analytics-list-zone-tags
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Web Analytics zone tags

`GET /accounts/{account_id}/rum/site_info/zone_tag/list`

Operation ID: `web-analytics-list-zone-tags`

Returns all zone tags associated with Web Analytics sites for an account. These can be used to filter or exclude zones in the zone selection input. This endpoint returns all results without pagination.

## Definition

```yaml
{"operationId": "web-analytics-list-zone-tags", "summary": "List Web Analytics zone tags", "description": "Returns all zone tags associated with Web Analytics sites for an account. These can be used to filter or exclude zones in the zone selection input. This endpoint returns all results without pagination.\n", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/rum_identifier"}}], "responses": {"200": {"description": "List of zone tags.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/rum_zone-tag-list-response"}}}}, "4XX": {"description": "Failure response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/rum_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Web Analytics"], "x-api-token-group": ["Account Settings Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
