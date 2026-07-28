---
title: Get AI Security for Apps Status
page_id: operation-get-zones-zone-id-ai-security-settings-8c72c075
path: operations/ai-security-for-apps
description: Get whether AI Security for Apps is enabled or disabled for a zone.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/ai-security/settings
operation_ids:
    - ai-security-settings-get
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get AI Security for Apps Status

`GET /zones/{zone_id}/ai-security/settings`

Operation ID: `ai-security-settings-get`

Get whether AI Security for Apps is enabled or disabled for a zone.

## Definition

```yaml
{"operationId": "ai-security-settings-get", "summary": "Get AI Security for Apps Status", "description": "Get whether AI Security for Apps is enabled or disabled for a zone.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/waf-product-api-bundle_zone_id"}}], "responses": {"200": {"description": "Get AI Security for Apps status response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/waf-product-api-bundle_response-settings"}}}}, "4XX": {"description": "Get AI Security for Apps status failure response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/waf-product-api-bundle_response-settings"}, {"$ref": "#/components/schemas/waf-product-api-bundle_api-response-common-failure-3"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["AI Security for Apps"]}
```
