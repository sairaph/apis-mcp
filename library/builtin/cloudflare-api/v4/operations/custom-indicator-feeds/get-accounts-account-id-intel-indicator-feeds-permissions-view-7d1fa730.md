---
title: List indicator feed permissions
page_id: operation-get-accounts-account-id-intel-indicator-feeds-permissions-view-ba94b649
path: operations/custom-indicator-feeds
description: Lists current access permissions for custom threat indicator feeds.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/intel/indicator-feeds/permissions/view
operation_ids:
    - custom-indicator-feeds-view-permissions
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List indicator feed permissions

`GET /accounts/{account_id}/intel/indicator-feeds/permissions/view`

Operation ID: `custom-indicator-feeds-view-permissions`

Lists current access permissions for custom threat indicator feeds.

## Definition

```yaml
{"operationId": "custom-indicator-feeds-view-permissions", "summary": "List indicator feed permissions", "description": "Lists current access permissions for custom threat indicator feeds.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/custom-indicator-feeds_identifier"}}], "responses": {"200": {"description": "Get indicator feed metadata", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/custom-indicator-feeds_permission_list_item_response"}}}}, "4XX": {"description": "Get indicator feeds response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/custom-indicator-feeds_permission_list_item_response"}, {"$ref": "#/components/schemas/custom-indicator-feeds_api_response_common_failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Custom Indicator Feeds"], "x-api-token-group": ["Intel Write", "Intel Read"]}
```
