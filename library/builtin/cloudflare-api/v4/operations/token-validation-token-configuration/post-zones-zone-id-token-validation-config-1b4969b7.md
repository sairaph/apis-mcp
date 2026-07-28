---
title: Create a new Token Validation configuration
page_id: operation-post-zones-zone-id-token-validation-config-58518f6b
path: operations/token-validation-token-configuration
description: Create a new Token Validation configuration
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /zones/{zone_id}/token_validation/config
operation_ids:
    - token-validation-config-create
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a new Token Validation configuration

`POST /zones/{zone_id}/token_validation/config`

Operation ID: `token-validation-config-create`

Create a new Token Validation configuration

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/api-shield_zone_id"}]
```

## Definition

```yaml
{"operationId": "token-validation-config-create", "summary": "Create a new Token Validation configuration", "description": "Create a new Token Validation configuration", "requestBody": {"$ref": "#/components/requestBodies/api-shield_create-config"}, "responses": {"200": {"$ref": "#/components/responses/api-shield_create-config-success"}, "4XX": {"$ref": "#/components/responses/api-shield_generic_failure"}}, "externalDocs": {"description": "Learn more about JSON Web Tokens Validation.", "url": "https://developers.cloudflare.com/api-shield/security/jwt-validation/"}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Token Validation Token Configuration"], "x-api-token-group": ["Account API Gateway", "Domain API Gateway"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "token-validation.configuration", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
