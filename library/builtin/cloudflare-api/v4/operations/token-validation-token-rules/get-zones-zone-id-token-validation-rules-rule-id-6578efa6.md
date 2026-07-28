---
title: Get a zone token validation rule
page_id: operation-get-zones-zone-id-token-validation-rules-rule-id-06be7752
path: operations/token-validation-token-rules
description: Get a zone token validation rule.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/token_validation/rules/{rule_id}
operation_ids:
    - token-validation-rules-get
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get a zone token validation rule

`GET /zones/{zone_id}/token_validation/rules/{rule_id}`

Operation ID: `token-validation-rules-get`

Get a zone token validation rule.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/api-shield_zone_id"}, {"$ref": "#/components/parameters/api-shield_rule_id"}]
```

## Definition

```yaml
{"operationId": "token-validation-rules-get", "summary": "Get a zone token validation rule", "description": "Get a zone token validation rule.", "responses": {"200": {"$ref": "#/components/responses/api-shield_get-rule-success"}, "4XX": {"$ref": "#/components/responses/api-shield_generic_failure"}}, "externalDocs": {"description": "Learn more about JSON Web Tokens Validation.", "url": "https://developers.cloudflare.com/api-shield/security/jwt-validation/"}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Token Validation Token Rules"], "x-api-token-group": ["Account API Gateway", "Account API Gateway Read", "Domain API Gateway", "Domain API Gateway Read"]}
```
