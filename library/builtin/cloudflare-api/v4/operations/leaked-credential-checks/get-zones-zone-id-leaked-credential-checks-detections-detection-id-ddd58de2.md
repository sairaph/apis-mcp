---
title: Get Leaked Credential Checks Custom Detection
page_id: operation-get-zones-zone-id-leaked-credential-checks-detections-detection-id-b4eeb191
path: operations/leaked-credential-checks
description: Get user-defined detection pattern for Leaked Credential Checks.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/leaked-credential-checks/detections/{detection_id}
operation_ids:
    - waf-product-api-leaked-credentials-get-detection
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Leaked Credential Checks Custom Detection

`GET /zones/{zone_id}/leaked-credential-checks/detections/{detection_id}`

Operation ID: `waf-product-api-leaked-credentials-get-detection`

Get user-defined detection pattern for Leaked Credential Checks.

## Definition

```yaml
{"operationId": "waf-product-api-leaked-credentials-get-detection", "summary": "Get Leaked Credential Checks Custom Detection", "description": "Get user-defined detection pattern for Leaked Credential Checks.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/waf-product-api-bundle_identifier"}}, {"name": "detection_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/waf-product-api-bundle_detection-id"}}], "responses": {"200": {"description": "Get Leaked Credential Checks custom detection response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/waf-product-api-bundle_response-custom-detection"}}}}, "4XX": {"description": "Update Leaked Credential Checks custom detection failure response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/waf-product-api-bundle_response-custom-detection"}, {"$ref": "#/components/schemas/waf-product-api-bundle_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Leaked Credential Checks"], "x-api-token-group": ["Zone WAF Write", "Zone WAF Read", "Account WAF Write", "Account WAF Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "leaked-credential-checks.detections", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
