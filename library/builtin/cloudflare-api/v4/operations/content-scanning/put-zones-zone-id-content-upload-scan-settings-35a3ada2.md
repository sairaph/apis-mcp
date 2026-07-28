---
title: Update Content Scanning Status
page_id: operation-put-zones-zone-id-content-upload-scan-settings-e00df660
path: operations/content-scanning
description: Update the Content Scanning status.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /zones/{zone_id}/content-upload-scan/settings
operation_ids:
    - waf-content-scanning-update-settings
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update Content Scanning Status

`PUT /zones/{zone_id}/content-upload-scan/settings`

Operation ID: `waf-content-scanning-update-settings`

Update the Content Scanning status.

## Definition

```yaml
{"operationId": "waf-content-scanning-update-settings", "summary": "Update Content Scanning Status", "description": "Update the Content Scanning status.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/waf-product-api-bundle_identifier"}}], "requestBody": {"description": "Content Scanning settings to update.", "required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"value": {"description": "The status value for Content Scanning.", "type": "string", "example": "enabled", "enum": ["enabled", "disabled"]}}, "required": ["value"]}}}}, "responses": {"200": {"description": "Update Content Scanning settings response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/waf-product-api-bundle_response-status-2"}}}}, "4XX": {"description": "Update Content Scanning settings failure response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/waf-product-api-bundle_api-response-common-failure-2"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Content Scanning"], "x-api-token-group": ["Zone WAF Write", "Account WAF Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "content-scanning", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
