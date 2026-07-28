---
title: Enable Content Scanning
page_id: operation-post-zones-zone-id-content-upload-scan-enable-3fa413a3
path: operations/content-scanning
description: Enable Content Scanning.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /zones/{zone_id}/content-upload-scan/enable
operation_ids:
    - waf-content-scanning-enable
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Enable Content Scanning

`POST /zones/{zone_id}/content-upload-scan/enable`

Operation ID: `waf-content-scanning-enable`

Enable Content Scanning.

## Definition

```yaml
{"operationId": "waf-content-scanning-enable", "summary": "Enable Content Scanning", "description": "Enable Content Scanning.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/waf-product-api-bundle_identifier"}}], "responses": {"200": {"description": "Enable Content Scanning response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/waf-product-api-bundle_api-response-common-2"}}}}, "4XX": {"description": "Enable Content Scanning failure response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/waf-product-api-bundle_api-response-common-failure-2"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Content Scanning"], "x-api-token-group": ["Zone WAF Write", "Account WAF Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "content-scanning", "x-fern-sdk-method-name": "enable", "x-forge-hidden": true}
```
