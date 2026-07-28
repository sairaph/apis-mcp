---
title: Disable Content Scanning
page_id: operation-post-zones-zone-id-content-upload-scan-disable-b98503eb
path: operations/content-scanning
description: Disable Content Scanning.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /zones/{zone_id}/content-upload-scan/disable
operation_ids:
    - waf-content-scanning-disable
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Disable Content Scanning

`POST /zones/{zone_id}/content-upload-scan/disable`

Operation ID: `waf-content-scanning-disable`

Disable Content Scanning.

## Definition

```yaml
{"operationId": "waf-content-scanning-disable", "summary": "Disable Content Scanning", "description": "Disable Content Scanning.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/waf-product-api-bundle_identifier"}}], "responses": {"200": {"description": "Disable Content Scanning response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/waf-product-api-bundle_api-response-common-2"}}}}, "4XX": {"description": "Disable Content Scanning failure response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/waf-product-api-bundle_api-response-common-failure-2"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Content Scanning"], "x-api-token-group": ["Zone WAF Write", "Account WAF Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "content-scanning", "x-fern-sdk-method-name": "disable", "x-forge-hidden": true, "x-forge-require-confirmation": "This operation turns off WAF content scanning for a zone."}
```
