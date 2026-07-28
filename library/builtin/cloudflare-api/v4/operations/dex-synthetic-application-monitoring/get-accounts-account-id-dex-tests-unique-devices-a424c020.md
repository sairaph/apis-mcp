---
title: Get count of devices targeted
page_id: operation-get-accounts-account-id-dex-tests-unique-devices-ed2281d2
path: operations/dex-synthetic-application-monitoring
description: Returns unique count of devices that have run synthetic application monitoring tests in the past 7 days.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/dex/tests/unique-devices
operation_ids:
    - dex-endpoints-tests-unique-devices
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get count of devices targeted

`GET /accounts/{account_id}/dex/tests/unique-devices`

Operation ID: `dex-endpoints-tests-unique-devices`

Returns unique count of devices that have run synthetic application monitoring tests in the past 7 days.

## Definition

```yaml
{"operationId": "dex-endpoints-tests-unique-devices", "summary": "Get count of devices targeted", "description": "Returns unique count of devices that have run synthetic application monitoring tests in the past 7 days.", "parameters": [{"name": "account_id", "in": "path", "description": "Unique Cloudflare account ID.", "required": true, "schema": {"$ref": "#/components/schemas/digital-experience-monitoring_account_identifier"}}, {"name": "testName", "in": "query", "description": "Optionally filter results by test name.", "schema": {"type": "string"}}, {"name": "deviceId", "in": "query", "description": "Optionally filter result stats to a specific device(s). Cannot be used in combination with colo param.", "schema": {"type": "array", "items": {"description": "Unique identifier for the physical device (UUID).", "example": "cb49c27f-7f97-49c5-b6f3-f7c01ead0fd7", "type": "string"}}}], "responses": {"200": {"description": "DEX unique devices targeted response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/digital-experience-monitoring_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/digital-experience-monitoring_unique_devices_response"}}}]}}}}, "4XX": {"description": "DEX unique devices targeted failure response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/digital-experience-monitoring_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}, {"api_token": []}, {"user_service_key": []}], "tags": ["DEX Synthetic Application Monitoring"], "x-api-token-group": ["Cloudflare DEX Write", "Cloudflare DEX Read", "Zero Trust Report", "Zero Trust Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.dex.overview.tests.unique-devices", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
