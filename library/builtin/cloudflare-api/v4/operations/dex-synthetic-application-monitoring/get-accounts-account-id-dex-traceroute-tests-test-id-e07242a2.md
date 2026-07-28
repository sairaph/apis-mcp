---
title: Get details and aggregate metrics for a traceroute test
page_id: operation-get-accounts-account-id-dex-traceroute-tests-test-id-9420857f
path: operations/dex-synthetic-application-monitoring
description: Get test details and aggregate performance metrics for a traceroute test for a given time period between 1 hour and 7 days.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/dex/traceroute-tests/{test_id}
operation_ids:
    - dex-endpoints-traceroute-test-details
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get details and aggregate metrics for a traceroute test

`GET /accounts/{account_id}/dex/traceroute-tests/{test_id}`

Operation ID: `dex-endpoints-traceroute-test-details`

Get test details and aggregate performance metrics for a traceroute test for a given time period between 1 hour and 7 days.

## Definition

```yaml
{"operationId": "dex-endpoints-traceroute-test-details", "summary": "Get details and aggregate metrics for a traceroute test", "description": "Get test details and aggregate performance metrics for a traceroute test for a given time period between 1 hour and 7 days.", "parameters": [{"name": "account_id", "in": "path", "description": "Unique Cloudflare account ID.", "required": true, "schema": {"$ref": "#/components/schemas/digital-experience-monitoring_account_identifier"}}, {"name": "test_id", "in": "path", "description": "Unique identifier for a specific test.", "required": true, "schema": {"$ref": "#/components/schemas/digital-experience-monitoring_uuid"}}, {"name": "deviceId", "in": "query", "description": "Optionally filter result stats to a specific device(s). Cannot be used in combination with colo param.", "schema": {"type": "array", "items": {"type": "string"}}}, {"name": "from", "in": "query", "description": "Start time for aggregate metrics in ISO ms.", "required": true, "schema": {"type": "string", "example": 1689520412000}}, {"name": "to", "in": "query", "description": "End time for aggregate metrics in ISO ms.", "required": true, "schema": {"type": "string", "example": 1689606812000}}, {"name": "interval", "in": "query", "description": "Time interval for aggregate time slots.", "required": true, "schema": {"type": "string", "enum": ["minute", "hour"]}}, {"name": "colo", "in": "query", "description": "Optionally filter result stats to a Cloudflare colo. Cannot be used in combination with deviceId param.", "schema": {"type": "string"}}], "responses": {"200": {"description": "DEX traceroute test details response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/digital-experience-monitoring_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/digital-experience-monitoring_traceroute_details_response"}}}]}}}}, "4XX": {"description": "DEX traceroute test details response failure response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/digital-experience-monitoring_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}, {"api_token": []}, {"user_service_key": []}], "tags": ["DEX Synthetic Application Monitoring"], "x-api-token-group": ["Cloudflare DEX Write", "Cloudflare DEX Read", "Zero Trust Report", "Zero Trust Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.dex.test-results.traceroute", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
