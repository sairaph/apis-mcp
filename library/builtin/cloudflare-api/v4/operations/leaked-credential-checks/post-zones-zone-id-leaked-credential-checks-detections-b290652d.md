---
title: Create Leaked Credential Checks Custom Detection
page_id: operation-post-zones-zone-id-leaked-credential-checks-detections-48a4051f
path: operations/leaked-credential-checks
description: Create user-defined detection pattern for Leaked Credential Checks.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /zones/{zone_id}/leaked-credential-checks/detections
operation_ids:
    - waf-product-api-leaked-credentials-create-detection
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create Leaked Credential Checks Custom Detection

`POST /zones/{zone_id}/leaked-credential-checks/detections`

Operation ID: `waf-product-api-leaked-credentials-create-detection`

Create user-defined detection pattern for Leaked Credential Checks.

## Definition

```yaml
{"operationId": "waf-product-api-leaked-credentials-create-detection", "summary": "Create Leaked Credential Checks Custom Detection", "description": "Create user-defined detection pattern for Leaked Credential Checks.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/waf-product-api-bundle_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/waf-product-api-bundle_custom-detection"}}}}, "responses": {"200": {"description": "Create Leaked Credential Checks custom detection response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/waf-product-api-bundle_response-custom-detection"}}}}, "4XX": {"description": "Create Leaked Credential Checks custom detection failure response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/waf-product-api-bundle_response-custom-detection"}, {"$ref": "#/components/schemas/waf-product-api-bundle_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Leaked Credential Checks"], "x-api-token-group": ["Zone WAF Write", "Account WAF Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "leaked-credential-checks.detections", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
