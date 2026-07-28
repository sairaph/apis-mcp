---
title: Get percentiles for an http test
page_id: operation-get-accounts-account-id-dex-http-tests-test-id-percentiles-78f20087
path: operations/dex-synthetic-application-monitoring
description: Get percentiles for an http test for a given time period between 1 hour and 7 days.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/dex/http-tests/{test_id}/percentiles
operation_ids:
    - dex-endpoints-http-test-percentiles
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get percentiles for an http test

`GET /accounts/{account_id}/dex/http-tests/{test_id}/percentiles`

Operation ID: `dex-endpoints-http-test-percentiles`

Get percentiles for an http test for a given time period between 1 hour and 7 days.

## Definition

```yaml
{"operationId": "dex-endpoints-http-test-percentiles", "summary": "Get percentiles for an http test", "description": "Get percentiles for an http test for a given time period between 1 hour and 7 days.", "parameters": [{"name": "account_id", "in": "path", "description": "Unique Cloudflare account ID.", "required": true, "schema": {"$ref": "#/components/schemas/digital-experience-monitoring_account_identifier"}}, {"name": "test_id", "in": "path", "description": "Unique identifier for a specific test.", "required": true, "schema": {"$ref": "#/components/schemas/digital-experience-monitoring_uuid"}}, {"name": "deviceId", "in": "query", "description": "Optionally filter result stats to a specific device(s). Cannot be used in combination with colo param.", "schema": {"type": "array", "items": {"description": "Unique identifier for the physical device (UUID).", "example": "cb49c27f-7f97-49c5-b6f3-f7c01ead0fd7", "type": "string"}}}, {"name": "from", "in": "query", "description": "Start time for the query in ISO (RFC3339 - ISO 8601) format.", "required": true, "schema": {"type": "string", "example": "2023-09-20T17:00:00Z"}}, {"name": "to", "in": "query", "description": "End time for the query in ISO (RFC3339 - ISO 8601) format.", "required": true, "schema": {"type": "string", "example": "2023-09-20T17:00:00Z"}}, {"name": "colo", "in": "query", "description": "Optionally filter result stats to a Cloudflare colo. Cannot be used in combination with deviceId param.", "schema": {"description": "Cloudflare colo airport code.", "type": "string", "example": "SJC"}}], "responses": {"200": {"description": "DEX HTTP test percentiles response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/digital-experience-monitoring_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/digital-experience-monitoring_http_details_percentiles_response"}}}]}}}}, "4XX": {"description": "DEX HTTP test percentiles failure response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/digital-experience-monitoring_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}, {"api_token": []}, {"user_service_key": []}], "tags": ["DEX Synthetic Application Monitoring"], "x-api-token-group": ["Cloudflare DEX Write", "Cloudflare DEX Read", "Zero Trust Report", "Zero Trust Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.dex.test-results.http.percentiles", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
