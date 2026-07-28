---
title: Secondary Zone Configuration Details
page_id: operation-get-zones-zone-id-secondary-dns-incoming-f1b1c9d8
path: operations/secondary-dns-secondary-zone
description: Get secondary zone configuration for incoming zone transfers.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/secondary_dns/incoming
operation_ids:
    - secondary-dns-(-secondary-zone)-secondary-zone-configuration-details
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Secondary Zone Configuration Details

`GET /zones/{zone_id}/secondary_dns/incoming`

Operation ID: `secondary-dns-(-secondary-zone)-secondary-zone-configuration-details`

Get secondary zone configuration for incoming zone transfers.

## Definition

```yaml
{"operationId": "secondary-dns-(-secondary-zone)-secondary-zone-configuration-details", "summary": "Secondary Zone Configuration Details", "description": "Get secondary zone configuration for incoming zone transfers.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/secondary-dns_identifier"}}], "responses": {"200": {"description": "Secondary Zone Configuration Details response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/secondary-dns_single_response_incoming"}}}}, "4XX": {"description": "Secondary Zone Configuration Details response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/secondary-dns_single_response_incoming"}, {"$ref": "#/components/schemas/secondary-dns_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Secondary DNS (Secondary Zone)"], "x-api-token-group": ["Zone Settings Write", "Zone Settings Read", "DNS Read", "Zone Write", "DNS Write"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "dns.zone-transfers.incoming", "x-fern-sdk-method-name": "get"}
```
