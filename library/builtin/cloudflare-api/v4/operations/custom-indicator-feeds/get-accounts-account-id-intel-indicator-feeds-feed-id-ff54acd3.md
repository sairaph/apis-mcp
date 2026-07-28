---
title: Get indicator feed metadata
page_id: operation-get-accounts-account-id-intel-indicator-feeds-feed-id-d8e7860a
path: operations/custom-indicator-feeds
description: Retrieves details for a specific custom threat indicator feed.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/intel/indicator-feeds/{feed_id}
operation_ids:
    - custom-indicator-feeds-get-indicator-feed-metadata
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get indicator feed metadata

`GET /accounts/{account_id}/intel/indicator-feeds/{feed_id}`

Operation ID: `custom-indicator-feeds-get-indicator-feed-metadata`

Retrieves details for a specific custom threat indicator feed.

## Definition

```yaml
{"operationId": "custom-indicator-feeds-get-indicator-feed-metadata", "summary": "Get indicator feed metadata", "description": "Retrieves details for a specific custom threat indicator feed.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/custom-indicator-feeds_identifier"}}, {"name": "feed_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/custom-indicator-feeds_feed_id"}}], "responses": {"200": {"description": "Get indicator feed metadata", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/custom-indicator-feeds_indicator_feed_metadata_response"}}}}, "4XX": {"description": "Get indicator feeds response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/custom-indicator-feeds_indicator_feed_metadata_response"}, {"$ref": "#/components/schemas/custom-indicator-feeds_api_response_common_failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Custom Indicator Feeds"], "x-api-token-group": ["Intel Write", "Intel Read"]}
```
