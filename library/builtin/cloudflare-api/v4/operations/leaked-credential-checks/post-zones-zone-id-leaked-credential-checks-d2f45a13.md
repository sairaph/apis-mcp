---
title: Set Leaked Credential Checks Status
page_id: operation-post-zones-zone-id-leaked-credential-checks-3b1670e0
path: operations/leaked-credential-checks
description: Updates the current status of Leaked Credential Checks.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /zones/{zone_id}/leaked-credential-checks
operation_ids:
    - waf-product-api-leaked-credentials-set-status
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Set Leaked Credential Checks Status

`POST /zones/{zone_id}/leaked-credential-checks`

Operation ID: `waf-product-api-leaked-credentials-set-status`

Updates the current status of Leaked Credential Checks.

## Definition

```yaml
{"operationId": "waf-product-api-leaked-credentials-set-status", "summary": "Set Leaked Credential Checks Status", "description": "Updates the current status of Leaked Credential Checks.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/waf-product-api-bundle_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/waf-product-api-bundle_status"}}}}, "responses": {"200": {"description": "Set Leaked Credential Checks status response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/waf-product-api-bundle_response-status"}}}}, "4XX": {"description": "Set Leaked Credential Checks status failure response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/waf-product-api-bundle_response-status"}, {"$ref": "#/components/schemas/waf-product-api-bundle_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Leaked Credential Checks"], "x-api-token-group": ["Zone WAF Write", "Account WAF Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "leaked-credential-checks", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
