---
title: List tag key summary
page_id: operation-get-accounts-account-id-tags-summary-f38c82d2
path: operations/resource-tagging
description: Lists all distinct tag keys and their distinct values across resources in an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/tags/summary
operation_ids:
    - tags-list-key-summary
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List tag key summary

`GET /accounts/{account_id}/tags/summary`

Operation ID: `tags-list-key-summary`

Lists all distinct tag keys and their distinct values across resources in an account.

## Definition

```yaml
{"operationId": "tags-list-key-summary", "summary": "List tag key summary", "description": "Lists all distinct tag keys and their distinct values across resources in an account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/resource-tagging_account_id"}}, {"$ref": "#/components/parameters/resource-tagging_cursor"}], "responses": {"200": {"description": "List tag key summary response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/resource-tagging_tag_key_summary_response_collection"}}}}, "4XX": {"description": "List tag key summary response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/resource-tagging_api-response-common-failure"}}}}, "5XX": {"description": "List tag key summary response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/resource-tagging_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Resource Tagging"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-stability": "beta"}
```
