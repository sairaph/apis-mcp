---
title: Review Scanned DNS Records
page_id: operation-post-zones-zone-id-dns-records-scan-review-809eccc3
path: operations/dns-records-for-a-zone
description: Accept or reject DNS records found by the DNS records scan. Accepted records will be permanently added to the zone, while rejected records will be permanently deleted.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /zones/{zone_id}/dns_records/scan/review
operation_ids:
    - dns-records-for-a-zone-apply-dns-scan-results
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Review Scanned DNS Records

`POST /zones/{zone_id}/dns_records/scan/review`

Operation ID: `dns-records-for-a-zone-apply-dns-scan-results`

Accept or reject DNS records found by the DNS records scan. Accepted records will be permanently added to the zone, while rejected records will be permanently deleted.

## Definition

```yaml
{"operationId": "dns-records-for-a-zone-apply-dns-scan-results", "summary": "Review Scanned DNS Records", "description": "Accept or reject DNS records found by the DNS records scan. Accepted records will be permanently added to the zone, while rejected records will be permanently deleted.\n", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/dns-records_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dns-records_dns-request-review-scan-object"}}}}, "responses": {"200": {"description": "Records reviewed successfully", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dns-records_dns_response_review_scan"}}}}, "4XX": {"description": "Review failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dns-records_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["DNS Records for a Zone"], "x-api-token-group": ["DNS Write"], "x-cfPermissionsRequired": {"enum": ["#dns_records:edit"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "dns.records", "x-fern-sdk-method-name": "scan-review"}
```
