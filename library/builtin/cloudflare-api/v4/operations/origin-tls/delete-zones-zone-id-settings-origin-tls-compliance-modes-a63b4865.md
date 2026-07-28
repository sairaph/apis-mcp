---
title: Delete Origin TLS Compliance Modes setting
page_id: operation-delete-zones-zone-id-settings-origin-tls-compliance-modes-5df1abdb
path: operations/origin-tls
description: Delete the Origin TLS Compliance Modes setting for the zone, removing any configured compliance constraint. After deletion, Cloudflare's default behavior applies (no compliance filtering of the key-exchange algorithm list sent to the origin).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /zones/{zone_id}/settings/origin_tls_compliance_modes
operation_ids:
    - zone-cache-settings-delete-origin-tls-compliance-modes-setting
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete Origin TLS Compliance Modes setting

`DELETE /zones/{zone_id}/settings/origin_tls_compliance_modes`

Operation ID: `zone-cache-settings-delete-origin-tls-compliance-modes-setting`

Delete the Origin TLS Compliance Modes setting for the zone, removing any configured compliance constraint. After deletion, Cloudflare's default behavior applies (no compliance filtering of the key-exchange algorithm list sent to the origin).

## Definition

```yaml
{"operationId": "zone-cache-settings-delete-origin-tls-compliance-modes-setting", "summary": "Delete Origin TLS Compliance Modes setting", "description": "Delete the Origin TLS Compliance Modes setting for the zone, removing any configured compliance constraint. After deletion, Cloudflare's default behavior applies (no compliance filtering of the key-exchange algorithm list sent to the origin).", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cache-rules_identifier"}}], "responses": {"200": {"description": "Delete Origin TLS Compliance Modes setting response.", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/cache-rules_zone_cache_settings_delete_response_single"}, {"properties": {"result": {"$ref": "#/components/schemas/cache-rules_origin_tls_compliance_modes"}}, "type": "object"}]}}}}, "4XX": {"description": "Delete Origin TLS Compliance Modes setting response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cache-rules_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Origin TLS"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
