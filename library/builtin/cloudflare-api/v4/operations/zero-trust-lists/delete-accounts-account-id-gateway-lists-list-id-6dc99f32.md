---
title: Delete Zero Trust list
page_id: operation-delete-accounts-account-id-gateway-lists-list-id-6c487788
path: operations/zero-trust-lists
description: Deletes a Zero Trust list.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/gateway/lists/{list_id}
operation_ids:
    - zero-trust-lists-delete-zero-trust-list
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete Zero Trust list

`DELETE /accounts/{account_id}/gateway/lists/{list_id}`

Operation ID: `zero-trust-lists-delete-zero-trust-list`

Deletes a Zero Trust list.

## Definition

```yaml
{"operationId": "zero-trust-lists-delete-zero-trust-list", "summary": "Delete Zero Trust list", "description": "Deletes a Zero Trust list.", "parameters": [{"name": "list_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zero-trust-gateway_uuid-2"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zero-trust-gateway_identifier-2"}}], "requestBody": {"required": true, "content": {"application/json": {}}}, "responses": {"200": {"description": "Delete Zero Trust list response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zero-trust-gateway_empty_response"}}}}, "4XX": {"description": "Delete Zero Trust list response failure.", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/zero-trust-gateway_empty_response"}, {"$ref": "#/components/schemas/zero-trust-gateway_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Zero Trust lists"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.gateway.lists", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
