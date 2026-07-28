---
title: Create a DLS prefix binding
page_id: operation-post-accounts-account-id-dls-regional-services-prefix-bindings-3c3452db
path: operations/prefix-bindings
description: |-
    Bind a CIDR from a BYOIP prefix to a region.

    This requires the **IP Prefixes Write** permission in addition to **DLS Write**, because the binding is created against a BYOIP prefix in Addressing.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/dls/regional_services/prefix_bindings
operation_ids:
    - publicCreatePrefixBinding
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a DLS prefix binding

`POST /accounts/{account_id}/dls/regional_services/prefix_bindings`

Operation ID: `publicCreatePrefixBinding`

Bind a CIDR from a BYOIP prefix to a region.

This requires the **IP Prefixes Write** permission in addition to **DLS Write**, because the binding is created against a BYOIP prefix in Addressing.

## Definition

```yaml
{"operationId": "publicCreatePrefixBinding", "summary": "Create a DLS prefix binding", "description": "Bind a CIDR from a BYOIP prefix to a region.\n\nThis requires the **IP Prefixes Write** permission in addition to **DLS Write**, because the binding is created against a BYOIP prefix in Addressing.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/dls_account_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dls_CreatePrefixBindingInput"}}}}, "responses": {"201": {"description": "Binding created.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dls_PrefixBindingResponse"}}}}, "400": {"description": "Bad request.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dls_bad_response"}}}}, "401": {"description": "Unauthorized.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dls_bad_response"}}}}, "403": {"description": "Forbidden.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dls_bad_response"}}}}, "409": {"description": "Conflict.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dls_bad_response"}}}}, "500": {"description": "Internal server error.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dls_bad_response"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Prefix Bindings"], "x-api-token-group": ["DLS: Write", "IP Prefixes: Write"]}
```
