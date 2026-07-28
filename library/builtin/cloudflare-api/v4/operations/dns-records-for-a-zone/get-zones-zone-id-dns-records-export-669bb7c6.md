---
title: Export DNS Records
page_id: operation-get-zones-zone-id-dns-records-export-a2ee7a6e
path: operations/dns-records-for-a-zone
description: |-
    You can export your [BIND config](https://en.wikipedia.org/wiki/Zone_file "Zone file") through this endpoint.

    See [the documentation](https://developers.cloudflare.com/dns/manage-dns-records/how-to/import-and-export/ "Import and export records") for more information.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/dns_records/export
operation_ids:
    - dns-records-for-a-zone-export-dns-records
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Export DNS Records

`GET /zones/{zone_id}/dns_records/export`

Operation ID: `dns-records-for-a-zone-export-dns-records`

You can export your [BIND config](https://en.wikipedia.org/wiki/Zone_file "Zone file") through this endpoint.

See [the documentation](https://developers.cloudflare.com/dns/manage-dns-records/how-to/import-and-export/ "Import and export records") for more information.

## Definition

```yaml
{"operationId": "dns-records-for-a-zone-export-dns-records", "summary": "Export DNS Records", "description": "You can export your [BIND config](https://en.wikipedia.org/wiki/Zone_file \"Zone file\") through this endpoint.\n\nSee [the documentation](https://developers.cloudflare.com/dns/manage-dns-records/how-to/import-and-export/ \"Import and export records\") for more information.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/dns-records_identifier"}}], "responses": {"200": {"description": "Export DNS Records response", "content": {"text/plain": {"schema": {"description": "Exported BIND zone file.", "type": "string", "example": "www.example.com. 300 IN  A 127.0.0.1\n"}}}}, "4XX": {"description": "Export DNS Records response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dns-records_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["DNS Records for a Zone"], "x-api-token-group": ["DNS Read", "DNS Write"], "x-cfPermissionsRequired": {"enum": ["#dns_records:read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "dns.records", "x-fern-sdk-method-name": "export"}
```
