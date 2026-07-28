---
title: List Web Analytics sites
page_id: operation-get-accounts-account-id-rum-site-info-list-50981b15
path: operations/web-analytics
description: Lists all Web Analytics sites of an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/rum/site_info/list
operation_ids:
    - web-analytics-list-sites
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Web Analytics sites

`GET /accounts/{account_id}/rum/site_info/list`

Operation ID: `web-analytics-list-sites`

Lists all Web Analytics sites of an account.

## Definition

```yaml
{"operationId": "web-analytics-list-sites", "summary": "List Web Analytics sites", "description": "Lists all Web Analytics sites of an account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/rum_identifier"}}, {"name": "per_page", "in": "query", "schema": {"$ref": "#/components/schemas/rum_per_page"}}, {"name": "page", "in": "query", "schema": {"$ref": "#/components/schemas/rum_page"}}, {"name": "order_by", "in": "query", "schema": {"$ref": "#/components/schemas/rum_order_by"}}], "responses": {"200": {"description": "List of Web Analytics sites.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/rum_sites-response-collection"}}}}, "4XX": {"description": "Failure response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/rum_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Web Analytics"], "x-api-token-group": ["Account Settings Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "rum.site-info", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
