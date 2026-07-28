---
title: Add a new image registry configuration
page_id: operation-post-accounts-account-id-containers-registries-77886200
path: operations/image-registries
description: Add a new image registry into your account, so then Cloudflare can pull docker images with public key JWT authentication
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/containers/registries
operation_ids:
    - createImageRegistry
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Add a new image registry configuration

`POST /accounts/{account_id}/containers/registries`

Operation ID: `createImageRegistry`

Add a new image registry into your account, so then Cloudflare can pull docker images with public key JWT authentication

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/cc_AccountId"}]
```

## Definition

```yaml
{"operationId": "createImageRegistry", "summary": "Add a new image registry configuration", "description": "Add a new image registry into your account, so then Cloudflare can pull docker images with public key JWT authentication", "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cc_CreateImageRegistryRequestBody"}}}}, "responses": {"201": {"description": "Created a new image registry in the account", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/cc_V4BaseResponse"}, {"properties": {"result": {"$ref": "#/components/schemas/cc_CustomerImageRegistry"}}, "required": ["result"], "type": "object"}]}}}}, "400": {"description": "Image registry input is malformed, see the error details", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cc_V4BaseErrorResponse"}}}}, "403": {"description": "The registry that is being added is not allowed", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cc_V4BaseErrorResponse"}}}}, "409": {"description": "The image registry already exists in the account", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cc_V4BaseErrorResponse"}}}}, "500": {"description": "An internal error has occurred.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cc_V4BaseErrorResponse"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Image Registries"], "x-api-token-group": ["Workers Containers Write"]}
```
