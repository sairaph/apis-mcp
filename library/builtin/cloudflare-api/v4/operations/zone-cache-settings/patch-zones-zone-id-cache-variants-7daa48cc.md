---
title: Change variants setting
page_id: operation-patch-zones-zone-id-cache-variants-57ed4193
path: operations/zone-cache-settings
description: 'Variant support enables caching variants of images with certain file extensions in addition to the original. This only applies when the origin server sends the ''Vary: Accept'' response header. If the origin server sends ''Vary: Accept'' but does not serve the variant requested, the response will not be cached. This will be indicated with BYPASS cache status in the response headers.'
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /zones/{zone_id}/cache/variants
operation_ids:
    - zone-cache-settings-change-variants-setting
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Change variants setting

`PATCH /zones/{zone_id}/cache/variants`

Operation ID: `zone-cache-settings-change-variants-setting`

Variant support enables caching variants of images with certain file extensions in addition to the original. This only applies when the origin server sends the 'Vary: Accept' response header. If the origin server sends 'Vary: Accept' but does not serve the variant requested, the response will not be cached. This will be indicated with BYPASS cache status in the response headers.

## Definition

```yaml
{"operationId": "zone-cache-settings-change-variants-setting", "summary": "Change variants setting", "description": "Variant support enables caching variants of images with certain file extensions in addition to the original. This only applies when the origin server sends the 'Vary: Accept' response header. If the origin server sends 'Vary: Accept' but does not serve the variant requested, the response will not be cached. This will be indicated with BYPASS cache status in the response headers.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cache-rules_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"value": {"$ref": "#/components/schemas/cache-rules_variants_value"}}, "required": ["value"]}}}}, "responses": {"200": {"description": "Change variants setting response.", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/cache-rules_zone_cache_settings_response_single"}, {"$ref": "#/components/schemas/cache-rules_variants_response_value"}]}}}}, "4XX": {"description": "Change variants setting response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cache-rules_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Zone Cache Settings"], "x-api-token-group": ["Zone Settings Write", "Zone Write"], "x-cfPermissionsRequired": {"enum": ["#zone_settings:edit"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": true}}
```
