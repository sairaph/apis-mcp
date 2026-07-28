---
title: List DLS prefix bindings for an account
page_id: operation-get-accounts-account-id-dls-regional-services-prefix-bindings-99fbe8c0
path: operations/prefix-bindings
description: List the BYOIP prefix bindings configured for an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/dls/regional_services/prefix_bindings
operation_ids:
    - publicListPrefixBindings
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List DLS prefix bindings for an account

`GET /accounts/{account_id}/dls/regional_services/prefix_bindings`

Operation ID: `publicListPrefixBindings`

List the BYOIP prefix bindings configured for an account.

## Definition

```yaml
{"operationId": "publicListPrefixBindings", "summary": "List DLS prefix bindings for an account", "description": "List the BYOIP prefix bindings configured for an account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/dls_account_identifier"}}, {"name": "cursor", "in": "query", "description": "Opaque token for cursor-based pagination. Omit for the first page. Pass the value from a previous response to fetch the next page.", "schema": {"type": "string"}}, {"name": "per_page", "in": "query", "schema": {"type": "integer", "default": 25, "maximum": 100, "minimum": 1}}], "responses": {"200": {"description": "List of bindings.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dls_PrefixBindingPaginatedListResponse"}}}}, "400": {"description": "Bad request.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dls_bad_response"}}}}, "401": {"description": "Unauthorized.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dls_bad_response"}}}}, "403": {"description": "Forbidden.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dls_bad_response"}}}}, "500": {"description": "Internal server error.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dls_bad_response"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Prefix Bindings"], "x-api-token-group": ["DLS: Read", "DLS: Write"]}
```
