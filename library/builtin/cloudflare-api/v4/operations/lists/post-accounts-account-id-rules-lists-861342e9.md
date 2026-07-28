---
title: Create a list
page_id: operation-post-accounts-account-id-rules-lists-b59924d4
path: operations/lists
description: Creates a new list of the specified kind.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/rules/lists
operation_ids:
    - lists-create-a-list
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a list

`POST /accounts/{account_id}/rules/lists`

Operation ID: `lists-create-a-list`

Creates a new list of the specified kind.

## Definition

```yaml
{"operationId": "lists-create-a-list", "summary": "Create a list", "description": "Creates a new list of the specified kind.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/lists_account_id"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"description": {"$ref": "#/components/schemas/lists_description"}, "kind": {"$ref": "#/components/schemas/lists_kind"}, "name": {"$ref": "#/components/schemas/lists_name"}}, "required": ["name", "kind"]}}}}, "responses": {"200": {"description": "Create a list response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/lists_list-response-collection"}}}}, "4XX": {"description": "Create a list response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/lists_list-response-collection"}, {"$ref": "#/components/schemas/lists_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Lists"], "x-api-token-group": ["Account Filter Lists Edit"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
