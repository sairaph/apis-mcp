---
title: Get over time aggregate details for devices by dimension
page_id: operation-get-accounts-account-id-dex-fleet-status-over-time-aa27e6d6
path: operations/dex-synthetic-application-monitoring
description: Get aggregate details for devices using WARP, up to 7 days.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/dex/fleet-status/over-time
operation_ids:
    - dex-fleet-status-over-time
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get over time aggregate details for devices by dimension

`GET /accounts/{account_id}/dex/fleet-status/over-time`

Operation ID: `dex-fleet-status-over-time`

Get aggregate details for devices using WARP, up to 7 days.

## Definition

```yaml
{"operationId": "dex-fleet-status-over-time", "summary": "Get over time aggregate details for devices by dimension", "description": "Get aggregate details for devices using WARP, up to 7 days.", "parameters": [{"name": "account_id", "in": "path", "description": "Unique Cloudflare account ID.", "required": true, "schema": {"$ref": "#/components/schemas/digital-experience-monitoring_account_identifier"}}, {"name": "to", "in": "query", "description": "End of the time range to query. Timestamp can be provided in ISO 8601 datetime format or milliseconds since epoch.", "required": true, "schema": {"$ref": "#/components/schemas/digital-experience-monitoring_timestamp"}}, {"name": "from", "in": "query", "description": "Start of the time range to query. Timestamp can be provided in ISO 8601 datetime format or milliseconds since epoch.", "required": true, "schema": {"$ref": "#/components/schemas/digital-experience-monitoring_timestamp"}}, {"name": "colo", "in": "query", "description": "Cloudflare colo airport code.", "schema": {"$ref": "#/components/schemas/digital-experience-monitoring_colo"}}, {"name": "device_id", "in": "query", "description": "Device-specific ID, given as UUID.", "schema": {"$ref": "#/components/schemas/digital-experience-monitoring_device_id"}}], "responses": {"200": {"description": "Get DEX devices aggregate success response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/digital-experience-monitoring_fleet_status_over_time_response"}}}}, "4XX": {"description": "Get DEX devices aggregate failure response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/digital-experience-monitoring_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}, {"api_token": []}, {"user_service_key": []}], "tags": ["DEX Synthetic Application Monitoring"], "x-api-token-group": ["Cloudflare DEX Write", "Cloudflare DEX Read", "Zero Trust Report", "Zero Trust Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.dex.devices.aggregates.over-time", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
