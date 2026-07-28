---
title: Get account billing usage
page_id: operation-get-accounts-account-id-billing-usage-9c9421fc
path: operations/usage-analytics
description: Retrieve billing usage analytics for an account. Returns time-series data for all billable product metrics including Stream, Media (Images), Rate Limiting, Load Balancing, Argo, Workers, Workers KV, Image Resizing, and Spectrum.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/billing/usage
operation_ids:
    - usage-analytics-get-account-billing-usage
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get account billing usage

`GET /accounts/{account_id}/billing/usage`

Operation ID: `usage-analytics-get-account-billing-usage`

Retrieve billing usage analytics for an account. Returns time-series data for all billable product metrics including Stream, Media (Images), Rate Limiting, Load Balancing, Argo, Workers, Workers KV, Image Resizing, and Spectrum.

## Definition

```yaml
{"operationId": "usage-analytics-get-account-billing-usage", "summary": "Get account billing usage", "description": "Retrieve billing usage analytics for an account. Returns time-series data for all billable product metrics including Stream, Media (Images), Rate Limiting, Load Balancing, Argo, Workers, Workers KV, Image Resizing, and Spectrum.\n", "parameters": [{"$ref": "#/components/parameters/usage-analytics_account_id"}, {"$ref": "#/components/parameters/usage-analytics_metrics"}, {"$ref": "#/components/parameters/usage-analytics_since"}, {"$ref": "#/components/parameters/usage-analytics_until"}, {"$ref": "#/components/parameters/usage-analytics_time_delta"}, {"$ref": "#/components/parameters/usage-analytics_limit"}, {"$ref": "#/components/parameters/usage-analytics_filters"}], "responses": {"200": {"description": "Usage analytics response.", "content": {"application/json": {"examples": {"success": {"summary": "Successful billing usage retrieval", "value": {"errors": [], "messages": [], "result": [{"argoAcceleratedBytes": 5000000, "imageResizingRequests": 15000, "loadBalancingQueries": 10000, "mediaUniqueTransformations": 45000, "rateLimitingRequestsAllowed": 50000, "spectrumBytesTransferred": 8000000, "streamMinutesViewed": 125000, "ts": 1693526400, "workersKVReads": 30000, "workersRequests": 200000}], "success": true}}}, "schema": {"$ref": "#/components/schemas/usage-analytics_billing_usage_response"}}}}, "400": {"$ref": "#/components/responses/usage-analytics_bad_request"}, "401": {"$ref": "#/components/responses/usage-analytics_unauthorized"}, "403": {"$ref": "#/components/responses/usage-analytics_forbidden"}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Usage Analytics"], "x-api-token-group": ["Billing Read"]}
```
