---
title: List available Page Rules settings
page_id: operation-get-zones-zone-id-pagerules-settings-3c2cd34c
path: operations/available-page-rules-settings
description: Returns a list of settings (and their details) that Page Rules can apply to matching requests.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/pagerules/settings
operation_ids:
    - available-page-rules-settings-list-available-page-rules-settings
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List available Page Rules settings

`GET /zones/{zone_id}/pagerules/settings`

Operation ID: `available-page-rules-settings-list-available-page-rules-settings`

Returns a list of settings (and their details) that Page Rules can apply to matching requests.

## Definition

```yaml
{"operationId": "available-page-rules-settings-list-available-page-rules-settings", "summary": "List available Page Rules settings", "description": "Returns a list of settings (and their details) that Page Rules can apply to matching requests.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zones_identifier-2"}, "example": "023e105f4ecef8ad9ca31a8372d0c353"}], "responses": {"200": {"description": "List available Page Rules settings response", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/zones_api-response-common-2"}, {"properties": {"result": {"$ref": "#/components/schemas/zones_settings"}}}]}}}}, "4XX": {"description": "List available Page Rules settings response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zones_api-response-common-failure-2"}}}}}, "deprecated": true, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Available Page Rules settings"], "x-api-token-group": ["Zone Read", "Zone Write", "Page Rules Write", "Page Rules Read"], "x-fern-availability": "deprecated", "x-fern-sdk-group-name": "page-rules.settings", "x-fern-sdk-method-name": "list", "x-forge-hidden": true, "x-forge-sunset": {"date": "2030-01-01T00:00:00Z"}}
```
