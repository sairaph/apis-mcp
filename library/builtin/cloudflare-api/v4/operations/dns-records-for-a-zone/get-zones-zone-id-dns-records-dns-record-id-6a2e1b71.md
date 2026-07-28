---
title: DNS Record Details
page_id: operation-get-zones-zone-id-dns-records-dns-record-id-8e565396
path: operations/dns-records-for-a-zone
description: Retrieves details for a specific DNS record in the zone.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/dns_records/{dns_record_id}
operation_ids:
    - dns-records-for-a-zone-dns-record-details
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# DNS Record Details

`GET /zones/{zone_id}/dns_records/{dns_record_id}`

Operation ID: `dns-records-for-a-zone-dns-record-details`

Retrieves details for a specific DNS record in the zone.

## Definition

```yaml
{"operationId": "dns-records-for-a-zone-dns-record-details", "summary": "DNS Record Details", "description": "Retrieves details for a specific DNS record in the zone.", "parameters": [{"name": "dns_record_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/dns-records_identifier"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/dns-records_identifier"}}, {"$ref": "#/components/parameters/dns-records_include_shadow_metadata"}], "responses": {"200": {"description": "DNS Record Details response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dns-records_dns_response_single"}}}}, "4XX": {"description": "DNS Record Details response failure", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/dns-records_dns_response_single"}, {"$ref": "#/components/schemas/dns-records_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["DNS Records for a Zone"], "x-api-token-group": ["DNS Read", "DNS Write"], "x-cfPermissionsRequired": {"enum": ["#dns_records:read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "dns.records", "x-fern-sdk-method-name": "get"}
```
