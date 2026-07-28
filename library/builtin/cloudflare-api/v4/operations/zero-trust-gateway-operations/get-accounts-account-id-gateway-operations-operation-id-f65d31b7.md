---
title: Zero Trust Gateway operation details
page_id: operation-get-accounts-account-id-gateway-operations-operation-id-da8acd5e
path: operations/zero-trust-gateway-operations
description: Fetch a single Zero Trust Gateway operation by its ID.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/gateway/operations/{operation_id}
operation_ids:
    - zero-trust-gateway-operations-zero-trust-gateway-operation-details
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Zero Trust Gateway operation details

`GET /accounts/{account_id}/gateway/operations/{operation_id}`

Operation ID: `zero-trust-gateway-operations-zero-trust-gateway-operation-details`

Fetch a single Zero Trust Gateway operation by its ID.

## Definition

```yaml
{"operationId": "zero-trust-gateway-operations-zero-trust-gateway-operation-details", "summary": "Zero Trust Gateway operation details", "description": "Fetch a single Zero Trust Gateway operation by its ID.", "parameters": [{"name": "operation_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zero-trust-gateway_uuid-2"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zero-trust-gateway_identifier-3"}}], "responses": {"200": {"description": "Zero Trust Gateway operation details response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zero-trust-gateway_single_response-2"}}}}, "4XX": {"description": "Zero Trust Gateway operation details response failure.", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/zero-trust-gateway_single_response-2"}, {"$ref": "#/components/schemas/zero-trust-gateway_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Zero Trust Gateway operations"], "x-api-token-group": ["Zero Trust Read", "Zero Trust Write"]}
```
