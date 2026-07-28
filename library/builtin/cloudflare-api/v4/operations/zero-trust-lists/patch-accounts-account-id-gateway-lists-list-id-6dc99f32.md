---
title: Patch Zero Trust list.
page_id: operation-patch-accounts-account-id-gateway-lists-list-id-4c2a262e
path: operations/zero-trust-lists
description: Appends or removes an item from a configured Zero Trust list.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/gateway/lists/{list_id}
operation_ids:
    - zero-trust-lists-patch-zero-trust-list
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Patch Zero Trust list.

`PATCH /accounts/{account_id}/gateway/lists/{list_id}`

Operation ID: `zero-trust-lists-patch-zero-trust-list`

Appends or removes an item from a configured Zero Trust list.

## Definition

```yaml
{"operationId": "zero-trust-lists-patch-zero-trust-list", "summary": "Patch Zero Trust list.", "description": "Appends or removes an item from a configured Zero Trust list.", "parameters": [{"name": "list_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zero-trust-gateway_uuid-2"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zero-trust-gateway_identifier-2"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"append": {"$ref": "#/components/schemas/zero-trust-gateway_items-input"}, "remove": {"description": "Lists of item values you want to remove.", "type": "array", "items": {"$ref": "#/components/schemas/zero-trust-gateway_value"}}}}}}}, "responses": {"200": {"description": "Patch Zero Trust list response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zero-trust-gateway_list_single_response"}}}}, "4XX": {"description": "Patch Zero Trust list response failure.", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/zero-trust-gateway_list_single_response"}, {"$ref": "#/components/schemas/zero-trust-gateway_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Zero Trust lists"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.gateway.lists", "x-fern-sdk-method-name": "edit", "x-forge-hidden": true}
```
