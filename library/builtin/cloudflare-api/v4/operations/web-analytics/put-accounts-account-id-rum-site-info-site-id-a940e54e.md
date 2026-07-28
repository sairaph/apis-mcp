---
title: Update a Web Analytics site
page_id: operation-put-accounts-account-id-rum-site-info-site-id-cb9ab132
path: operations/web-analytics
description: Updates an existing Web Analytics site.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/rum/site_info/{site_id}
operation_ids:
    - web-analytics-update-site
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update a Web Analytics site

`PUT /accounts/{account_id}/rum/site_info/{site_id}`

Operation ID: `web-analytics-update-site`

Updates an existing Web Analytics site.

## Definition

```yaml
{"operationId": "web-analytics-update-site", "summary": "Update a Web Analytics site", "description": "Updates an existing Web Analytics site.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/rum_identifier"}}, {"name": "site_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/rum_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/rum_update-site-request"}}}}, "responses": {"200": {"description": "Updated Web Analytics site.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/rum_site-response-single"}}}}, "4XX": {"description": "Failure response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/rum_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Web Analytics"], "x-api-token-group": ["Account Settings Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "rum.site-info", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
