---
title: Delete a Custom Scan Expression
page_id: operation-delete-zones-zone-id-content-upload-scan-payloads-expression-id-cd0ac970
path: operations/content-scanning
description: Delete a Content Scan Custom Expression.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /zones/{zone_id}/content-upload-scan/payloads/{expression_id}
operation_ids:
    - waf-content-scanning-delete-custom-scan-expressions
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete a Custom Scan Expression

`DELETE /zones/{zone_id}/content-upload-scan/payloads/{expression_id}`

Operation ID: `waf-content-scanning-delete-custom-scan-expressions`

Delete a Content Scan Custom Expression.

## Definition

```yaml
{"operationId": "waf-content-scanning-delete-custom-scan-expressions", "summary": "Delete a Custom Scan Expression", "description": "Delete a Content Scan Custom Expression.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/waf-product-api-bundle_identifier"}}, {"name": "expression_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/waf-product-api-bundle_custom-scan-id"}}], "responses": {"200": {"description": "Delete Content Scan custom scan expressions response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/waf-product-api-bundle_response-custom-scan-collection"}}}}, "4XX": {"description": "Delete Content Scan custom scan expressions failure response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/waf-product-api-bundle_response-custom-scan-collection"}, {"$ref": "#/components/schemas/waf-product-api-bundle_api-response-common-failure-2"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Content Scanning"], "x-api-token-group": ["Zone WAF Write", "Account WAF Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "content-scanning.payloads", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
