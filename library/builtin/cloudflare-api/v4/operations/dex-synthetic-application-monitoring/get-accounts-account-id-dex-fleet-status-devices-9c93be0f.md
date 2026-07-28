---
title: List details of devices using WARP.
page_id: operation-get-accounts-account-id-dex-fleet-status-devices-085c1dde
path: operations/dex-synthetic-application-monitoring
description: List details of devices using WARP.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/dex/fleet-status/devices
operation_ids:
    - dex-fleet-status-devices
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List details of devices using WARP.

`GET /accounts/{account_id}/dex/fleet-status/devices`

Operation ID: `dex-fleet-status-devices`

List details of devices using WARP.

## Definition

```yaml
{"operationId": "dex-fleet-status-devices", "summary": "List details of devices using WARP.", "description": "List details of devices using WARP.", "parameters": [{"name": "account_id", "in": "path", "description": "Unique Cloudflare account ID.", "required": true, "schema": {"$ref": "#/components/schemas/digital-experience-monitoring_account_identifier"}}, {"name": "to", "in": "query", "description": "End of the time range to query. Timestamp can be provided in ISO 8601 datetime format or milliseconds since epoch.", "required": true, "schema": {"$ref": "#/components/schemas/digital-experience-monitoring_timestamp"}}, {"name": "from", "in": "query", "description": "Start of the time range to query. Timestamp can be provided in ISO 8601 datetime format or milliseconds since epoch.", "required": true, "schema": {"$ref": "#/components/schemas/digital-experience-monitoring_timestamp"}}, {"name": "page", "in": "query", "description": "Page number of paginated results.", "required": true, "schema": {"$ref": "#/components/schemas/digital-experience-monitoring_page"}}, {"name": "per_page", "in": "query", "description": "Number of results per page.", "required": true, "schema": {"$ref": "#/components/schemas/digital-experience-monitoring_per_page"}}, {"name": "sort_by", "in": "query", "description": "Dimension to sort results by.", "schema": {"$ref": "#/components/schemas/digital-experience-monitoring_sort_by"}}, {"name": "colo", "in": "query", "description": "Cloudflare colo airport code.", "schema": {"$ref": "#/components/schemas/digital-experience-monitoring_colo"}}, {"name": "device_id", "in": "query", "description": "Device-specific ID, given as UUID.", "schema": {"$ref": "#/components/schemas/digital-experience-monitoring_device_id"}}, {"name": "mode", "in": "query", "description": "The mode under which the WARP client is run.", "schema": {"$ref": "#/components/schemas/digital-experience-monitoring_mode"}}, {"name": "status", "in": "query", "description": "Network status.", "schema": {"$ref": "#/components/schemas/digital-experience-monitoring_status"}}, {"name": "platform", "in": "query", "description": "Operating system.", "schema": {"$ref": "#/components/schemas/digital-experience-monitoring_platform"}}, {"name": "version", "in": "query", "description": "WARP client version.", "schema": {"$ref": "#/components/schemas/digital-experience-monitoring_version"}}, {"name": "source", "in": "query", "description": "Source:\n  * `hourly` - device details aggregated hourly, up to 7 days prior\n  * `last_seen` - device details, up to 60 minutes prior. Time windows exceeding 60 minutes will be rejected from June 1st, 2026. Please use 'hourly' or 'raw' instead for longer time ranges.\n  * `raw` - device details, up to 7 days prior\n", "schema": {"$ref": "#/components/schemas/digital-experience-monitoring_source"}}], "responses": {"200": {"description": "List devices response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/digital-experience-monitoring_fleet_status_devices_response"}}}}, "4XX": {"description": "List devices response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/digital-experience-monitoring_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}, {"api_token": []}, {"user_service_key": []}], "tags": ["DEX Synthetic Application Monitoring"], "x-api-token-group": ["Cloudflare DEX Write", "Cloudflare DEX Read", "Zero Trust Report", "Zero Trust Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.dex.devices", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
