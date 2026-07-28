---
title: Get live aggregate device details by dimension
page_id: operation-get-accounts-account-id-dex-fleet-status-live-43b4f4bb
path: operations/dex-synthetic-application-monitoring
description: Get details for live (up to 60 minutes) devices using WARP.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/dex/fleet-status/live
operation_ids:
    - dex-fleet-status-live
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get live aggregate device details by dimension

`GET /accounts/{account_id}/dex/fleet-status/live`

Operation ID: `dex-fleet-status-live`

Get details for live (up to 60 minutes) devices using WARP.

## Definition

```yaml
{"operationId": "dex-fleet-status-live", "summary": "Get live aggregate device details by dimension", "description": "Get details for live (up to 60 minutes) devices using WARP.", "parameters": [{"name": "account_id", "in": "path", "description": "Unique Cloudflare account ID.", "required": true, "schema": {"$ref": "#/components/schemas/digital-experience-monitoring_account_identifier"}}, {"name": "since_minutes", "in": "query", "description": "Number of minutes before current time.", "required": true, "schema": {"$ref": "#/components/schemas/digital-experience-monitoring_since_minutes"}}], "responses": {"200": {"description": "Get device details (live) response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/digital-experience-monitoring_fleet_status_live_response"}}}}, "4XX": {"description": "Get device details (live) response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/digital-experience-monitoring_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}, {"api_token": []}, {"user_service_key": []}], "tags": ["DEX Synthetic Application Monitoring"], "x-api-token-group": ["Cloudflare DEX Write", "Cloudflare DEX Read", "Zero Trust Report", "Zero Trust Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.dex.devices.aggregates.live", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
