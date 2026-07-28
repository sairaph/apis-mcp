---
title: Get a single Token Configuration
page_id: operation-get-zones-zone-id-token-validation-config-config-id-926b6247
path: operations/token-validation-token-configuration
description: Get a single Token Configuration
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/token_validation/config/{config_id}
operation_ids:
    - token-validation-config-get
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get a single Token Configuration

`GET /zones/{zone_id}/token_validation/config/{config_id}`

Operation ID: `token-validation-config-get`

Get a single Token Configuration

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/api-shield_zone_id"}, {"$ref": "#/components/parameters/api-shield_config_id"}]
```

## Definition

```yaml
{"operationId": "token-validation-config-get", "summary": "Get a single Token Configuration", "description": "Get a single Token Configuration", "responses": {"200": {"$ref": "#/components/responses/api-shield_get-config-success"}, "4XX": {"$ref": "#/components/responses/api-shield_generic_failure"}}, "externalDocs": {"description": "Learn more about JSON Web Tokens Validation.", "url": "https://developers.cloudflare.com/api-shield/security/jwt-validation/"}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Token Validation Token Configuration"], "x-api-token-group": ["Account API Gateway", "Account API Gateway Read", "Domain API Gateway", "Domain API Gateway Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "token-validation.configuration", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
