---
title: Trigger DNS Record Scan
page_id: operation-post-zones-zone-id-dns-records-scan-trigger-4c042192
path: operations/dns-records-for-a-zone
description: Initiates an asynchronous scan for common DNS records on your domain. Note that this **does not** automatically add records to your zone. The scan runs in the background, and results can be reviewed later using the `/scan/review` endpoints. Useful if you haven't updated your nameservers yet.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /zones/{zone_id}/dns_records/scan/trigger
operation_ids:
    - dns-records-for-a-zone-trigger-dns-scan
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Trigger DNS Record Scan

`POST /zones/{zone_id}/dns_records/scan/trigger`

Operation ID: `dns-records-for-a-zone-trigger-dns-scan`

Initiates an asynchronous scan for common DNS records on your domain. Note that this **does not** automatically add records to your zone. The scan runs in the background, and results can be reviewed later using the `/scan/review` endpoints. Useful if you haven't updated your nameservers yet.

## Definition

```yaml
{"operationId": "dns-records-for-a-zone-trigger-dns-scan", "summary": "Trigger DNS Record Scan", "description": "Initiates an asynchronous scan for common DNS records on your domain. Note that this **does not** automatically add records to your zone. The scan runs in the background, and results can be reviewed later using the `/scan/review` endpoints. Useful if you haven't updated your nameservers yet.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/dns-records_identifier"}}], "requestBody": {"content": {"application/json": {}}}, "responses": {"200": {"description": "Trigger DNS Records Scan Response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dns-records_dns_response_trigger_scan"}}}}, "4XX": {"description": "Trigger DNS Records Scan response failure", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/dns-records_dns_response_trigger_scan"}, {"$ref": "#/components/schemas/dns-records_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["DNS Records for a Zone"], "x-api-token-group": ["DNS Write"], "x-cfPermissionsRequired": {"enum": ["#dns_records:edit"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "dns.records", "x-fern-sdk-method-name": "scan-trigger"}
```
