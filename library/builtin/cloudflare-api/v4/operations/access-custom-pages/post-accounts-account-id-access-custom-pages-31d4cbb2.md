---
title: Create a custom page
page_id: operation-post-accounts-account-id-access-custom-pages-f898310e
path: operations/access-custom-pages
description: Create a custom page
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/access/custom_pages
operation_ids:
    - access-custom-pages-create-a-custom-page
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a custom page

`POST /accounts/{account_id}/access/custom_pages`

Operation ID: `access-custom-pages-create-a-custom-page`

Create a custom page

## Definition

```yaml
{"operationId": "access-custom-pages-create-a-custom-page", "summary": "Create a custom page", "description": "Create a custom page", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}], "requestBody": {"content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_custom_page"}}}}, "responses": {"201": {"description": "Create a custom page response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_single_response_without_html"}}}}, "4XX": {"description": "Create a custom page response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Access custom pages"], "x-api-token-group": ["Access: Custom Pages Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.access.custom-pages", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
