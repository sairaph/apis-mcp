---
title: Get Zero Trust list items
page_id: operation-get-accounts-account-id-gateway-lists-list-id-items-750fa6d9
path: operations/zero-trust-lists
description: Fetch all items in a single Zero Trust list.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/gateway/lists/{list_id}/items
operation_ids:
    - zero-trust-lists-zero-trust-list-items
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Zero Trust list items

`GET /accounts/{account_id}/gateway/lists/{list_id}/items`

Operation ID: `zero-trust-lists-zero-trust-list-items`

Fetch all items in a single Zero Trust list.

## Definition

```yaml
{"operationId": "zero-trust-lists-zero-trust-list-items", "summary": "Get Zero Trust list items", "description": "Fetch all items in a single Zero Trust list.", "parameters": [{"name": "list_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zero-trust-gateway_uuid-2"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zero-trust-gateway_identifier-2"}}], "responses": {"200": {"description": "Get Zero Trust list items response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zero-trust-gateway_list_item_response_collection"}}}}, "4XX": {"description": "Get Zero Trust list items response failure.", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/zero-trust-gateway_list_item_response_collection"}, {"$ref": "#/components/schemas/zero-trust-gateway_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Zero Trust lists"]}
```
