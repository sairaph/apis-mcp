---
title: Get Zero Trust list details
page_id: operation-get-accounts-account-id-gateway-lists-list-id-ec3e357b
path: operations/zero-trust-lists
description: Fetch a single Zero Trust list.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/gateway/lists/{list_id}
operation_ids:
    - zero-trust-lists-zero-trust-list-details
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Zero Trust list details

`GET /accounts/{account_id}/gateway/lists/{list_id}`

Operation ID: `zero-trust-lists-zero-trust-list-details`

Fetch a single Zero Trust list.

## Definition

```yaml
{"operationId": "zero-trust-lists-zero-trust-list-details", "summary": "Get Zero Trust list details", "description": "Fetch a single Zero Trust list.", "parameters": [{"name": "list_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zero-trust-gateway_uuid-2"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zero-trust-gateway_identifier-2"}}], "responses": {"200": {"description": "Get Zero Trust list details response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zero-trust-gateway_list_single_response"}}}}, "4XX": {"description": "Get Zero Trust list details response failure.", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/zero-trust-gateway_list_single_response"}, {"$ref": "#/components/schemas/zero-trust-gateway_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Zero Trust lists"]}
```
