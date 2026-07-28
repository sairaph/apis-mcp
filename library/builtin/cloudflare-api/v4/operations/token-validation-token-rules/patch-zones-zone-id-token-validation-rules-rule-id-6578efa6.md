---
title: Edit a zone token validation rule
page_id: operation-patch-zones-zone-id-token-validation-rules-rule-id-0da0a12b
path: operations/token-validation-token-rules
description: Edit a zone token validation rule.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /zones/{zone_id}/token_validation/rules/{rule_id}
operation_ids:
    - token-validation-rules-edit
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Edit a zone token validation rule

`PATCH /zones/{zone_id}/token_validation/rules/{rule_id}`

Operation ID: `token-validation-rules-edit`

Edit a zone token validation rule.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/api-shield_zone_id"}, {"$ref": "#/components/parameters/api-shield_rule_id"}]
```

## Definition

```yaml
{"operationId": "token-validation-rules-edit", "summary": "Edit a zone token validation rule", "description": "Edit a zone token validation rule.", "requestBody": {"$ref": "#/components/requestBodies/api-shield_edit-rule"}, "responses": {"200": {"$ref": "#/components/responses/api-shield_edit-rule-success"}, "4XX": {"$ref": "#/components/responses/api-shield_generic_failure"}}, "externalDocs": {"description": "Learn more about JSON Web Tokens Validation.", "url": "https://developers.cloudflare.com/api-shield/security/jwt-validation/"}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Token Validation Token Rules"], "x-api-token-group": ["Account API Gateway", "Domain API Gateway"]}
```
