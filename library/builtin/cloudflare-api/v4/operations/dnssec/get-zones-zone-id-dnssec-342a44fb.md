---
title: DNSSEC Details
page_id: operation-get-zones-zone-id-dnssec-16faa3b5
path: operations/dnssec
description: Details about DNSSEC status and configuration.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/dnssec
operation_ids:
    - dnssec-dnssec-details
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# DNSSEC Details

`GET /zones/{zone_id}/dnssec`

Operation ID: `dnssec-dnssec-details`

Details about DNSSEC status and configuration.

## Definition

```yaml
{"operationId": "dnssec-dnssec-details", "summary": "DNSSEC Details", "description": "Details about DNSSEC status and configuration.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/dnssec_identifier"}}], "responses": {"200": {"description": "DNSSEC Details response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dnssec_dnssec_response_single"}}}}, "4XX": {"description": "DNSSEC Details response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dnssec_dnssec_response_single"}, {"$ref": "#/components/schemas/dnssec_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["DNSSEC"], "x-api-token-group": ["DNS Read", "DNS Write"], "x-cfPermissionsRequired": {"enum": ["#zone_settings:read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "dns.dnssec", "x-fern-sdk-method-name": "get"}
```
