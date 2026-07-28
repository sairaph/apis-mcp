---
title: Update a Zero Trust Gateway rule
page_id: operation-put-accounts-account-id-gateway-rules-rule-id-a24103ff
path: operations/zero-trust-gateway-rules
description: Update a configured Zero Trust Gateway rule.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/gateway/rules/{rule_id}
operation_ids:
    - zero-trust-gateway-rules-update-zero-trust-gateway-rule
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update a Zero Trust Gateway rule

`PUT /accounts/{account_id}/gateway/rules/{rule_id}`

Operation ID: `zero-trust-gateway-rules-update-zero-trust-gateway-rule`

Update a configured Zero Trust Gateway rule.

## Definition

```yaml
{"operationId": "zero-trust-gateway-rules-update-zero-trust-gateway-rule", "summary": "Update a Zero Trust Gateway rule", "description": "Update a configured Zero Trust Gateway rule.", "parameters": [{"name": "rule_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zero-trust-gateway_uuid-2"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zero-trust-gateway_identifier-2"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"action": {"$ref": "#/components/schemas/zero-trust-gateway_action"}, "description": {"$ref": "#/components/schemas/zero-trust-gateway_description-2"}, "device_posture": {"$ref": "#/components/schemas/zero-trust-gateway_device_posture"}, "enabled": {"$ref": "#/components/schemas/zero-trust-gateway_enabled"}, "expiration": {"$ref": "#/components/schemas/zero-trust-gateway_expiration"}, "filters": {"$ref": "#/components/schemas/zero-trust-gateway_filters"}, "identity": {"$ref": "#/components/schemas/zero-trust-gateway_identity"}, "name": {"$ref": "#/components/schemas/zero-trust-gateway_name-3"}, "precedence": {"$ref": "#/components/schemas/zero-trust-gateway_precedence"}, "rule_settings": {"$ref": "#/components/schemas/zero-trust-gateway_rule-settings"}, "schedule": {"$ref": "#/components/schemas/zero-trust-gateway_schedule"}, "traffic": {"$ref": "#/components/schemas/zero-trust-gateway_traffic"}}, "required": ["name", "action"]}}}}, "responses": {"200": {"description": "Update a Zero Trust Gateway rule response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zero-trust-gateway_single_response-4"}}}}, "4XX": {"description": "Update a Zero Trust Gateway rule response failure.", "content": {"application/json": {"schema": {"allOf": [{"type": "object"}, {"$ref": "#/components/schemas/zero-trust-gateway_single_response-4"}, {"$ref": "#/components/schemas/zero-trust-gateway_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Zero Trust Gateway rules"]}
```
