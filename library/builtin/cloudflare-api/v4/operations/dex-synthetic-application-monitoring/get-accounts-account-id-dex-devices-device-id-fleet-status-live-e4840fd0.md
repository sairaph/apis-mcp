---
title: Get the latest status of a device.
page_id: operation-get-accounts-account-id-dex-devices-device-id-fleet-status-live-5417698b
path: operations/dex-synthetic-application-monitoring
description: Get the latest status of a device given device_id from the device_state table.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/dex/devices/{device_id}/fleet-status/live
operation_ids:
    - devices-live-status
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get the latest status of a device.

`GET /accounts/{account_id}/dex/devices/{device_id}/fleet-status/live`

Operation ID: `devices-live-status`

Get the latest status of a device given device_id from the device_state table.

## Definition

```yaml
{"operationId": "devices-live-status", "summary": "Get the latest status of a device.", "description": "Get the latest status of a device given device_id from the device_state table.", "parameters": [{"name": "account_id", "in": "path", "description": "Unique Cloudflare account ID.", "required": true, "schema": {"$ref": "#/components/schemas/digital-experience-monitoring_account_identifier"}}, {"name": "device_id", "in": "path", "description": "Device-specific ID, given as UUID.", "required": true, "schema": {"$ref": "#/components/schemas/digital-experience-monitoring_device_id"}}, {"name": "since_minutes", "in": "query", "description": "Number of minutes before current time.", "required": true, "schema": {"$ref": "#/components/schemas/digital-experience-monitoring_since_minutes"}}, {"name": "time_now", "in": "query", "description": "Current time in ISO format.", "schema": {"$ref": "#/components/schemas/digital-experience-monitoring_time_now"}, "deprecated": true, "x-stainless-deprecation-message": "This parameter is deprecated and will be removed in the future."}, {"name": "colo", "in": "query", "description": "List of data centers to filter results.", "schema": {"$ref": "#/components/schemas/digital-experience-monitoring_colo"}}], "responses": {"200": {"description": "Get the latest status of a device.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/digital-experience-monitoring_device"}}}}, "4XX": {"description": "Get the latest status of a device failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/digital-experience-monitoring_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}, {"api_token": []}, {"user_service_key": []}], "tags": ["DEX Synthetic Application Monitoring"], "x-api-token-group": ["Cloudflare DEX Write", "Cloudflare DEX Read", "Zero Trust Report", "Zero Trust Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.dex.devices", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
