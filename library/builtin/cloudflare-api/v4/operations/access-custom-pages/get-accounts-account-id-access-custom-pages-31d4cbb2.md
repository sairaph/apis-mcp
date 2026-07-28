---
title: List custom pages
page_id: operation-get-accounts-account-id-access-custom-pages-b69cddcf
path: operations/access-custom-pages
description: List custom pages
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/access/custom_pages
operation_ids:
    - access-custom-pages-list-custom-pages
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List custom pages

`GET /accounts/{account_id}/access/custom_pages`

Operation ID: `access-custom-pages-list-custom-pages`

List custom pages

## Definition

```yaml
{"operationId": "access-custom-pages-list-custom-pages", "summary": "List custom pages", "description": "List custom pages", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}, {"$ref": "#/components/parameters/access_page"}, {"name": "per_page", "in": "query", "schema": {"description": "Number of results per page.", "type": "integer", "default": 50, "maximum": 1000}}], "responses": {"200": {"description": "List custom pages response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_response_collection-10"}}}}, "4XX": {"description": "List custom pages response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Access custom pages"], "x-api-token-group": ["Access: Custom Pages Write", "Access: Custom Pages Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.access.custom-pages", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
