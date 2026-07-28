---
title: List Zero Trust lists
page_id: operation-get-accounts-account-id-gateway-lists-520c34ac
path: operations/zero-trust-lists
description: Fetch all Zero Trust lists for an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/gateway/lists
operation_ids:
    - zero-trust-lists-list-zero-trust-lists
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Zero Trust lists

`GET /accounts/{account_id}/gateway/lists`

Operation ID: `zero-trust-lists-list-zero-trust-lists`

Fetch all Zero Trust lists for an account.

## Definition

```yaml
{"operationId": "zero-trust-lists-list-zero-trust-lists", "summary": "List Zero Trust lists", "description": "Fetch all Zero Trust lists for an account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zero-trust-gateway_identifier-2"}}, {"name": "type", "in": "query", "schema": {"$ref": "#/components/schemas/zero-trust-gateway_type-2"}, "deprecated": true}], "responses": {"200": {"description": "List Zero Trust lists response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zero-trust-gateway_response_collection-3"}}}}, "4XX": {"description": "List Zero Trust lists response failure.", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/zero-trust-gateway_response_collection-3"}, {"$ref": "#/components/schemas/zero-trust-gateway_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Zero Trust lists"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.gateway.lists", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
