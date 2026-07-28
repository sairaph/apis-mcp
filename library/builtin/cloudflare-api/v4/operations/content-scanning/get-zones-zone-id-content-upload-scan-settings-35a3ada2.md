---
title: Get Content Scanning Status
page_id: operation-get-zones-zone-id-content-upload-scan-settings-fcba5dcd
path: operations/content-scanning
description: Retrieve the current status of Content Scanning.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/content-upload-scan/settings
operation_ids:
    - waf-content-scanning-get-status
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Content Scanning Status

`GET /zones/{zone_id}/content-upload-scan/settings`

Operation ID: `waf-content-scanning-get-status`

Retrieve the current status of Content Scanning.

## Definition

```yaml
{"operationId": "waf-content-scanning-get-status", "summary": "Get Content Scanning Status", "description": "Retrieve the current status of Content Scanning.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/waf-product-api-bundle_identifier"}}], "responses": {"200": {"description": "Get Content Scanning status response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/waf-product-api-bundle_response-status-2"}}}}, "4XX": {"description": "Get Content Scanning status failure response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/waf-product-api-bundle_api-response-common-failure-2"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Content Scanning"], "x-api-token-group": ["Zone WAF Write", "Zone WAF Read", "Account WAF Write", "Account WAF Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "content-scanning.settings", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
