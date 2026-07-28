---
title: List DNSSEC ZSKs
page_id: operation-get-zones-zone-id-dnssec-zsk-34a96a1c
path: operations/dnssec
description: List the Zone Signing Keys (ZSKs) that DNSSEC uses for the zone.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/dnssec/zsk
operation_ids:
    - dnssec-list-dnssec-zsks
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List DNSSEC ZSKs

`GET /zones/{zone_id}/dnssec/zsk`

Operation ID: `dnssec-list-dnssec-zsks`

List the Zone Signing Keys (ZSKs) that DNSSEC uses for the zone.

## Definition

```yaml
{"operationId": "dnssec-list-dnssec-zsks", "summary": "List DNSSEC ZSKs", "description": "List the Zone Signing Keys (ZSKs) that DNSSEC uses for the zone.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/dnssec_identifier"}}], "responses": {"200": {"description": "List DNSSEC ZSKs response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dnssec_dnssec_zsk_response_collection"}}}}, "4XX": {"description": "List DNSSEC ZSKs response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dnssec_dnssec_zsk_response_collection"}, {"$ref": "#/components/schemas/dnssec_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["DNSSEC"], "x-api-token-group": ["DNS Read", "DNS Write"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "dns.dnssec.list", "x-fern-sdk-method-name": "zsks"}
```
