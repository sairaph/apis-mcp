---
title: Create a Zero Trust Gateway rule
page_id: operation-post-accounts-account-id-gateway-rules-ca242374
path: operations/zero-trust-gateway-rules
description: Create a new Zero Trust Gateway rule.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/gateway/rules
operation_ids:
    - zero-trust-gateway-rules-create-zero-trust-gateway-rule
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a Zero Trust Gateway rule

`POST /accounts/{account_id}/gateway/rules`

Operation ID: `zero-trust-gateway-rules-create-zero-trust-gateway-rule`

Create a new Zero Trust Gateway rule.

## Definition

```yaml
{"operationId": "zero-trust-gateway-rules-create-zero-trust-gateway-rule", "summary": "Create a Zero Trust Gateway rule", "description": "Create a new Zero Trust Gateway rule.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zero-trust-gateway_identifier-2"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"action": {"$ref": "#/components/schemas/zero-trust-gateway_action"}, "description": {"$ref": "#/components/schemas/zero-trust-gateway_description-2"}, "device_posture": {"$ref": "#/components/schemas/zero-trust-gateway_device_posture"}, "enabled": {"$ref": "#/components/schemas/zero-trust-gateway_enabled"}, "expiration": {"$ref": "#/components/schemas/zero-trust-gateway_expiration"}, "filters": {"$ref": "#/components/schemas/zero-trust-gateway_filters"}, "identity": {"$ref": "#/components/schemas/zero-trust-gateway_identity"}, "name": {"$ref": "#/components/schemas/zero-trust-gateway_name-3"}, "precedence": {"$ref": "#/components/schemas/zero-trust-gateway_precedence"}, "rule_settings": {"$ref": "#/components/schemas/zero-trust-gateway_rule-settings"}, "schedule": {"$ref": "#/components/schemas/zero-trust-gateway_schedule"}, "traffic": {"$ref": "#/components/schemas/zero-trust-gateway_traffic"}}, "required": ["name", "action"]}}}}, "responses": {"200": {"description": "Create a Zero Trust Gateway rule response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zero-trust-gateway_single_response-4"}}}}, "4XX": {"description": "Create a Zero Trust Gateway rule response failure.", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"type": "object"}, {"$ref": "#/components/schemas/zero-trust-gateway_single_response-4"}, {"$ref": "#/components/schemas/zero-trust-gateway_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Zero Trust Gateway rules"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.gateway.rules", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
