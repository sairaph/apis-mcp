---
title: Get account Stream usage
page_id: operation-get-accounts-account-id-stream-usage-a78e57e6
path: operations/usage-analytics
description: Retrieve Stream usage analytics for an account. Returns time-series data for Stream billable minutes viewed across all zones in the account. The gateway rewrites this path before forwarding to the backend usage handler.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/stream/usage
operation_ids:
    - usage-analytics-get-account-stream-usage
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get account Stream usage

`GET /accounts/{account_id}/stream/usage`

Operation ID: `usage-analytics-get-account-stream-usage`

Retrieve Stream usage analytics for an account. Returns time-series data for Stream billable minutes viewed across all zones in the account. The gateway rewrites this path before forwarding to the backend usage handler.

## Definition

```yaml
{"operationId": "usage-analytics-get-account-stream-usage", "summary": "Get account Stream usage", "description": "Retrieve Stream usage analytics for an account. Returns time-series data for Stream billable minutes viewed across all zones in the account. The gateway rewrites this path before forwarding to the backend usage handler.\n", "parameters": [{"$ref": "#/components/parameters/usage-analytics_account_id"}, {"$ref": "#/components/parameters/usage-analytics_metrics"}, {"$ref": "#/components/parameters/usage-analytics_since"}, {"$ref": "#/components/parameters/usage-analytics_until"}, {"$ref": "#/components/parameters/usage-analytics_time_delta"}, {"$ref": "#/components/parameters/usage-analytics_limit"}, {"$ref": "#/components/parameters/usage-analytics_filters"}], "responses": {"200": {"description": "Usage analytics response.", "content": {"application/json": {"examples": {"success": {"summary": "Successful Stream usage retrieval", "value": {"errors": [], "messages": [], "result": [{"streamMinutesViewed": 125000, "ts": 1693526400}, {"streamMinutesViewed": 98000, "ts": 1693530000}], "success": true}}}, "schema": {"$ref": "#/components/schemas/usage-analytics_stream_usage_response"}}}}, "400": {"$ref": "#/components/responses/usage-analytics_bad_request"}, "401": {"$ref": "#/components/responses/usage-analytics_unauthorized"}, "403": {"$ref": "#/components/responses/usage-analytics_forbidden"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Usage Analytics"], "x-api-token-group": ["Stream Read"]}
```
