---
title: Delete a Zero Trust Gateway rule
page_id: operation-delete-accounts-account-id-gateway-rules-rule-id-e429bad9
path: operations/zero-trust-gateway-rules
description: Delete a Zero Trust Gateway rule.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/gateway/rules/{rule_id}
operation_ids:
    - zero-trust-gateway-rules-delete-zero-trust-gateway-rule
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete a Zero Trust Gateway rule

`DELETE /accounts/{account_id}/gateway/rules/{rule_id}`

Operation ID: `zero-trust-gateway-rules-delete-zero-trust-gateway-rule`

Delete a Zero Trust Gateway rule.

## Definition

```yaml
{"operationId": "zero-trust-gateway-rules-delete-zero-trust-gateway-rule", "summary": "Delete a Zero Trust Gateway rule", "description": "Delete a Zero Trust Gateway rule.", "parameters": [{"name": "rule_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zero-trust-gateway_uuid-2"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zero-trust-gateway_identifier-2"}}], "requestBody": {"required": true, "content": {"application/json": {}}}, "responses": {"200": {"description": "Delete a Zero Trust Gateway rule response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zero-trust-gateway_empty_response"}}}}, "4XX": {"description": "Delete a Zero Trust Gateway rule response failure.", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/zero-trust-gateway_empty_response"}, {"$ref": "#/components/schemas/zero-trust-gateway_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Zero Trust Gateway rules"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.gateway.rules", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
