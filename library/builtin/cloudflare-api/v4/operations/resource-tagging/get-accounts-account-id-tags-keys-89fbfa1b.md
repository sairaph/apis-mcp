---
title: List tag keys
page_id: operation-get-accounts-account-id-tags-keys-0ad67a60
path: operations/resource-tagging
description: Lists all distinct tag keys used across resources in an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/tags/keys
operation_ids:
    - tags-list-keys
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List tag keys

`GET /accounts/{account_id}/tags/keys`

Operation ID: `tags-list-keys`

Lists all distinct tag keys used across resources in an account.

## Definition

```yaml
{"operationId": "tags-list-keys", "summary": "List tag keys", "description": "Lists all distinct tag keys used across resources in an account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/resource-tagging_account_id"}}, {"$ref": "#/components/parameters/resource-tagging_cursor"}], "responses": {"200": {"description": "List tag keys response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/resource-tagging_tag_keys_response_collection"}}}}, "4XX": {"description": "List tag keys response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/resource-tagging_api-response-common-failure"}}}}, "5XX": {"description": "List tag keys response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/resource-tagging_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Resource Tagging"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "beta", "x-fern-sdk-group-name": "keys", "x-fern-sdk-method-name": "list", "x-stability": "beta"}
```
