---
title: Get indicator feeds owned by this account
page_id: operation-get-accounts-account-id-intel-indicator-feeds-e8cc68b3
path: operations/custom-indicator-feeds
description: Retrieves details for all accessible custom threat indicator feeds.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/intel/indicator-feeds
operation_ids:
    - custom-indicator-feeds-get-indicator-feeds
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get indicator feeds owned by this account

`GET /accounts/{account_id}/intel/indicator-feeds`

Operation ID: `custom-indicator-feeds-get-indicator-feeds`

Retrieves details for all accessible custom threat indicator feeds.

## Definition

```yaml
{"operationId": "custom-indicator-feeds-get-indicator-feeds", "summary": "Get indicator feeds owned by this account", "description": "Retrieves details for all accessible custom threat indicator feeds.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/custom-indicator-feeds_identifier"}}], "responses": {"200": {"description": "Get indicator feeds response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/custom-indicator-feeds_indicator_feed_response"}}}}, "4XX": {"description": "Get indicator feeds response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/custom-indicator-feeds_indicator_feed_response"}, {"$ref": "#/components/schemas/custom-indicator-feeds_api_response_common_failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Custom Indicator Feeds"], "x-api-token-group": ["Intel Write", "Intel Read"]}
```
