---
title: Get TLD details
page_id: operation-get-radar-tlds-tld-c50bae48
path: operations/radar-top-level-domains
description: Retrieves the requested TLD information.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /radar/tlds/{tld}
operation_ids:
    - radar-get-tld-details
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get TLD details

`GET /radar/tlds/{tld}`

Operation ID: `radar-get-tld-details`

Retrieves the requested TLD information.

## Definition

```yaml
{"operationId": "radar-get-tld-details", "summary": "Get TLD details", "description": "Retrieves the requested TLD information.", "parameters": [{"name": "tld", "in": "path", "description": "Top-level domain.", "required": true, "schema": {"description": "Top-level domain.", "type": "string", "example": "com", "maxLength": 63, "minLength": 2, "pattern": "^[a-z0-9](?:[a-z0-9-]*[a-z0-9])$"}}, {"name": "format", "in": "query", "description": "Format in which results will be returned.", "schema": {"description": "Format in which results will be returned.", "type": "string", "example": "json", "enum": ["JSON", "CSV"]}}], "responses": {"200": {"description": "Successful response.", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "object", "properties": {"tld": {"type": "object", "properties": {"manager": {"description": "The organization that manages the TLD.", "type": "string", "example": "VeriSign Global Registry Services"}, "tld": {"description": "The actual TLD.", "type": "string", "example": "com"}, "type": {"description": "The type of TLD.", "type": "string", "example": "GENERIC"}}, "required": ["tld", "type", "manager"]}}, "required": ["tld"]}, "success": {"type": "boolean", "example": true}}, "required": ["result", "success"]}}}}, "404": {"description": "Not found.", "content": {"application/json": {"schema": {"type": "object", "properties": {"error": {"type": "string", "example": "Not Found."}}, "required": ["error"]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Radar Top-Level Domains"], "x-api-token-group": ["User Details Write", "User Details Read"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "radar.tlds", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
