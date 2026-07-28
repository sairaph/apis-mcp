---
title: List Scanned DNS Records
page_id: operation-get-zones-zone-id-dns-records-scan-review-f115d580
path: operations/dns-records-for-a-zone
description: Retrieves the list of DNS records discovered up to this point by the asynchronous scan. These records are temporary until explicitly accepted or rejected via `POST /scan/review`. Additional records may be discovered by the scan later.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/dns_records/scan/review
operation_ids:
    - dns-records-for-a-zone-review-dns-scan
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Scanned DNS Records

`GET /zones/{zone_id}/dns_records/scan/review`

Operation ID: `dns-records-for-a-zone-review-dns-scan`

Retrieves the list of DNS records discovered up to this point by the asynchronous scan. These records are temporary until explicitly accepted or rejected via `POST /scan/review`. Additional records may be discovered by the scan later.

## Definition

```yaml
{"operationId": "dns-records-for-a-zone-review-dns-scan", "summary": "List Scanned DNS Records", "description": "Retrieves the list of DNS records discovered up to this point by the asynchronous scan. These records are temporary until explicitly accepted or rejected via `POST /scan/review`. Additional records may be discovered by the scan later.\n", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/dns-records_identifier"}}], "responses": {"200": {"description": "List of discovered DNS records", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dns-records_dns_response_collection"}}}}, "4XX": {"description": "Scan review failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dns-records_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["DNS Records for a Zone"], "x-api-token-group": ["DNS Read", "DNS Write"], "x-cfPermissionsRequired": {"enum": ["#dns_records:read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "dns.records", "x-fern-sdk-method-name": "scan-list"}
```
