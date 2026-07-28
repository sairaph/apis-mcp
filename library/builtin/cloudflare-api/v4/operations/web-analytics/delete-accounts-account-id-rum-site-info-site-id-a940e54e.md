---
title: Delete a Web Analytics site
page_id: operation-delete-accounts-account-id-rum-site-info-site-id-b7cd7137
path: operations/web-analytics
description: Deletes an existing Web Analytics site.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/rum/site_info/{site_id}
operation_ids:
    - web-analytics-delete-site
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete a Web Analytics site

`DELETE /accounts/{account_id}/rum/site_info/{site_id}`

Operation ID: `web-analytics-delete-site`

Deletes an existing Web Analytics site.

## Definition

```yaml
{"operationId": "web-analytics-delete-site", "summary": "Delete a Web Analytics site", "description": "Deletes an existing Web Analytics site.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/rum_identifier"}}, {"name": "site_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/rum_identifier"}}], "responses": {"200": {"description": "Deleted Web Analytics site identifier.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/rum_site-tag-response-single"}}}}, "4XX": {"description": "Failure response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/rum_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Web Analytics"], "x-api-token-group": ["Account Settings Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "rum.site-info", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
