---
title: Update a custom page
page_id: operation-put-accounts-account-id-access-custom-pages-custom-page-id-de510f04
path: operations/access-custom-pages
description: Update a custom page
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/access/custom_pages/{custom_page_id}
operation_ids:
    - access-custom-pages-update-a-custom-page
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update a custom page

`PUT /accounts/{account_id}/access/custom_pages/{custom_page_id}`

Operation ID: `access-custom-pages-update-a-custom-page`

Update a custom page

## Definition

```yaml
{"operationId": "access-custom-pages-update-a-custom-page", "summary": "Update a custom page", "description": "Update a custom page", "parameters": [{"name": "custom_page_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_uuid"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}], "requestBody": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_custom_page"}}}}, "responses": {"200": {"description": "Update a custom page response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_single_response_without_html"}}}}, "4XX": {"description": "Update a custom page response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Access custom pages"], "x-api-token-group": ["Access: Custom Pages Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.access.custom-pages", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
