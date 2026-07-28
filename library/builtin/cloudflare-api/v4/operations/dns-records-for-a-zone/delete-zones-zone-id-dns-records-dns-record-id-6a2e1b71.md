---
title: Delete DNS Record
page_id: operation-delete-zones-zone-id-dns-records-dns-record-id-92fe6d70
path: operations/dns-records-for-a-zone
description: Permanently removes a DNS record from the zone.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /zones/{zone_id}/dns_records/{dns_record_id}
operation_ids:
    - dns-records-for-a-zone-delete-dns-record
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete DNS Record

`DELETE /zones/{zone_id}/dns_records/{dns_record_id}`

Operation ID: `dns-records-for-a-zone-delete-dns-record`

Permanently removes a DNS record from the zone.

## Definition

```yaml
{"operationId": "dns-records-for-a-zone-delete-dns-record", "summary": "Delete DNS Record", "description": "Permanently removes a DNS record from the zone.", "parameters": [{"name": "dns_record_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/dns-records_identifier"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/dns-records_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {}}}, "responses": {"200": {"description": "Delete DNS Record response", "content": {"application/json": {"schema": {"type": "object", "properties": {"result": {"type": "object", "properties": {"id": {"$ref": "#/components/schemas/dns-records_identifier"}}}}}}}}, "4XX": {"description": "Delete DNS Record response failure", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"properties": {"result": {"type": "object", "properties": {"id": {"$ref": "#/components/schemas/dns-records_identifier"}}}}, "type": "object"}, {"$ref": "#/components/schemas/dns-records_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["DNS Records for a Zone"], "x-api-token-group": ["DNS Write"], "x-cfPermissionsRequired": {"enum": ["#dns_records:edit"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "dns.records", "x-fern-sdk-method-name": "delete"}
```
