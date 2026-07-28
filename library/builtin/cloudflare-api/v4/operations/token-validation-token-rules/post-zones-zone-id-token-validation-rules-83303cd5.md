---
title: Create a token validation rule
page_id: operation-post-zones-zone-id-token-validation-rules-9ff2b3f2
path: operations/token-validation-token-rules
description: Create a token validation rule.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /zones/{zone_id}/token_validation/rules
operation_ids:
    - token-validation-rules-create
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a token validation rule

`POST /zones/{zone_id}/token_validation/rules`

Operation ID: `token-validation-rules-create`

Create a token validation rule.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/api-shield_zone_id"}]
```

## Definition

```yaml
{"operationId": "token-validation-rules-create", "summary": "Create a token validation rule", "description": "Create a token validation rule.", "requestBody": {"$ref": "#/components/requestBodies/api-shield_create-rule"}, "responses": {"200": {"$ref": "#/components/responses/api-shield_create-rule-success"}, "4XX": {"$ref": "#/components/responses/api-shield_generic_failure"}}, "externalDocs": {"description": "Learn more about JSON Web Tokens Validation.", "url": "https://developers.cloudflare.com/api-shield/security/jwt-validation/"}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Token Validation Token Rules"], "x-api-token-group": ["Account API Gateway", "Domain API Gateway"]}
```
