---
title: Get details for a specific traceroute test run
page_id: operation-get-accounts-account-id-dex-traceroute-test-results-test-result-id-netwo-c3895c1f
path: operations/dex-synthetic-application-monitoring
description: Get a breakdown of hops and performance metrics for a specific traceroute test run
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/dex/traceroute-test-results/{test_result_id}/network-path
operation_ids:
    - dex-endpoints-traceroute-test-result-network-path
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get details for a specific traceroute test run

`GET /accounts/{account_id}/dex/traceroute-test-results/{test_result_id}/network-path`

Operation ID: `dex-endpoints-traceroute-test-result-network-path`

Get a breakdown of hops and performance metrics for a specific traceroute test run

## Definition

```yaml
{"operationId": "dex-endpoints-traceroute-test-result-network-path", "summary": "Get details for a specific traceroute test run", "description": "Get a breakdown of hops and performance metrics for a specific traceroute test run", "parameters": [{"name": "account_id", "in": "path", "description": "Unique Cloudflare account ID.", "required": true, "schema": {"$ref": "#/components/schemas/digital-experience-monitoring_account_identifier"}}, {"name": "test_result_id", "in": "path", "description": "Unique identifier for a specific traceroute test.", "required": true, "schema": {"$ref": "#/components/schemas/digital-experience-monitoring_uuid"}}], "responses": {"200": {"description": "DEX traceroute test result network path response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/digital-experience-monitoring_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/digital-experience-monitoring_traceroute_test_result_network_path_response"}}}]}}}}, "4XX": {"description": "DEX traceroute test result network path failure response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/digital-experience-monitoring_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}, {"api_token": []}, {"user_service_key": []}], "tags": ["DEX Synthetic Application Monitoring"], "x-api-token-group": ["Cloudflare DEX Write", "Cloudflare DEX Read", "Zero Trust Report", "Zero Trust Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.dex.test-results.traceroute.network-path", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
