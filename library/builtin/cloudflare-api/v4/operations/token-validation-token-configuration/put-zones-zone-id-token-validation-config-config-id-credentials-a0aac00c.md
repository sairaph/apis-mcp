---
title: Update Token Configuration credentials
page_id: operation-put-zones-zone-id-token-validation-config-config-id-credentials-fd7447fc
path: operations/token-validation-token-configuration
description: 'Update Token Configuration credentials with full replacement semantics. Key identities (`{alg,kid}`) must be unique within the request. Symmetric keys (`kty: "oct"`) require `k`; `k: null` is invalid.'
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /zones/{zone_id}/token_validation/config/{config_id}/credentials
operation_ids:
    - token-validation-config-credentials-update
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update Token Configuration credentials

`PUT /zones/{zone_id}/token_validation/config/{config_id}/credentials`

Operation ID: `token-validation-config-credentials-update`

Update Token Configuration credentials with full replacement semantics. Key identities (`{alg,kid}`) must be unique within the request. Symmetric keys (`kty: "oct"`) require `k`; `k: null` is invalid.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/api-shield_zone_id"}, {"$ref": "#/components/parameters/api-shield_config_id"}]
```

## Definition

```yaml
{"operationId": "token-validation-config-credentials-update", "summary": "Update Token Configuration credentials", "description": "Update Token Configuration credentials with full replacement semantics. Key identities (`{alg,kid}`) must be unique within the request. Symmetric keys (`kty: \"oct\"`) require `k`; `k: null` is invalid.", "requestBody": {"$ref": "#/components/requestBodies/api-shield_update-config-credentials"}, "responses": {"200": {"$ref": "#/components/responses/api-shield_update-config-credentials-success"}, "4XX": {"$ref": "#/components/responses/api-shield_generic_failure"}}, "externalDocs": {"description": "Learn more about JSON Web Tokens Validation.", "url": "https://developers.cloudflare.com/api-shield/security/jwt-validation/"}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Token Validation Token Configuration"], "x-api-token-group": ["Account API Gateway", "Domain API Gateway"]}
```
