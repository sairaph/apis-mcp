---
title: Get the list of configured registries in the account
page_id: operation-get-accounts-account-id-containers-registries-7639f88b
path: operations/image-registries
description: Get the list of configured registries in the account
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/containers/registries
operation_ids:
    - listImageRegistries
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get the list of configured registries in the account

`GET /accounts/{account_id}/containers/registries`

Operation ID: `listImageRegistries`

Get the list of configured registries in the account

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/cc_AccountId"}]
```

## Definition

```yaml
{"operationId": "listImageRegistries", "summary": "Get the list of configured registries in the account", "description": "Get the list of configured registries in the account", "responses": {"200": {"description": "The list of registries that are added in the account", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/cc_V4BaseResponse"}, {"properties": {"result": {"type": "array", "items": {"$ref": "#/components/schemas/cc_CustomerImageRegistry"}}}, "required": ["result"], "type": "object"}]}}}}, "500": {"description": "An internal error has occurred.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cc_V4BaseErrorResponse"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Image Registries"], "x-api-token-group": ["Workers Containers Write", "Workers Containers Read"]}
```
