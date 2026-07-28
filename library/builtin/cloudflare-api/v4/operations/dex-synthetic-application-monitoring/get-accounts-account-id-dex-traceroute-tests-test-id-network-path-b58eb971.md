---
title: Get network path breakdown for a traceroute test
page_id: operation-get-accounts-account-id-dex-traceroute-tests-test-id-network-path-f20dd7bc
path: operations/dex-synthetic-application-monitoring
description: Get a breakdown of metrics by hop for individual traceroute test runs.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/dex/traceroute-tests/{test_id}/network-path
operation_ids:
    - dex-endpoints-traceroute-test-network-path
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get network path breakdown for a traceroute test

`GET /accounts/{account_id}/dex/traceroute-tests/{test_id}/network-path`

Operation ID: `dex-endpoints-traceroute-test-network-path`

Get a breakdown of metrics by hop for individual traceroute test runs.

## Definition

```yaml
{"operationId": "dex-endpoints-traceroute-test-network-path", "summary": "Get network path breakdown for a traceroute test", "description": "Get a breakdown of metrics by hop for individual traceroute test runs.", "parameters": [{"name": "account_id", "in": "path", "description": "Unique Cloudflare account ID.", "required": true, "schema": {"$ref": "#/components/schemas/digital-experience-monitoring_account_identifier"}}, {"name": "test_id", "in": "path", "description": "Unique identifier for a specific test.", "required": true, "schema": {"$ref": "#/components/schemas/digital-experience-monitoring_uuid"}}, {"name": "deviceId", "in": "query", "description": "Device to filter traceroute result runs to.", "required": true, "schema": {"type": "string"}}, {"name": "from", "in": "query", "description": "Start time for aggregate metrics in ISO ms.", "required": true, "schema": {"type": "string", "example": 1689520412000}}, {"name": "to", "in": "query", "description": "End time for aggregate metrics in ISO ms.", "required": true, "schema": {"type": "string", "example": 1689606812000}}, {"name": "interval", "in": "query", "description": "Time interval for aggregate time slots.", "required": true, "schema": {"type": "string", "enum": ["minute", "hour"]}}], "responses": {"200": {"description": "DEX traceroute test network path response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/digital-experience-monitoring_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/digital-experience-monitoring_traceroute_test_network_path_response"}}}]}}}}, "4XX": {"description": "DEX traceroute test network path failure response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/digital-experience-monitoring_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}, {"api_token": []}, {"user_service_key": []}], "tags": ["DEX Synthetic Application Monitoring"], "x-api-token-group": ["Cloudflare DEX Write", "Cloudflare DEX Read", "Zero Trust Report", "Zero Trust Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.dex.tests.network-path", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
