---
title: Create DNS Record
page_id: operation-post-zones-zone-id-dns-records-377203ce
path: operations/dns-records-for-a-zone
description: |-
    Create a new DNS record for a zone.

    Notes:
    - A/AAAA records cannot exist on the same name as CNAME records.
    - NS records cannot exist on the same name as any other record type.
    - Domain names are always represented in Punycode, even if Unicode
      characters were used when creating the record.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /zones/{zone_id}/dns_records
operation_ids:
    - dns-records-for-a-zone-create-dns-record
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create DNS Record

`POST /zones/{zone_id}/dns_records`

Operation ID: `dns-records-for-a-zone-create-dns-record`

Create a new DNS record for a zone.

Notes:
- A/AAAA records cannot exist on the same name as CNAME records.
- NS records cannot exist on the same name as any other record type.
- Domain names are always represented in Punycode, even if Unicode
  characters were used when creating the record.

## Definition

```yaml
{"operationId": "dns-records-for-a-zone-create-dns-record", "summary": "Create DNS Record", "description": "Create a new DNS record for a zone.\n\nNotes:\n- A/AAAA records cannot exist on the same name as CNAME records.\n- NS records cannot exist on the same name as any other record type.\n- Domain names are always represented in Punycode, even if Unicode\n  characters were used when creating the record.\n", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/dns-records_identifier"}}, {"$ref": "#/components/parameters/dns-records_include_shadow_metadata"}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dns-records_dns-record-post"}}}}, "responses": {"200": {"description": "Create DNS Record response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dns-records_dns_response_single"}}}}, "4XX": {"description": "Create DNS Record response failure", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/dns-records_dns_response_single"}, {"$ref": "#/components/schemas/dns-records_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["DNS Records for a Zone"], "x-api-token-group": ["DNS Write"], "x-cfPermissionsRequired": {"enum": ["#dns_records:edit"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "dns.records", "x-fern-sdk-method-name": "create"}
```
