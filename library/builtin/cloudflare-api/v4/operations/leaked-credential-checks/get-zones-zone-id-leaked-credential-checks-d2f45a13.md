---
title: Get Leaked Credential Checks Status
page_id: operation-get-zones-zone-id-leaked-credential-checks-ddd10a7f
path: operations/leaked-credential-checks
description: Retrieves the current status of Leaked Credential Checks.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/leaked-credential-checks
operation_ids:
    - waf-product-api-leaked-credentials-get-status
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Leaked Credential Checks Status

`GET /zones/{zone_id}/leaked-credential-checks`

Operation ID: `waf-product-api-leaked-credentials-get-status`

Retrieves the current status of Leaked Credential Checks.

## Definition

```yaml
{"operationId": "waf-product-api-leaked-credentials-get-status", "summary": "Get Leaked Credential Checks Status", "description": "Retrieves the current status of Leaked Credential Checks.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/waf-product-api-bundle_identifier"}}], "responses": {"200": {"description": "Get Leaked Credential Checks status response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/waf-product-api-bundle_response-status"}}}}, "4XX": {"description": "Get Leaked Credential Checks status failure response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/waf-product-api-bundle_response-status"}, {"$ref": "#/components/schemas/waf-product-api-bundle_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Leaked Credential Checks"], "x-api-token-group": ["Zone WAF Write", "Zone WAF Read", "Account WAF Write", "Account WAF Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "leaked-credential-checks", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
