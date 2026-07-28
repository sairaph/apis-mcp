---
title: Set AI Security for Apps Status
page_id: operation-put-zones-zone-id-ai-security-settings-c5442698
path: operations/ai-security-for-apps
description: |-
    Enable or disable AI Security for Apps for a zone.

    Changes can take up to a minute to propagate to the zone.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /zones/{zone_id}/ai-security/settings
operation_ids:
    - ai-security-settings-put
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Set AI Security for Apps Status

`PUT /zones/{zone_id}/ai-security/settings`

Operation ID: `ai-security-settings-put`

Enable or disable AI Security for Apps for a zone.

Changes can take up to a minute to propagate to the zone.

## Definition

```yaml
{"operationId": "ai-security-settings-put", "summary": "Set AI Security for Apps Status", "description": "Enable or disable AI Security for Apps for a zone.\n\nChanges can take up to a minute to propagate to the zone.\n", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/waf-product-api-bundle_zone_id"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/waf-product-api-bundle_settings"}}}}, "responses": {"200": {"description": "Set AI Security for Apps status response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/waf-product-api-bundle_response-settings"}}}}, "4XX": {"description": "Set AI Security for Apps status failure response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/waf-product-api-bundle_response-settings"}, {"$ref": "#/components/schemas/waf-product-api-bundle_api-response-common-failure-3"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["AI Security for Apps"]}
```
