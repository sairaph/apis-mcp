---
title: Change Origin TLS Compliance Modes setting
page_id: operation-patch-zones-zone-id-settings-origin-tls-compliance-modes-4c4a4720
path: operations/origin-tls
description: 'Update the set of TLS compliance modes for the zone. PATCH performs a full replace of the modes list, not a merge — the request body is treated as the complete new list, and any modes not present in it are removed. (To remove a single mode from an existing configuration, send the updated list without it.) The request body must be of the form `{"value": ["fips", "pqh"]}`. Currently supported modes are `fips` and `pqh`; an empty list clears the constraint. Future modes (e.g. `cnsa2`) may be added; clients should treat unknown values as opaque strings. Invalid mode values are rejected with a 4xx response.'
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /zones/{zone_id}/settings/origin_tls_compliance_modes
operation_ids:
    - zone-cache-settings-change-origin-tls-compliance-modes-setting
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Change Origin TLS Compliance Modes setting

`PATCH /zones/{zone_id}/settings/origin_tls_compliance_modes`

Operation ID: `zone-cache-settings-change-origin-tls-compliance-modes-setting`

Update the set of TLS compliance modes for the zone. PATCH performs a full replace of the modes list, not a merge — the request body is treated as the complete new list, and any modes not present in it are removed. (To remove a single mode from an existing configuration, send the updated list without it.) The request body must be of the form `{"value": ["fips", "pqh"]}`. Currently supported modes are `fips` and `pqh`; an empty list clears the constraint. Future modes (e.g. `cnsa2`) may be added; clients should treat unknown values as opaque strings. Invalid mode values are rejected with a 4xx response.

## Definition

```yaml
{"operationId": "zone-cache-settings-change-origin-tls-compliance-modes-setting", "summary": "Change Origin TLS Compliance Modes setting", "description": "Update the set of TLS compliance modes for the zone. PATCH performs a full replace of the modes list, not a merge — the request body is treated as the complete new list, and any modes not present in it are removed. (To remove a single mode from an existing configuration, send the updated list without it.) The request body must be of the form `{\"value\": [\"fips\", \"pqh\"]}`. Currently supported modes are `fips` and `pqh`; an empty list clears the constraint. Future modes (e.g. `cnsa2`) may be added; clients should treat unknown values as opaque strings. Invalid mode values are rejected with a 4xx response.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cache-rules_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"value": {"$ref": "#/components/schemas/cache-rules_origin_tls_compliance_modes_value"}}, "required": ["value"]}}}}, "responses": {"200": {"description": "Change Origin TLS Compliance Modes setting response.", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/cache-rules_zone_cache_settings_response_single"}, {"$ref": "#/components/schemas/cache-rules_origin_tls_compliance_modes_response_value"}]}}}}, "4XX": {"description": "Change Origin TLS Compliance Modes setting response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cache-rules_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Origin TLS"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
