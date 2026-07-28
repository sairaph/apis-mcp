---
title: List tagged resources
page_id: operation-get-accounts-account-id-tags-resources-e3deb558
path: operations/resource-tagging
description: Lists all tagged resources for an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/tags/resources
operation_ids:
    - tags-list
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List tagged resources

`GET /accounts/{account_id}/tags/resources`

Operation ID: `tags-list`

Lists all tagged resources for an account.

## Definition

```yaml
{"operationId": "tags-list", "summary": "List tagged resources", "description": "Lists all tagged resources for an account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/resource-tagging_account_id"}}, {"name": "type", "in": "query", "description": "Filter by resource type. Can be repeated to filter by multiple types (OR logic). Example: ?type=zone&type=worker", "schema": {"type": "array", "items": {"$ref": "#/components/schemas/resource-tagging_resource_type"}}, "explode": true, "style": "form"}, {"name": "name", "in": "query", "description": "Filter by resource name. Performs a case-insensitive substring match. Example: ?name=my-zone", "schema": {"type": "string", "maxLength": 256}}, {"name": "id", "in": "query", "description": "Filter by resource ID. Can be repeated up to 50 times to filter by multiple IDs. Example: ?id=abc&id=def", "schema": {"type": "array", "items": {"$ref": "#/components/schemas/resource-tagging_resource_id"}, "maxItems": 50}, "explode": true, "style": "form"}, {"$ref": "#/components/parameters/resource-tagging_tag"}, {"$ref": "#/components/parameters/resource-tagging_cursor"}], "responses": {"200": {"description": "List tagged resources response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/resource-tagging_tagged_resource_response_collection"}}}}, "4XX": {"description": "List tagged resources response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/resource-tagging_api-response-common-failure"}}}}, "5XX": {"description": "List tagged resources response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/resource-tagging_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Resource Tagging"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "beta", "x-fern-sdk-group-name": "tags.resources", "x-fern-sdk-method-name": "list", "x-stability": "beta"}
```
