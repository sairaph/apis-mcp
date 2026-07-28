---
title: Update Zero Trust list
page_id: operation-put-accounts-account-id-gateway-lists-list-id-7b21e9eb
path: operations/zero-trust-lists
description: Updates a configured Zero Trust list. Skips updating list items if not included in the payload. A non empty list items will overwrite the existing list.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/gateway/lists/{list_id}
operation_ids:
    - zero-trust-lists-update-zero-trust-list
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update Zero Trust list

`PUT /accounts/{account_id}/gateway/lists/{list_id}`

Operation ID: `zero-trust-lists-update-zero-trust-list`

Updates a configured Zero Trust list. Skips updating list items if not included in the payload. A non empty list items will overwrite the existing list.

## Definition

```yaml
{"operationId": "zero-trust-lists-update-zero-trust-list", "summary": "Update Zero Trust list", "description": "Updates a configured Zero Trust list. Skips updating list items if not included in the payload. A non empty list items will overwrite the existing list.", "parameters": [{"name": "list_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zero-trust-gateway_uuid-2"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zero-trust-gateway_identifier-2"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"description": {"$ref": "#/components/schemas/zero-trust-gateway_description"}, "items": {"$ref": "#/components/schemas/zero-trust-gateway_items-input"}, "name": {"$ref": "#/components/schemas/zero-trust-gateway_name"}}, "required": ["name"]}}}}, "responses": {"200": {"description": "Update Zero Trust list response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zero-trust-gateway_list_single_response"}}}}, "4XX": {"description": "Update Zero Trust list response failure.", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/zero-trust-gateway_list_single_response"}, {"$ref": "#/components/schemas/zero-trust-gateway_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Zero Trust lists"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.gateway.lists", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
