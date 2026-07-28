---
title: Delete a zone token validation rule
page_id: operation-delete-zones-zone-id-token-validation-rules-rule-id-6e6ab292
path: operations/token-validation-token-rules
description: Delete a zone token validation rule.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /zones/{zone_id}/token_validation/rules/{rule_id}
operation_ids:
    - token-validation-rules-delete
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete a zone token validation rule

`DELETE /zones/{zone_id}/token_validation/rules/{rule_id}`

Operation ID: `token-validation-rules-delete`

Delete a zone token validation rule.

## Path Parameters

```yaml
[{"$ref": "#/components/parameters/api-shield_zone_id"}, {"$ref": "#/components/parameters/api-shield_rule_id"}]
```

## Definition

```yaml
{"operationId": "token-validation-rules-delete", "summary": "Delete a zone token validation rule", "description": "Delete a zone token validation rule.", "responses": {"200": {"$ref": "#/components/responses/api-shield_delete-rule-success"}, "4XX": {"$ref": "#/components/responses/api-shield_generic_failure"}}, "externalDocs": {"description": "Learn more about JSON Web Tokens Validation.", "url": "https://developers.cloudflare.com/api-shield/security/jwt-validation/"}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Token Validation Token Rules"], "x-api-token-group": ["Account API Gateway", "Domain API Gateway"]}
```
