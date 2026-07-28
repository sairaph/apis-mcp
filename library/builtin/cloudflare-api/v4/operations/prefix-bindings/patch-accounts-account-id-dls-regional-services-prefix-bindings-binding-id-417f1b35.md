---
title: Update a DLS prefix binding
page_id: operation-patch-accounts-account-id-dls-regional-services-prefix-bindings-binding-26adfea2
path: operations/prefix-bindings
description: |-
    Update the region of an existing BYOIP prefix binding.

    Like creating a binding, this requires **IP Prefixes Write** in addition to **DLS Write**.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/dls/regional_services/prefix_bindings/{binding_id}
operation_ids:
    - publicPatchPrefixBinding
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update a DLS prefix binding

`PATCH /accounts/{account_id}/dls/regional_services/prefix_bindings/{binding_id}`

Operation ID: `publicPatchPrefixBinding`

Update the region of an existing BYOIP prefix binding.

Like creating a binding, this requires **IP Prefixes Write** in addition to **DLS Write**.

## Definition

```yaml
{"operationId": "publicPatchPrefixBinding", "summary": "Update a DLS prefix binding", "description": "Update the region of an existing BYOIP prefix binding.\n\nLike creating a binding, this requires **IP Prefixes Write** in addition to **DLS Write**.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/dls_account_identifier"}}, {"name": "binding_id", "in": "path", "required": true, "schema": {"description": "Unique identifier for the prefix binding.", "type": "string", "example": "a1b2c3d4-e5f6-7890-abcd-ef1234567890"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dls_UpdatePrefixBindingInput"}}}}, "responses": {"200": {"description": "Binding updated.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dls_PrefixBindingResponse"}}}}, "400": {"description": "Bad request.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dls_bad_response"}}}}, "401": {"description": "Unauthorized.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dls_bad_response"}}}}, "403": {"description": "Forbidden.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dls_bad_response"}}}}, "404": {"description": "Not found.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dls_bad_response"}}}}, "500": {"description": "Internal server error.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dls_bad_response"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Prefix Bindings"], "x-api-token-group": ["DLS: Write", "IP Prefixes: Write"]}
```
