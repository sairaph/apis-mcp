---
title: List Zero Trust Gateway operations
page_id: operation-get-accounts-account-id-gateway-operations-603adbe8
path: operations/zero-trust-gateway-operations
description: List all Zero Trust Gateway operations for an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/gateway/operations
operation_ids:
    - zero-trust-gateway-operations-list-zero-trust-gateway-operations
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Zero Trust Gateway operations

`GET /accounts/{account_id}/gateway/operations`

Operation ID: `zero-trust-gateway-operations-list-zero-trust-gateway-operations`

List all Zero Trust Gateway operations for an account.

## Definition

```yaml
{"operationId": "zero-trust-gateway-operations-list-zero-trust-gateway-operations", "summary": "List Zero Trust Gateway operations", "description": "List all Zero Trust Gateway operations for an account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zero-trust-gateway_identifier-3"}}], "responses": {"200": {"description": "List Zero Trust Gateway operations response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zero-trust-gateway_response_collection-4"}}}}, "4XX": {"description": "List Zero Trust Gateway operations response failure.", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/zero-trust-gateway_response_collection-4"}, {"$ref": "#/components/schemas/zero-trust-gateway_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Zero Trust Gateway operations"], "x-api-token-group": ["Zero Trust Read", "Zero Trust Write"]}
```
