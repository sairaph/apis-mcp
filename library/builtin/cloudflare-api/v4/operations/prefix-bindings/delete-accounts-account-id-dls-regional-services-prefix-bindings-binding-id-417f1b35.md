---
title: Delete a DLS prefix binding
page_id: operation-delete-accounts-account-id-dls-regional-services-prefix-bindings-binding-143a1a5f
path: operations/prefix-bindings
description: |-
    Delete a BYOIP prefix binding.

    Like creating a binding, this requires **IP Prefixes Write** in addition to **DLS Write**.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/dls/regional_services/prefix_bindings/{binding_id}
operation_ids:
    - publicDeletePrefixBinding
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete a DLS prefix binding

`DELETE /accounts/{account_id}/dls/regional_services/prefix_bindings/{binding_id}`

Operation ID: `publicDeletePrefixBinding`

Delete a BYOIP prefix binding.

Like creating a binding, this requires **IP Prefixes Write** in addition to **DLS Write**.

## Definition

```yaml
{"operationId": "publicDeletePrefixBinding", "summary": "Delete a DLS prefix binding", "description": "Delete a BYOIP prefix binding.\n\nLike creating a binding, this requires **IP Prefixes Write** in addition to **DLS Write**.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/dls_account_identifier"}}, {"name": "binding_id", "in": "path", "required": true, "schema": {"description": "Unique identifier for the prefix binding.", "type": "string", "example": "a1b2c3d4-e5f6-7890-abcd-ef1234567890"}}], "responses": {"200": {"description": "Binding deleted successfully.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dls_good_response"}}}}, "401": {"description": "Unauthorized.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dls_bad_response"}}}}, "403": {"description": "Forbidden.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dls_bad_response"}}}}, "404": {"description": "Not found.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dls_bad_response"}}}}, "500": {"description": "Internal server error.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dls_bad_response"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Prefix Bindings"], "x-api-token-group": ["DLS: Write", "IP Prefixes: Write"]}
```
