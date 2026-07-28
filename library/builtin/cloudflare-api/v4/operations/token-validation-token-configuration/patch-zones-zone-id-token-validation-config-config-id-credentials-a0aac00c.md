---
title: Edit Token Configuration credentials
page_id: operation-patch-zones-zone-id-token-validation-config-config-id-credentials-37c4f881
path: operations/token-validation-token-configuration
description: 'Edit Token Configuration credentials. The provided `keys` array defines the full resulting key set (stored keys omitted from payload are removed). For each provided key identity (`{alg,kid}`), payload fields overwrite the stored key before validation and omitted fields inherit from the stored key. Key identities must be unique within the request. Existing symmetric keys (`kty: "oct"`) preserve stored key material when `k` is omitted; send `k` to rotate. `k: null` is invalid.'
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /zones/{zone_id}/token_validation/config/{config_id}/credentials
operation_ids:
    - token-validation-config-credentials-edit
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Edit Token Configuration credentials

`PATCH /zones/{zone_id}/token_validation/config/{config_id}/credentials`

Operation ID: `token-validation-config-credentials-edit`

Edit Token Configuration credentials. The provided `keys` array defines the full resulting key set (stored keys omitted from payload are removed). For each provided key identity (`{alg,kid}`), payload fields overwrite the stored key before validation and omitted fields inherit from the stored key. Key identities must be unique within the request. Existing symmetric keys (`kty: "oct"`) preserve stored key material when `k` is omitted; send `k` to rotate. `k: null` is invalid.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/api-shield_zone_id"}, {"$ref": "#/components/parameters/api-shield_config_id"}]
```

## Definition

```yaml
{"operationId": "token-validation-config-credentials-edit", "summary": "Edit Token Configuration credentials", "description": "Edit Token Configuration credentials. The provided `keys` array defines the full resulting key set (stored keys omitted from payload are removed). For each provided key identity (`{alg,kid}`), payload fields overwrite the stored key before validation and omitted fields inherit from the stored key. Key identities must be unique within the request. Existing symmetric keys (`kty: \"oct\"`) preserve stored key material when `k` is omitted; send `k` to rotate. `k: null` is invalid.", "requestBody": {"$ref": "#/components/requestBodies/api-shield_edit-config-credentials"}, "responses": {"200": {"$ref": "#/components/responses/api-shield_update-config-credentials-success"}, "4XX": {"$ref": "#/components/responses/api-shield_generic_failure"}}, "externalDocs": {"description": "Learn more about JSON Web Tokens Validation.", "url": "https://developers.cloudflare.com/api-shield/security/jwt-validation/"}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Token Validation Token Configuration"], "x-api-token-group": ["Account API Gateway", "Domain API Gateway"]}
```
