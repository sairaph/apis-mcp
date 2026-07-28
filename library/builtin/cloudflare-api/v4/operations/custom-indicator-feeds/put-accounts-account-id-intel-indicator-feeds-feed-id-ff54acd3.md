---
title: Update indicator feed metadata
page_id: operation-put-accounts-account-id-intel-indicator-feeds-feed-id-33ecfbf6
path: operations/custom-indicator-feeds
description: Revises details for a specific custom threat indicator feed.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/intel/indicator-feeds/{feed_id}
operation_ids:
    - custom-indicator-feeds-update-indicator-feed-metadata
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update indicator feed metadata

`PUT /accounts/{account_id}/intel/indicator-feeds/{feed_id}`

Operation ID: `custom-indicator-feeds-update-indicator-feed-metadata`

Revises details for a specific custom threat indicator feed.

## Definition

```yaml
{"operationId": "custom-indicator-feeds-update-indicator-feed-metadata", "summary": "Update indicator feed metadata", "description": "Revises details for a specific custom threat indicator feed.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/custom-indicator-feeds_identifier"}}, {"name": "feed_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/custom-indicator-feeds_feed_id"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/custom-indicator-feeds_update_public_field_request"}}}}, "responses": {"200": {"description": "Get update public field response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/custom-indicator-feeds_update_public_field_response"}}}}, "4XX": {"description": "Get update public field response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/custom-indicator-feeds_update_public_field_response"}, {"$ref": "#/components/schemas/custom-indicator-feeds_api_response_common_failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Custom Indicator Feeds"]}
```
