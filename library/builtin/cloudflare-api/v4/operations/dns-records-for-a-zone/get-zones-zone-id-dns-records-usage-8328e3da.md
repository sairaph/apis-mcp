---
title: Get DNS Record Usage
page_id: operation-get-zones-zone-id-dns-records-usage-7b03bbfd
path: operations/dns-records-for-a-zone
description: Get the current DNS record usage for a zone, including the number of records and the quota limit.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/dns_records/usage
operation_ids:
    - dns-records-for-a-zone-get-usage
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get DNS Record Usage

`GET /zones/{zone_id}/dns_records/usage`

Operation ID: `dns-records-for-a-zone-get-usage`

Get the current DNS record usage for a zone, including the number of records and the quota limit.

## Definition

```yaml
{"operationId": "dns-records-for-a-zone-get-usage", "summary": "Get DNS Record Usage", "description": "Get the current DNS record usage for a zone, including the number of records and the quota limit.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/dns-records_identifier"}}], "responses": {"200": {"description": "Get DNS Record Usage response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dns-records_dns_response_zone_usage"}}}}, "4XX": {"description": "Get DNS Record Usage response failure.", "content": {"application/json": {"schema": {"type": "object", "allOf": [{"$ref": "#/components/schemas/dns-records_dns_response_zone_usage"}, {"$ref": "#/components/schemas/dns-records_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["DNS Records for a Zone"], "x-api-token-group": ["Zone DNS Settings Write", "Zone DNS Settings Read", "DNS Read", "DNS Write"], "x-cfPermissionsRequired": {"enum": ["#dns_records:read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "dns.usage.zone", "x-fern-sdk-method-name": "get"}
```
