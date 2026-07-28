---
title: List Leaked Credential Checks Custom Detections
page_id: operation-get-zones-zone-id-leaked-credential-checks-detections-54f86ad3
path: operations/leaked-credential-checks
description: List user-defined detection patterns for Leaked Credential Checks.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/leaked-credential-checks/detections
operation_ids:
    - waf-product-api-leaked-credentials-list-detections
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Leaked Credential Checks Custom Detections

`GET /zones/{zone_id}/leaked-credential-checks/detections`

Operation ID: `waf-product-api-leaked-credentials-list-detections`

List user-defined detection patterns for Leaked Credential Checks.

## Definition

```yaml
{"operationId": "waf-product-api-leaked-credentials-list-detections", "summary": "List Leaked Credential Checks Custom Detections", "description": "List user-defined detection patterns for Leaked Credential Checks.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/waf-product-api-bundle_identifier"}}], "responses": {"200": {"description": "List Leaked Credential Checks custom detections response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/waf-product-api-bundle_response-custom-detection-collection"}}}}, "4XX": {"description": "List Leaked Credential Checks custom detections failure response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/waf-product-api-bundle_response-custom-detection-collection"}, {"$ref": "#/components/schemas/waf-product-api-bundle_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Leaked Credential Checks"], "x-api-token-group": ["Zone WAF Write", "Zone WAF Read", "Account WAF Write", "Account WAF Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "leaked-credential-checks.detections", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
