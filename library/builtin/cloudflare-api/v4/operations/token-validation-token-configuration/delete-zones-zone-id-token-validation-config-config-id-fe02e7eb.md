---
title: Delete Token Configuration
page_id: operation-delete-zones-zone-id-token-validation-config-config-id-b47e9f46
path: operations/token-validation-token-configuration
description: Delete Token Configuration
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /zones/{zone_id}/token_validation/config/{config_id}
operation_ids:
    - token-validation-config-delete
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete Token Configuration

`DELETE /zones/{zone_id}/token_validation/config/{config_id}`

Operation ID: `token-validation-config-delete`

Delete Token Configuration

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/api-shield_zone_id"}, {"$ref": "#/components/parameters/api-shield_config_id"}]
```

## Definition

```yaml
{"operationId": "token-validation-config-delete", "summary": "Delete Token Configuration", "description": "Delete Token Configuration", "responses": {"200": {"$ref": "#/components/responses/api-shield_delete-config-success"}, "4XX": {"$ref": "#/components/responses/api-shield_generic_failure"}}, "externalDocs": {"description": "Learn more about JSON Web Tokens Validation.", "url": "https://developers.cloudflare.com/api-shield/security/jwt-validation/"}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Token Validation Token Configuration"], "x-api-token-group": ["Account API Gateway", "Domain API Gateway"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "token-validation.configuration", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
