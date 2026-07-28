---
title: Create new indicator feed
page_id: operation-post-accounts-account-id-intel-indicator-feeds-0fde216e
path: operations/custom-indicator-feeds
description: Creates a new custom threat indicator feed for sharing threat intelligence data.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/intel/indicator-feeds
operation_ids:
    - custom-indicator-feeds-create-indicator-feeds
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create new indicator feed

`POST /accounts/{account_id}/intel/indicator-feeds`

Operation ID: `custom-indicator-feeds-create-indicator-feeds`

Creates a new custom threat indicator feed for sharing threat intelligence data.

## Definition

```yaml
{"operationId": "custom-indicator-feeds-create-indicator-feeds", "summary": "Create new indicator feed", "description": "Creates a new custom threat indicator feed for sharing threat intelligence data.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/custom-indicator-feeds_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/custom-indicator-feeds_create_feed"}}}}, "responses": {"200": {"description": "Create indicator feed response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/custom-indicator-feeds_create_feed_response"}}}}, "4XX": {"description": "Get indicator feeds failure response", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/custom-indicator-feeds_create_feed_response"}, {"$ref": "#/components/schemas/custom-indicator-feeds_api_response_common_failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Custom Indicator Feeds"]}
```
