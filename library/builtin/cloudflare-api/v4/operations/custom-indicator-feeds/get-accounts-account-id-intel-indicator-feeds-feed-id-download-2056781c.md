---
title: Download indicator feed data
page_id: operation-get-accounts-account-id-intel-indicator-feeds-feed-id-download-49826c75
path: operations/custom-indicator-feeds
description: Downloads the content of a custom threat indicator feed.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/intel/indicator-feeds/{feed_id}/download
operation_ids:
    - custom-indicator-feeds-download-indicator-feed-data
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Download indicator feed data

`GET /accounts/{account_id}/intel/indicator-feeds/{feed_id}/download`

Operation ID: `custom-indicator-feeds-download-indicator-feed-data`

Downloads the content of a custom threat indicator feed.

## Definition

```yaml
{"operationId": "custom-indicator-feeds-download-indicator-feed-data", "summary": "Download indicator feed data", "description": "Downloads the content of a custom threat indicator feed.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/custom-indicator-feeds_identifier"}}, {"name": "feed_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/custom-indicator-feeds_feed_id"}}], "responses": {"200": {"description": "Get indicator feed metadata", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/custom-indicator-feeds_update_feed_response"}}}}, "4XX": {"description": "Get indicator feeds response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/custom-indicator-feeds_api_response_common_failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Custom Indicator Feeds"], "x-api-token-group": ["Intel Write", "Intel Read"]}
```
