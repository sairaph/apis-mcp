---
title: Get a custom page
page_id: operation-get-accounts-account-id-access-custom-pages-custom-page-id-0ef00346
path: operations/access-custom-pages
description: Fetches a custom page and also returns its HTML.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/access/custom_pages/{custom_page_id}
operation_ids:
    - access-custom-pages-get-a-custom-page
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get a custom page

`GET /accounts/{account_id}/access/custom_pages/{custom_page_id}`

Operation ID: `access-custom-pages-get-a-custom-page`

Fetches a custom page and also returns its HTML.

## Definition

```yaml
{"operationId": "access-custom-pages-get-a-custom-page", "summary": "Get a custom page", "description": "Fetches a custom page and also returns its HTML.", "parameters": [{"name": "custom_page_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_uuid"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}], "responses": {"200": {"description": "Get a custom page response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_single_response-11"}}}}, "4XX": {"description": "Get a custom page response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Access custom pages"], "x-api-token-group": ["Access: Custom Pages Write", "Access: Custom Pages Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.access.custom-pages", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
