---
title: Get Account-Level Metrics
page_id: operation-get-accounts-account-id-r2-metrics-332371e6
path: operations/r2-account
description: Get Storage/Object Count Metrics across all buckets in your account. Note that Account-Level Metrics may not immediately reflect the latest data.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/r2/metrics
operation_ids:
    - r2-get-account-level-metrics
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Account-Level Metrics

`GET /accounts/{account_id}/r2/metrics`

Operation ID: `r2-get-account-level-metrics`

Get Storage/Object Count Metrics across all buckets in your account. Note that Account-Level Metrics may not immediately reflect the latest data.

## Definition

```yaml
{"operationId": "r2-get-account-level-metrics", "summary": "Get Account-Level Metrics", "description": "Get Storage/Object Count Metrics across all buckets in your account. Note that Account-Level Metrics may not immediately reflect the latest data.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/r2_account_identifier"}}], "responses": {"200": {"description": "Get Account-Level Metrics response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/r2_v4_response"}, {"properties": {"result": {"$ref": "#/components/schemas/r2_account_level_metrics"}}, "type": "object"}]}}}}, "4XX": {"description": "Get Account-Level Metrics response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/r2_v4_response_failure"}}}}}, "security": [{"api_token": []}], "tags": ["R2 Account"], "x-cfPermissionsRequired": null, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "r2.buckets.metrics", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
