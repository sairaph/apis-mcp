---
title: List tag values
page_id: operation-get-accounts-account-id-tags-values-tag-key-b9c6ec3f
path: operations/resource-tagging
description: Lists all distinct values for a given tag key, optionally filtered by resource type.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/tags/values/{tag_key}
operation_ids:
    - tags-list-values
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List tag values

`GET /accounts/{account_id}/tags/values/{tag_key}`

Operation ID: `tags-list-values`

Lists all distinct values for a given tag key, optionally filtered by resource type.

## Definition

```yaml
{"operationId": "tags-list-values", "summary": "List tag values", "description": "Lists all distinct values for a given tag key, optionally filtered by resource type.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/resource-tagging_account_id"}}, {"name": "tag_key", "in": "path", "description": "The tag key to retrieve values for.", "required": true, "schema": {"type": "string", "example": "environment"}}, {"$ref": "#/components/parameters/resource-tagging_type"}, {"$ref": "#/components/parameters/resource-tagging_cursor"}], "responses": {"200": {"description": "List tag values response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/resource-tagging_tag_values_response_collection"}}}}, "4XX": {"description": "List tag values response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/resource-tagging_api-response-common-failure"}}}}, "5XX": {"description": "List tag values response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/resource-tagging_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Resource Tagging"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "beta", "x-fern-sdk-group-name": "values", "x-fern-sdk-method-name": "list", "x-stability": "beta"}
```
