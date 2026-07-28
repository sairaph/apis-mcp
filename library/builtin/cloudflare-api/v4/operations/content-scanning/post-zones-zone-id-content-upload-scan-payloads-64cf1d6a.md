---
title: Add Custom Scan Expressions
page_id: operation-post-zones-zone-id-content-upload-scan-payloads-ec8d9272
path: operations/content-scanning
description: Add custom scan expressions for Content Scanning.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /zones/{zone_id}/content-upload-scan/payloads
operation_ids:
    - waf-content-scanning-add-custom-scan-expressions
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Add Custom Scan Expressions

`POST /zones/{zone_id}/content-upload-scan/payloads`

Operation ID: `waf-content-scanning-add-custom-scan-expressions`

Add custom scan expressions for Content Scanning.

## Definition

```yaml
{"operationId": "waf-content-scanning-add-custom-scan-expressions", "summary": "Add Custom Scan Expressions", "description": "Add custom scan expressions for Content Scanning.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/waf-product-api-bundle_identifier"}}], "requestBody": {"description": "Array of custom scan expressions to add.", "required": true, "content": {"application/json": {"schema": {"type": "array", "items": {"properties": {"payload": {"$ref": "#/components/schemas/waf-product-api-bundle_custom-scan-payload"}}, "required": ["payload"], "type": "object"}}}}}, "responses": {"200": {"description": "Add custom scan expressions for Content Scanning.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/waf-product-api-bundle_response-custom-scan-collection"}}}}, "4XX": {"description": "List existing Content Scan custom scan expressions failure response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/waf-product-api-bundle_response-custom-scan-collection"}, {"$ref": "#/components/schemas/waf-product-api-bundle_api-response-common-failure-2"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Content Scanning"], "x-api-token-group": ["Zone WAF Write", "Account WAF Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "content-scanning.payloads", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
