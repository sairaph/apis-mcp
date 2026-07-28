---
title: Create Zero Trust list
page_id: operation-post-accounts-account-id-gateway-lists-6fbdea7d
path: operations/zero-trust-lists
description: Creates a new Zero Trust list.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/gateway/lists
operation_ids:
    - zero-trust-lists-create-zero-trust-list
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create Zero Trust list

`POST /accounts/{account_id}/gateway/lists`

Operation ID: `zero-trust-lists-create-zero-trust-list`

Creates a new Zero Trust list.

## Definition

```yaml
{"operationId": "zero-trust-lists-create-zero-trust-list", "summary": "Create Zero Trust list", "description": "Creates a new Zero Trust list.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zero-trust-gateway_identifier-2"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"description": {"$ref": "#/components/schemas/zero-trust-gateway_description"}, "items": {"$ref": "#/components/schemas/zero-trust-gateway_items-input"}, "name": {"$ref": "#/components/schemas/zero-trust-gateway_name"}, "type": {"$ref": "#/components/schemas/zero-trust-gateway_type-2"}}, "required": ["name", "type"]}}}}, "responses": {"200": {"description": "Create Zero Trust list response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zero-trust-gateway_single_response_with_list_items"}}}}, "4XX": {"description": "Create Zero Trust list response failure.", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/zero-trust-gateway_single_response_with_list_items"}, {"$ref": "#/components/schemas/zero-trust-gateway_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Zero Trust lists"]}
```
