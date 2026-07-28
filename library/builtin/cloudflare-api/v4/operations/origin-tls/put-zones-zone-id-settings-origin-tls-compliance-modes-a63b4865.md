---
title: Replace Origin TLS Compliance Modes setting
page_id: operation-put-zones-zone-id-settings-origin-tls-compliance-modes-8a121051
path: operations/origin-tls
description: 'Replace the entire set of TLS compliance modes for the zone with the list provided in the request body. PUT performs a full replace, not a merge — any modes not present in the request body are removed. The request body must be of the form `{"value": ["fips", "pqh"]}`. Currently supported modes are `fips` and `pqh`; an empty list clears the constraint. Future modes (e.g. `cnsa2`) may be added; clients should treat unknown values as opaque strings. Invalid mode values are rejected with a 4xx response.'
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /zones/{zone_id}/settings/origin_tls_compliance_modes
operation_ids:
    - zone-cache-settings-replace-origin-tls-compliance-modes-setting
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Replace Origin TLS Compliance Modes setting

`PUT /zones/{zone_id}/settings/origin_tls_compliance_modes`

Operation ID: `zone-cache-settings-replace-origin-tls-compliance-modes-setting`

Replace the entire set of TLS compliance modes for the zone with the list provided in the request body. PUT performs a full replace, not a merge — any modes not present in the request body are removed. The request body must be of the form `{"value": ["fips", "pqh"]}`. Currently supported modes are `fips` and `pqh`; an empty list clears the constraint. Future modes (e.g. `cnsa2`) may be added; clients should treat unknown values as opaque strings. Invalid mode values are rejected with a 4xx response.

## Definition

```yaml
{"operationId": "zone-cache-settings-replace-origin-tls-compliance-modes-setting", "summary": "Replace Origin TLS Compliance Modes setting", "description": "Replace the entire set of TLS compliance modes for the zone with the list provided in the request body. PUT performs a full replace, not a merge — any modes not present in the request body are removed. The request body must be of the form `{\"value\": [\"fips\", \"pqh\"]}`. Currently supported modes are `fips` and `pqh`; an empty list clears the constraint. Future modes (e.g. `cnsa2`) may be added; clients should treat unknown values as opaque strings. Invalid mode values are rejected with a 4xx response.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cache-rules_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"value": {"$ref": "#/components/schemas/cache-rules_origin_tls_compliance_modes_value"}}, "required": ["value"]}}}}, "responses": {"200": {"description": "Replace Origin TLS Compliance Modes setting response.", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/cache-rules_zone_cache_settings_response_single"}, {"$ref": "#/components/schemas/cache-rules_origin_tls_compliance_modes_response_value"}]}}}}, "4XX": {"description": "Replace Origin TLS Compliance Modes setting response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cache-rules_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Origin TLS"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
