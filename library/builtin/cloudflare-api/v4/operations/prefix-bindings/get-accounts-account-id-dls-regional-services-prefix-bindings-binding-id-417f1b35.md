---
title: Get a DLS prefix binding
page_id: operation-get-accounts-account-id-dls-regional-services-prefix-bindings-binding-id-0b6e828c
path: operations/prefix-bindings
description: Retrieve a single BYOIP prefix binding by ID.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/dls/regional_services/prefix_bindings/{binding_id}
operation_ids:
    - publicGetPrefixBinding
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get a DLS prefix binding

`GET /accounts/{account_id}/dls/regional_services/prefix_bindings/{binding_id}`

Operation ID: `publicGetPrefixBinding`

Retrieve a single BYOIP prefix binding by ID.

## Definition

```yaml
{"operationId": "publicGetPrefixBinding", "summary": "Get a DLS prefix binding", "description": "Retrieve a single BYOIP prefix binding by ID.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/dls_account_identifier"}}, {"name": "binding_id", "in": "path", "required": true, "schema": {"description": "Unique identifier for the prefix binding.", "type": "string", "example": "a1b2c3d4-e5f6-7890-abcd-ef1234567890"}}], "responses": {"200": {"description": "Binding found.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dls_PrefixBindingResponse"}}}}, "401": {"description": "Unauthorized.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dls_bad_response"}}}}, "403": {"description": "Forbidden.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dls_bad_response"}}}}, "404": {"description": "Not found.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dls_bad_response"}}}}, "500": {"description": "Internal server error.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dls_bad_response"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Prefix Bindings"], "x-api-token-group": ["DLS: Read", "DLS: Write"]}
```
