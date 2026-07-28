---
title: List account shares
page_id: operation-get-accounts-account-id-shares-0dd57110
path: operations/resource-sharing
description: Lists all account shares.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/shares
operation_ids:
    - shares-list
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List account shares

`GET /accounts/{account_id}/shares`

Operation ID: `shares-list`

Lists all account shares.

## Definition

```yaml
{"operationId": "shares-list", "summary": "List account shares", "description": "Lists all account shares.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/resource-sharing_account_id"}}, {"$ref": "#/components/parameters/resource-sharing_status"}, {"$ref": "#/components/parameters/resource-sharing_kind"}, {"$ref": "#/components/parameters/resource-sharing_target_type"}, {"$ref": "#/components/parameters/resource-sharing_resource_types"}, {"$ref": "#/components/parameters/resource-sharing_order"}, {"$ref": "#/components/parameters/resource-sharing_direction"}, {"$ref": "#/components/parameters/resource-sharing_page"}, {"$ref": "#/components/parameters/resource-sharing_per_page"}, {"$ref": "#/components/parameters/resource-sharing_include_resources"}, {"$ref": "#/components/parameters/resource-sharing_include_recipient_counts"}, {"$ref": "#/components/parameters/resource-sharing_tag"}], "responses": {"200": {"description": "List account shares response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/resource-sharing_share_response_collection"}}}}, "4XX": {"description": "List account shares response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/resource-sharing_api-response-common-failure"}]}}}}, "5XX": {"description": "List account shares response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/resource-sharing_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Resource Sharing"], "x-api-token-group": null, "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "resource-sharing", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
