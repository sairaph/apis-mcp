---
title: Get the status over time for a device
page_id: operation-get-accounts-account-id-dex-devices-device-id-fleet-status-over-time-8bd79013
path: operations/dex-synthetic-application-monitoring
description: Get time-bucketed status metrics for a specific device.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/dex/devices/{device_id}/fleet-status/over-time
operation_ids:
    - dex-device-status-over-time
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get the status over time for a device

`GET /accounts/{account_id}/dex/devices/{device_id}/fleet-status/over-time`

Operation ID: `dex-device-status-over-time`

Get time-bucketed status metrics for a specific device.

## Definition

```yaml
{"operationId": "dex-device-status-over-time", "summary": "Get the status over time for a device", "description": "Get time-bucketed status metrics for a specific device.", "parameters": [{"name": "account_id", "in": "path", "description": "Unique Cloudflare account ID.", "required": true, "schema": {"$ref": "#/components/schemas/digital-experience-monitoring_account_identifier"}}, {"name": "device_id", "in": "path", "description": "Device-specific ID, given as UUID.", "required": true, "schema": {"$ref": "#/components/schemas/digital-experience-monitoring_device_id"}}, {"name": "from", "in": "query", "description": "Start of the time range to query. Timestamp can be provided in ISO 8601 datetime format or milliseconds since epoch.", "required": true, "schema": {"$ref": "#/components/schemas/digital-experience-monitoring_timestamp"}}, {"name": "to", "in": "query", "description": "End of the time range to query. Timestamp can be provided in ISO 8601 datetime format or milliseconds since epoch.", "required": true, "schema": {"$ref": "#/components/schemas/digital-experience-monitoring_timestamp"}}, {"name": "interval", "in": "query", "description": "Time interval for aggregate time slots.", "required": true, "schema": {"type": "string", "enum": ["minute", "hour"]}}, {"name": "colo", "in": "query", "description": "List of data centers to filter results.", "schema": {"$ref": "#/components/schemas/digital-experience-monitoring_colo"}}], "responses": {"200": {"description": "Device status over time response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/digital-experience-monitoring_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/digital-experience-monitoring_device_status_over_time_result"}}}]}}}}, "4XX": {"description": "Device status over time failure response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/digital-experience-monitoring_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}, {"api_token": []}, {"user_service_key": []}], "tags": ["DEX Synthetic Application Monitoring"], "x-api-token-group": ["Cloudflare DEX Write", "Cloudflare DEX Read", "Zero Trust Report", "Zero Trust Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.dex.devices.over-time", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
