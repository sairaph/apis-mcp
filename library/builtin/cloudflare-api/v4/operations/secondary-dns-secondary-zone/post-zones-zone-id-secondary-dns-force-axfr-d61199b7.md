---
title: Force AXFR
page_id: operation-post-zones-zone-id-secondary-dns-force-axfr-4cf5a5e4
path: operations/secondary-dns-secondary-zone
description: Sends AXFR zone transfer request to primary nameserver(s).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /zones/{zone_id}/secondary_dns/force_axfr
operation_ids:
    - secondary-dns-(-secondary-zone)-force-axfr
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Force AXFR

`POST /zones/{zone_id}/secondary_dns/force_axfr`

Operation ID: `secondary-dns-(-secondary-zone)-force-axfr`

Sends AXFR zone transfer request to primary nameserver(s).

## Definition

```yaml
{"operationId": "secondary-dns-(-secondary-zone)-force-axfr", "summary": "Force AXFR", "description": "Sends AXFR zone transfer request to primary nameserver(s).", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/secondary-dns_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {}}}, "responses": {"200": {"description": "Force AXFR response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/secondary-dns_force_response"}}}}, "4XX": {"description": "Force AXFR response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/secondary-dns_force_response"}, {"$ref": "#/components/schemas/secondary-dns_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Secondary DNS (Secondary Zone)"], "x-api-token-group": ["Zone Settings Write", "Zone Write", "DNS Write"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "dns.zone-transfers.force-axfr", "x-fern-sdk-method-name": "create"}
```
