---
title: Create a Web Analytics site
page_id: operation-post-accounts-account-id-rum-site-info-106f1c1f
path: operations/web-analytics
description: Creates a new Web Analytics site.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/rum/site_info
operation_ids:
    - web-analytics-create-site
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a Web Analytics site

`POST /accounts/{account_id}/rum/site_info`

Operation ID: `web-analytics-create-site`

Creates a new Web Analytics site.

## Definition

```yaml
{"operationId": "web-analytics-create-site", "summary": "Create a Web Analytics site", "description": "Creates a new Web Analytics site.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/rum_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/rum_create-site-request"}}}}, "responses": {"200": {"description": "Created Web Analytics site.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/rum_site-response-single"}}}}, "4XX": {"description": "Failure response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/rum_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Web Analytics"], "x-api-token-group": ["Account Settings Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "rum.site-info", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
