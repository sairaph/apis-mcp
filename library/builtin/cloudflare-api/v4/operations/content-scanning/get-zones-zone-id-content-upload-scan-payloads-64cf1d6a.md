---
title: List Existing Custom Scan Expressions
page_id: operation-get-zones-zone-id-content-upload-scan-payloads-67837dfc
path: operations/content-scanning
description: Get a list of existing custom scan expressions for Content Scanning.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/content-upload-scan/payloads
operation_ids:
    - waf-content-scanning-list-custom-scan-expressions
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Existing Custom Scan Expressions

`GET /zones/{zone_id}/content-upload-scan/payloads`

Operation ID: `waf-content-scanning-list-custom-scan-expressions`

Get a list of existing custom scan expressions for Content Scanning.

## Definition

```yaml
{"operationId": "waf-content-scanning-list-custom-scan-expressions", "summary": "List Existing Custom Scan Expressions", "description": "Get a list of existing custom scan expressions for Content Scanning.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/waf-product-api-bundle_identifier"}}], "responses": {"200": {"description": "List existing Content Scan custom scan expressions response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/waf-product-api-bundle_response-custom-scan-collection"}}}}, "4XX": {"description": "List existing Content Scan custom scan expressions failure response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/waf-product-api-bundle_response-custom-scan-collection"}, {"$ref": "#/components/schemas/waf-product-api-bundle_api-response-common-failure-2"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Content Scanning"], "x-api-token-group": ["Zone WAF Write", "Zone WAF Read", "Account WAF Write", "Account WAF Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "content-scanning.payloads", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
