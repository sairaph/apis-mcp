---
title: Get indicator feed data
page_id: operation-get-accounts-account-id-intel-indicator-feeds-feed-id-data-aa24e7c3
path: operations/custom-indicator-feeds
description: Retrieves the raw data entries in a custom threat indicator feed.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/intel/indicator-feeds/{feed_id}/data
operation_ids:
    - custom-indicator-feeds-get-indicator-feed-data
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get indicator feed data

`GET /accounts/{account_id}/intel/indicator-feeds/{feed_id}/data`

Operation ID: `custom-indicator-feeds-get-indicator-feed-data`

Retrieves the raw data entries in a custom threat indicator feed.

## Definition

```yaml
{"operationId": "custom-indicator-feeds-get-indicator-feed-data", "summary": "Get indicator feed data", "description": "Retrieves the raw data entries in a custom threat indicator feed.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/custom-indicator-feeds_identifier"}}, {"name": "feed_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/custom-indicator-feeds_feed_id"}}], "responses": {"200": {"description": "Get indicator feed metadata", "content": {"text/csv": {"schema": {"type": "string"}}}}, "4XX": {"description": "Get indicator feeds response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/custom-indicator-feeds_api_response_common_failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Custom Indicator Feeds"], "x-api-token-group": ["Intel Write", "Intel Read"]}
```
