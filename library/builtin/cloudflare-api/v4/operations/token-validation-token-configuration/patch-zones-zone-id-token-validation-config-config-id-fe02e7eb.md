---
title: Edit an existing Token Configuration
page_id: operation-patch-zones-zone-id-token-validation-config-config-id-1fc3bccc
path: operations/token-validation-token-configuration
description: Edit fields of an existing Token Configuration
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /zones/{zone_id}/token_validation/config/{config_id}
operation_ids:
    - token-validation-config-edit
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Edit an existing Token Configuration

`PATCH /zones/{zone_id}/token_validation/config/{config_id}`

Operation ID: `token-validation-config-edit`

Edit fields of an existing Token Configuration

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/api-shield_zone_id"}, {"$ref": "#/components/parameters/api-shield_config_id"}]
```

## Definition

```yaml
{"operationId": "token-validation-config-edit", "summary": "Edit an existing Token Configuration", "description": "Edit fields of an existing Token Configuration", "requestBody": {"$ref": "#/components/requestBodies/api-shield_edit-config"}, "responses": {"200": {"$ref": "#/components/responses/api-shield_edit-config-success"}, "4XX": {"$ref": "#/components/responses/api-shield_generic_failure"}}, "externalDocs": {"description": "Learn more about JSON Web Tokens Validation.", "url": "https://developers.cloudflare.com/api-shield/security/jwt-validation/"}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Token Validation Token Configuration"], "x-api-token-group": ["Account API Gateway", "Domain API Gateway"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "token-validation.configuration", "x-fern-sdk-method-name": "edit", "x-forge-hidden": true}
```
