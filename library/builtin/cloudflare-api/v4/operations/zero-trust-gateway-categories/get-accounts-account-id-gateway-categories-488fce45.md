---
title: List categories
page_id: operation-get-accounts-account-id-gateway-categories-52b03ed8
path: operations/zero-trust-gateway-categories
description: List all categories.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/gateway/categories
operation_ids:
    - zero-trust-gateway-categories-list-categories
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List categories

`GET /accounts/{account_id}/gateway/categories`

Operation ID: `zero-trust-gateway-categories-list-categories`

List all categories.

## Definition

```yaml
{"operationId": "zero-trust-gateway-categories-list-categories", "summary": "List categories", "description": "List all categories.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zero-trust-gateway_identifier-3"}}], "responses": {"200": {"description": "List categories response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zero-trust-gateway_response_collection-8"}}}}, "4XX": {"description": "List categories response failure.", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/zero-trust-gateway_response_collection-8"}, {"$ref": "#/components/schemas/zero-trust-gateway_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Zero Trust Gateway categories"], "x-api-token-group": ["Zero Trust Read", "Zero Trust Write"]}
```
