---
title: Delete a registry from the account
page_id: operation-delete-accounts-account-id-containers-registries-domain-1286bf02
path: operations/image-registries
description: Delete a registry from the account, this will make Cloudchamber unable to pull images from the registry
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/containers/registries/{domain}
operation_ids:
    - deleteImageRegistry
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete a registry from the account

`DELETE /accounts/{account_id}/containers/registries/{domain}`

Operation ID: `deleteImageRegistry`

Delete a registry from the account, this will make Cloudchamber unable to pull images from the registry

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/cc_AccountId"}]
```

## Definition

```yaml
{"operationId": "deleteImageRegistry", "summary": "Delete a registry from the account", "description": "Delete a registry from the account, this will make Cloudchamber unable to pull images from the registry", "parameters": [{"name": "domain", "in": "path", "required": true, "schema": {"description": "The domain to delete", "type": "string"}}], "responses": {"200": {"description": "The image registry is deleted", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/cc_V4BaseResponse"}, {"properties": {"result": {"$ref": "#/components/schemas/cc_DeleteRegistryResponseBody"}}, "required": ["result"], "type": "object"}]}}}}, "403": {"description": "The image registry cannot be deleted", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cc_V4BaseErrorResponse"}}}}, "404": {"description": "The image registry does not exist", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cc_V4BaseErrorResponse"}}}}, "409": {"description": "The image registry cannot be deleted because it is still referenced by applications on the account", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cc_V4BaseErrorResponse"}}}}, "500": {"description": "An internal error has occurred.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cc_V4BaseErrorResponse"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Image Registries"], "x-api-token-group": ["Workers Containers Write"]}
```
