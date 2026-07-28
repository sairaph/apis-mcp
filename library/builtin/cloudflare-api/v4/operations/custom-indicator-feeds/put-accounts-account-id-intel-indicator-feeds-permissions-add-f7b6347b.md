---
title: Grant permission to indicator feed
page_id: operation-put-accounts-account-id-intel-indicator-feeds-permissions-add-db7de2dd
path: operations/custom-indicator-feeds
description: Grants access permissions for a custom threat indicator feed to other accounts.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/intel/indicator-feeds/permissions/add
operation_ids:
    - custom-indicator-feeds-add-permission
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Grant permission to indicator feed

`PUT /accounts/{account_id}/intel/indicator-feeds/permissions/add`

Operation ID: `custom-indicator-feeds-add-permission`

Grants access permissions for a custom threat indicator feed to other accounts.

## Definition

```yaml
{"operationId": "custom-indicator-feeds-add-permission", "summary": "Grant permission to indicator feed", "description": "Grants access permissions for a custom threat indicator feed to other accounts.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/custom-indicator-feeds_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/custom-indicator-feeds_permissions_request"}}}}, "responses": {"200": {"description": "Get indicator feed metadata", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/custom-indicator-feeds_permissions_response"}}}}, "4XX": {"description": "Get indicator feeds response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/custom-indicator-feeds_permissions_response"}, {"$ref": "#/components/schemas/custom-indicator-feeds_api_response_common_failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Custom Indicator Feeds"]}
```
