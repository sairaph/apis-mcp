---
title: Delete variants setting
page_id: operation-delete-zones-zone-id-cache-variants-3e9502f8
path: operations/zone-cache-settings
description: 'Variant support enables caching variants of images with certain file extensions in addition to the original. This only applies when the origin server sends the ''Vary: Accept'' response header. If the origin server sends ''Vary: Accept'' but does not serve the variant requested, the response will not be cached. This will be indicated with BYPASS cache status in the response headers.'
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /zones/{zone_id}/cache/variants
operation_ids:
    - zone-cache-settings-delete-variants-setting
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete variants setting

`DELETE /zones/{zone_id}/cache/variants`

Operation ID: `zone-cache-settings-delete-variants-setting`

Variant support enables caching variants of images with certain file extensions in addition to the original. This only applies when the origin server sends the 'Vary: Accept' response header. If the origin server sends 'Vary: Accept' but does not serve the variant requested, the response will not be cached. This will be indicated with BYPASS cache status in the response headers.

## Definition

```yaml
{"operationId": "zone-cache-settings-delete-variants-setting", "summary": "Delete variants setting", "description": "Variant support enables caching variants of images with certain file extensions in addition to the original. This only applies when the origin server sends the 'Vary: Accept' response header. If the origin server sends 'Vary: Accept' but does not serve the variant requested, the response will not be cached. This will be indicated with BYPASS cache status in the response headers.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cache-rules_identifier"}}], "responses": {"200": {"description": "Delete variants setting response.", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/cache-rules_zone_cache_settings_delete_response_single"}, {"properties": {"result": {"$ref": "#/components/schemas/cache-rules_variants"}}, "type": "object"}]}}}}, "4XX": {"description": "Delete variants setting response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cache-rules_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Zone Cache Settings"], "x-api-token-group": ["Zone Settings Write", "Zone Write"], "x-cfPermissionsRequired": {"enum": ["#zone_settings:edit"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": false, "pro": true}}
```
