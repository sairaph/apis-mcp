---
title: List WARP change events.
page_id: operation-get-accounts-account-id-dex-warp-change-events-d14bf110
path: operations/warp-change-events
description: List WARP configuration and enablement toggle change events by device.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/dex/warp-change-events
operation_ids:
    - list-warp-change-events
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List WARP change events.

`GET /accounts/{account_id}/dex/warp-change-events`

Operation ID: `list-warp-change-events`

List WARP configuration and enablement toggle change events by device.

## Definition

```yaml
{"operationId": "list-warp-change-events", "summary": "List WARP change events.", "description": "List WARP configuration and enablement toggle change events by device.", "parameters": [{"name": "account_id", "in": "path", "description": "Unique Cloudflare account ID.", "required": true, "schema": {"$ref": "#/components/schemas/digital-experience-monitoring_account_identifier"}}, {"name": "page", "in": "query", "description": "Page number of paginated results.", "required": true, "schema": {"type": "number", "example": 1, "default": 1, "minimum": 1}}, {"name": "per_page", "in": "query", "description": "Number of results per page.", "required": true, "schema": {"type": "number", "example": 10, "maximum": 50, "minimum": 1}}, {"name": "from", "in": "query", "description": "Start time for the query in ISO (RFC3339 - ISO 8601) format.", "required": true, "schema": {"type": "string", "example": "2023-09-20T17:00:00Z"}}, {"name": "to", "in": "query", "description": "End time for the query in ISO (RFC3339 - ISO 8601) format.", "required": true, "schema": {"type": "string", "example": "2023-09-20T17:00:00Z"}}, {"name": "type", "in": "query", "description": "Filter events by type 'config' or 'toggle'.", "schema": {"type": "string", "enum": ["config", "toggle"]}}, {"name": "toggle", "in": "query", "description": "Filter events by type toggle value. Applicable to type='toggle' events only.", "schema": {"type": "string", "enum": ["on", "off"]}}, {"name": "config_name", "in": "query", "description": "Filter events by WARP configuration name changed from or to. Applicable to type='config' events only.", "schema": {"type": "string", "example": "MASQUE"}}, {"name": "account_name", "in": "query", "description": "Filter events by account name.", "schema": {"type": "string", "example": "Myorg"}}, {"name": "sort_order", "in": "query", "description": "Sort response by event timestamp.", "schema": {"type": "string", "default": "ASC", "enum": ["ASC", "DESC"]}}], "responses": {"200": {"description": "Success response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/digital-experience-monitoring_api-response-collection"}, {"properties": {"result": {"$ref": "#/components/schemas/digital-experience-monitoring_warp_events_response"}}}]}}}}, "4XX": {"description": "List WARP change events failure response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/digital-experience-monitoring_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}, {"api_token": []}, {"user_service_key": []}], "tags": ["WARP Change Events"], "x-api-token-group": ["Cloudflare DEX Write", "Cloudflare DEX Read", "Zero Trust Report", "Zero Trust Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.dex.warp-change-events", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
