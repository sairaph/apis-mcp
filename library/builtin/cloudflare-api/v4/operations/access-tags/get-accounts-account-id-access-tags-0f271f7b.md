---
title: List tags
page_id: operation-get-accounts-account-id-access-tags-d581eb63
path: operations/access-tags
description: List tags
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/access/tags
operation_ids:
    - access-tags-list-tags
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List tags

`GET /accounts/{account_id}/access/tags`

Operation ID: `access-tags-list-tags`

List tags

## Definition

```yaml
{"operationId": "access-tags-list-tags", "summary": "List tags", "description": "List tags", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}, {"$ref": "#/components/parameters/access_page"}, {"name": "per_page", "in": "query", "schema": {"description": "Number of results per page.", "type": "integer", "default": 50, "maximum": 1000}}], "responses": {"200": {"description": "List tags response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_response_collection-15"}}}}, "4XX": {"description": "List tags response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Access tags"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.access.tags", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
