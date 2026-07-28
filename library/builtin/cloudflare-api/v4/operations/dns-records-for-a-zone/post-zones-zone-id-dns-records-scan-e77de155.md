---
title: Scan DNS Records
page_id: operation-post-zones-zone-id-dns-records-scan-c6682792
path: operations/dns-records-for-a-zone
description: Scan for common DNS records on your domain and automatically add them to your zone. Useful if you haven't updated your nameservers yet.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /zones/{zone_id}/dns_records/scan
operation_ids:
    - dns-records-for-a-zone-scan-dns-records
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Scan DNS Records

`POST /zones/{zone_id}/dns_records/scan`

Operation ID: `dns-records-for-a-zone-scan-dns-records`

Scan for common DNS records on your domain and automatically add them to your zone. Useful if you haven't updated your nameservers yet.

## Definition

```yaml
{"operationId": "dns-records-for-a-zone-scan-dns-records", "summary": "Scan DNS Records", "description": "Scan for common DNS records on your domain and automatically add them to your zone. Useful if you haven't updated your nameservers yet.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/dns-records_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {}}}, "responses": {"200": {"description": "Scan DNS Records response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dns-records_dns_response_import_scan"}}}}, "4XX": {"description": "Scan DNS Records response failure", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/dns-records_dns_response_import_scan"}, {"$ref": "#/components/schemas/dns-records_api-response-common-failure"}]}}}}}, "deprecated": true, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["DNS Records for a Zone"], "x-api-token-group": ["DNS Write"], "x-cfPermissionsRequired": {"enum": ["#dns_records:edit"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "dns.records", "x-fern-sdk-method-name": "scan", "x-stainless-deprecation-message": "This endpoint is deprecated in favor of a new asynchronous version. Please use the [/scan/trigger](https://developers.cloudflare.com/api/resources/dns/subresources/records/methods/scan/trigger) and [/scan/review](https://developers.cloudflare.com/api/resources/dns/subresources/records/methods/scan/review) endpoints instead.\n"}
```
