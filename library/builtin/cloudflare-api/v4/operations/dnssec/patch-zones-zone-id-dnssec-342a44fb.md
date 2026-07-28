---
title: Edit DNSSEC Status
page_id: operation-patch-zones-zone-id-dnssec-60cc95df
path: operations/dnssec
description: Enable or disable DNSSEC.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /zones/{zone_id}/dnssec
operation_ids:
    - dnssec-edit-dnssec-status
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Edit DNSSEC Status

`PATCH /zones/{zone_id}/dnssec`

Operation ID: `dnssec-edit-dnssec-status`

Enable or disable DNSSEC.

## Definition

```yaml
{"operationId": "dnssec-edit-dnssec-status", "summary": "Edit DNSSEC Status", "description": "Enable or disable DNSSEC.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/dnssec_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"dnssec_multi_signer": {"$ref": "#/components/schemas/dnssec_dnssec_multi_signer"}, "dnssec_presigned": {"$ref": "#/components/schemas/dnssec_dnssec_presigned"}, "dnssec_use_nsec3": {"$ref": "#/components/schemas/dnssec_dnssec_use_nsec3"}, "status": {"description": "Status of DNSSEC, based on user-desired state and presence of necessary records.", "type": "string", "example": "active", "enum": ["active", "disabled"]}}}}}}, "responses": {"200": {"description": "Edit DNSSEC Status response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dnssec_dnssec_response_single"}}}}, "4XX": {"description": "Edit DNSSEC Status response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dnssec_dnssec_response_single"}, {"$ref": "#/components/schemas/dnssec_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["DNSSEC"], "x-api-token-group": ["DNS Write"], "x-cfPermissionsRequired": {"enum": ["#zone_settings:edit"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "dns.dnssec", "x-fern-sdk-method-name": "edit"}
```
