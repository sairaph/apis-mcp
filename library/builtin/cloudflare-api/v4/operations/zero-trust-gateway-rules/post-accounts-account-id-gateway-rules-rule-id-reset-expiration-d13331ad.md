---
title: Reset the expiration of a Zero Trust Gateway Rule
page_id: operation-post-accounts-account-id-gateway-rules-rule-id-reset-expiration-4a125988
path: operations/zero-trust-gateway-rules
description: Resets the expiration of a Zero Trust Gateway Rule if its duration elapsed and it has a default duration. The Zero Trust Gateway Rule must have values  for both `expiration.expires_at` and `expiration.duration`.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/gateway/rules/{rule_id}/reset_expiration
operation_ids:
    - zero-trust-gateway-rules-reset-expiration-zero-trust-gateway-rule
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Reset the expiration of a Zero Trust Gateway Rule

`POST /accounts/{account_id}/gateway/rules/{rule_id}/reset_expiration`

Operation ID: `zero-trust-gateway-rules-reset-expiration-zero-trust-gateway-rule`

Resets the expiration of a Zero Trust Gateway Rule if its duration elapsed and it has a default duration. The Zero Trust Gateway Rule must have values  for both `expiration.expires_at` and `expiration.duration`.

## Definition

```yaml
{"operationId": "zero-trust-gateway-rules-reset-expiration-zero-trust-gateway-rule", "summary": "Reset the expiration of a Zero Trust Gateway Rule", "description": "Resets the expiration of a Zero Trust Gateway Rule if its duration elapsed and it has a default duration. The Zero Trust Gateway Rule must have values  for both `expiration.expires_at` and `expiration.duration`.", "parameters": [{"name": "rule_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zero-trust-gateway_uuid-2"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zero-trust-gateway_identifier-2"}}], "responses": {"200": {"description": "Reset the expiration of a Zero Trust Gateway rule response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zero-trust-gateway_single_response-4"}}}}, "4XX": {"description": "Reset the expiration of a Zero Trust Gateway rule response failure.", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/zero-trust-gateway_single_response-4"}, {"$ref": "#/components/schemas/zero-trust-gateway_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Zero Trust Gateway rules"]}
```
