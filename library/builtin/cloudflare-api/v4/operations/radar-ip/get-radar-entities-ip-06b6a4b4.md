---
title: Get IP address details
page_id: operation-get-radar-entities-ip-0b03f8f7
path: operations/radar-ip
description: Retrieves IP address information.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /radar/entities/ip
operation_ids:
    - radar-get-entities-ip
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get IP address details

`GET /radar/entities/ip`

Operation ID: `radar-get-entities-ip`

Retrieves IP address information.

## Definition

```yaml
{"operationId": "radar-get-entities-ip", "summary": "Get IP address details", "description": "Retrieves IP address information.", "parameters": [{"name": "ip", "in": "query", "description": "IP address.", "required": true, "schema": {"description": "IP address.", "type": "string", "format": "ip", "example": "8.8.8.8"}}, {"name": "format", "in": "query", "description": "Format in which results will be returned.", "schema": {"description": "Format in which results will be returned.", "type": "string", "example": "json", "enum": ["JSON", "CSV"]}}], "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "object", "properties": {"ip": {"type": "object", "properties": {"asn": {"type": "string", "example": "15169"}, "asnLocation": {"type": "string", "example": "US"}, "asnName": {"type": "string", "example": "GOOGLE"}, "asnOrgName": {"type": "string", "example": "Google LLC"}, "ip": {"type": "string", "example": "8.8.8.8"}, "ipVersion": {"type": "string", "example": "IPv4"}, "location": {"type": "string", "example": "GB"}, "locationName": {"type": "string", "example": "United Kingdom"}}, "required": ["ip", "ipVersion", "location", "locationName", "asn", "asnName", "asnLocation", "asnOrgName"]}}, "required": ["ip"]}, "success": {"type": "boolean", "example": true}}, "required": ["result", "success"]}}}}, "404": {"description": "Not found.", "content": {"application/json": {"schema": {"type": "object", "properties": {"error": {"type": "string", "example": "Not Found."}}, "required": ["error"]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Radar IP"], "x-api-token-group": ["User Details Write", "User Details Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "radar.entities", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
