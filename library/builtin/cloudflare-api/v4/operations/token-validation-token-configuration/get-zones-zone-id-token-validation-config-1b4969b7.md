---
title: List token validation configurations
page_id: operation-get-zones-zone-id-token-validation-config-c9ec24c9
path: operations/token-validation-token-configuration
description: Lists all token validation configurations for this zone
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/token_validation/config
operation_ids:
    - token-validation-config-list
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List token validation configurations

`GET /zones/{zone_id}/token_validation/config`

Operation ID: `token-validation-config-list`

Lists all token validation configurations for this zone

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/api-shield_zone_id"}]
```

## Definition

```yaml
{"operationId": "token-validation-config-list", "summary": "List token validation configurations", "description": "Lists all token validation configurations for this zone", "parameters": [{"$ref": "#/components/parameters/api-shield_page"}, {"$ref": "#/components/parameters/api-shield_per_page"}], "responses": {"200": {"$ref": "#/components/responses/api-shield_list-configs-success"}, "4XX": {"$ref": "#/components/responses/api-shield_generic_failure"}}, "externalDocs": {"description": "Learn more about JSON Web Tokens Validation.", "url": "https://developers.cloudflare.com/api-shield/security/jwt-validation/"}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Token Validation Token Configuration"], "x-api-token-group": ["Account API Gateway", "Account API Gateway Read", "Domain API Gateway", "Domain API Gateway Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "token-validation.configuration", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
