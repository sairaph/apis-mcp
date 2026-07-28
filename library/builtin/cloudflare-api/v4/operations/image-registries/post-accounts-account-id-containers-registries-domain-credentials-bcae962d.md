---
title: Generate a JWT to interact with the specified image registry.
page_id: operation-post-accounts-account-id-containers-registries-domain-credentials-23e0ab9f
path: operations/image-registries
description: Generates temporary credentials for accessing Cloudflare's container image registry. Used for pulling and pushing container images.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/containers/registries/{domain}/credentials
operation_ids:
    - generateImageRegistryCredentials
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Generate a JWT to interact with the specified image registry.

`POST /accounts/{account_id}/containers/registries/{domain}/credentials`

Operation ID: `generateImageRegistryCredentials`

Generates temporary credentials for accessing Cloudflare's container image registry. Used for pulling and pushing container images.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/cc_AccountId"}]
```

## Definition

```yaml
{"operationId": "generateImageRegistryCredentials", "summary": "Generate a JWT to interact with the specified image registry.", "description": "Generates temporary credentials for accessing Cloudflare's container image registry. Used for pulling and pushing container images.", "parameters": [{"name": "domain", "in": "path", "required": true, "schema": {"description": "The domain to get credentials for.", "type": "string", "example": "registry.cloudflare.com"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cc_ImageRegistryCredentialsConfiguration"}}}}, "responses": {"201": {"description": "Credentials with 'pull' or 'push' permissions to access the registry", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/cc_V4BaseResponse"}, {"properties": {"result": {"$ref": "#/components/schemas/cc_AccountRegistryToken"}}, "required": ["result"], "type": "object"}]}}}}, "400": {"description": "Bad Request for Public API.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cc_V4BaseErrorResponse"}}}}, "403": {"description": "The requested token permissions are not allowed for this account", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cc_V4BaseErrorResponse"}}}}, "404": {"description": "The image registry does not exist", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cc_V4BaseErrorResponse"}}}}, "409": {"description": "The registry was configured as public, so credentials can not be generated", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cc_V4BaseErrorResponse"}}}}, "500": {"description": "InternalError500.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cc_V4BaseErrorResponse"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Image Registries"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "workers.registries.credentials", "x-fern-sdk-method-name": "generate", "x-forge-hidden": true}
```
