---
title: Get a Web Analytics site
page_id: operation-get-accounts-account-id-rum-site-info-site-id-bbd98112
path: operations/web-analytics
description: Retrieves a Web Analytics site.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/rum/site_info/{site_id}
operation_ids:
    - web-analytics-get-site
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get a Web Analytics site

`GET /accounts/{account_id}/rum/site_info/{site_id}`

Operation ID: `web-analytics-get-site`

Retrieves a Web Analytics site.

## Definition

```yaml
{"operationId": "web-analytics-get-site", "summary": "Get a Web Analytics site", "description": "Retrieves a Web Analytics site.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/rum_identifier"}}, {"name": "site_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/rum_identifier"}}], "responses": {"200": {"description": "Web Analytics site.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/rum_site-response-single"}}}}, "4XX": {"description": "Failure response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/rum_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Web Analytics"], "x-api-token-group": ["Account Settings Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "rum.site-info", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
