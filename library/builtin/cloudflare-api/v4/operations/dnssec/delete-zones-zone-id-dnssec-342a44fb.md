---
title: Delete DNSSEC records
page_id: operation-delete-zones-zone-id-dnssec-533c0f01
path: operations/dnssec
description: Delete DNSSEC.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /zones/{zone_id}/dnssec
operation_ids:
    - dnssec-delete-dnssec-records
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete DNSSEC records

`DELETE /zones/{zone_id}/dnssec`

Operation ID: `dnssec-delete-dnssec-records`

Delete DNSSEC.

## Definition

```yaml
{"operationId": "dnssec-delete-dnssec-records", "summary": "Delete DNSSEC records", "description": "Delete DNSSEC.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/dnssec_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {}}}, "responses": {"200": {"description": "Delete DNSSEC records response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dnssec_delete_dnssec_response_single"}}}}, "4XX": {"description": "Delete DNSSEC records response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dnssec_delete_dnssec_response_single"}, {"$ref": "#/components/schemas/dnssec_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["DNSSEC"], "x-api-token-group": ["DNS Write"], "x-cfPermissionsRequired": {"enum": ["#zone_settings:edit"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "dns.dnssec", "x-fern-sdk-method-name": "delete"}
```
