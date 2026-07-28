---
title: Get Zero Trust Gateway rule details.
page_id: operation-get-accounts-account-id-gateway-rules-rule-id-7893963a
path: operations/zero-trust-gateway-rules
description: Get a single Zero Trust Gateway rule.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/gateway/rules/{rule_id}
operation_ids:
    - zero-trust-gateway-rules-zero-trust-gateway-rule-details
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Zero Trust Gateway rule details.

`GET /accounts/{account_id}/gateway/rules/{rule_id}`

Operation ID: `zero-trust-gateway-rules-zero-trust-gateway-rule-details`

Get a single Zero Trust Gateway rule.

## Definition

```yaml
{"operationId": "zero-trust-gateway-rules-zero-trust-gateway-rule-details", "summary": "Get Zero Trust Gateway rule details.", "description": "Get a single Zero Trust Gateway rule.", "parameters": [{"name": "rule_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zero-trust-gateway_uuid-2"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zero-trust-gateway_identifier-2"}}], "responses": {"200": {"description": "Get Zero Trust Gateway rule details response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zero-trust-gateway_single_response-4"}}}}, "4XX": {"description": "Get Zero Trust Gateway rule details response failure.", "content": {"application/json": {"schema": {"allOf": [{"type": "object"}, {"$ref": "#/components/schemas/zero-trust-gateway_single_response-4"}, {"$ref": "#/components/schemas/zero-trust-gateway_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Zero Trust Gateway rules"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.gateway.rules", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
