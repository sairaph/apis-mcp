---
title: Patch a Zero Trust Gateway rule
page_id: operation-patch-accounts-account-id-gateway-rules-rule-id-d616844f
path: operations/zero-trust-gateway-rules
description: Update select fields of an existing Zero Trust Gateway rule. Only the fields provided in the request body are updated. This endpoint supports a limited subset of fields (`name`, `description`, `precedence`, `enabled`). To update other fields such as `action`, `traffic`, `identity`, `device_posture`, `rule_settings`, `schedule`, or `expiration`, use the PUT endpoint for a full rule replacement.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/gateway/rules/{rule_id}
operation_ids:
    - zero-trust-gateway-rules-patch-zero-trust-gateway-rule
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Patch a Zero Trust Gateway rule

`PATCH /accounts/{account_id}/gateway/rules/{rule_id}`

Operation ID: `zero-trust-gateway-rules-patch-zero-trust-gateway-rule`

Update select fields of an existing Zero Trust Gateway rule. Only the fields provided in the request body are updated. This endpoint supports a limited subset of fields (`name`, `description`, `precedence`, `enabled`). To update other fields such as `action`, `traffic`, `identity`, `device_posture`, `rule_settings`, `schedule`, or `expiration`, use the PUT endpoint for a full rule replacement.

## Definition

```yaml
{"operationId": "zero-trust-gateway-rules-patch-zero-trust-gateway-rule", "summary": "Patch a Zero Trust Gateway rule", "description": "Update select fields of an existing Zero Trust Gateway rule. Only the fields provided in the request body are updated. This endpoint supports a limited subset of fields (`name`, `description`, `precedence`, `enabled`). To update other fields such as `action`, `traffic`, `identity`, `device_posture`, `rule_settings`, `schedule`, or `expiration`, use the PUT endpoint for a full rule replacement.", "parameters": [{"name": "rule_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zero-trust-gateway_uuid-2"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zero-trust-gateway_identifier-2"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"description": {"$ref": "#/components/schemas/zero-trust-gateway_description-2"}, "enabled": {"$ref": "#/components/schemas/zero-trust-gateway_enabled"}, "name": {"$ref": "#/components/schemas/zero-trust-gateway_name-3"}, "precedence": {"$ref": "#/components/schemas/zero-trust-gateway_precedence"}}}}}}, "responses": {"200": {"description": "Patch a Zero Trust Gateway rule response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zero-trust-gateway_single_response-4"}}}}, "4XX": {"description": "Patch a Zero Trust Gateway rule response failure.", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/zero-trust-gateway_single_response-4"}, {"$ref": "#/components/schemas/zero-trust-gateway_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Zero Trust Gateway rules"]}
```
