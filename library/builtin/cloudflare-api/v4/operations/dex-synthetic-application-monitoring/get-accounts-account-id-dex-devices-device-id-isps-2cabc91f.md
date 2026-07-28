---
title: List device ISPs
page_id: operation-get-accounts-account-id-dex-devices-device-id-isps-66eb654c
path: operations/dex-synthetic-application-monitoring
description: List ISP information observed for a specific device during traceroute tests.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/dex/devices/{device_id}/isps
operation_ids:
    - dex-endpoints-list-device-isps
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List device ISPs

`GET /accounts/{account_id}/dex/devices/{device_id}/isps`

Operation ID: `dex-endpoints-list-device-isps`

List ISP information observed for a specific device during traceroute tests.

## Definition

```yaml
{"operationId": "dex-endpoints-list-device-isps", "summary": "List device ISPs", "description": "List ISP information observed for a specific device during traceroute tests.", "parameters": [{"name": "account_id", "in": "path", "description": "Unique Cloudflare account ID.", "required": true, "schema": {"$ref": "#/components/schemas/digital-experience-monitoring_account_identifier"}}, {"name": "device_id", "in": "path", "description": "Device-specific ID, given as UUID.", "required": true, "schema": {"$ref": "#/components/schemas/digital-experience-monitoring_uuid"}}, {"name": "page", "in": "query", "description": "Page number of paginated results. Mutually exclusive with cursor.", "schema": {"type": "integer", "example": 1, "default": 1, "minimum": 1}}, {"name": "per_page", "in": "query", "description": "Number of items per page", "required": true, "schema": {"type": "integer", "example": 10, "maximum": 50, "minimum": 1}}, {"name": "cursor", "in": "query", "description": "Cursor for cursor-based pagination. Mutually exclusive with page.", "schema": {"type": "string"}}, {"name": "sort_by", "in": "query", "description": "The field to sort results by.", "schema": {"type": "string", "default": "time_start", "enum": ["time_start"]}}, {"name": "sort_order", "in": "query", "description": "The order to sort results.", "schema": {"type": "string", "default": "DESC", "enum": ["ASC", "DESC"]}}, {"name": "from", "in": "query", "description": "Start time for the query in ISO 8601 format.", "schema": {"type": "string", "format": "date-time"}}, {"name": "to", "in": "query", "description": "End time for the query in ISO 8601 format.", "schema": {"type": "string", "format": "date-time"}}], "responses": {"200": {"description": "List of ISPs observed for the device.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/digital-experience-monitoring_api-response-collection"}, {"properties": {"result": {"$ref": "#/components/schemas/digital-experience-monitoring_device_isps_response"}}}]}}}}, "4XX": {"description": "List device ISPs failure response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/digital-experience-monitoring_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}, {"api_token": []}, {"user_service_key": []}], "tags": ["DEX Synthetic Application Monitoring"], "x-api-token-group": ["Cloudflare DEX Write", "Cloudflare DEX Read", "Zero Trust Report", "Zero Trust Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.dex.isps", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
