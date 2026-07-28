---
title: Primary Zone Configuration Details
page_id: operation-get-zones-zone-id-secondary-dns-outgoing-ab187039
path: operations/secondary-dns-primary-zone
description: Get primary zone configuration for outgoing zone transfers.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/secondary_dns/outgoing
operation_ids:
    - secondary-dns-(-primary-zone)-primary-zone-configuration-details
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Primary Zone Configuration Details

`GET /zones/{zone_id}/secondary_dns/outgoing`

Operation ID: `secondary-dns-(-primary-zone)-primary-zone-configuration-details`

Get primary zone configuration for outgoing zone transfers.

## Definition

```yaml
{"operationId": "secondary-dns-(-primary-zone)-primary-zone-configuration-details", "summary": "Primary Zone Configuration Details", "description": "Get primary zone configuration for outgoing zone transfers.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/secondary-dns_identifier"}}], "responses": {"200": {"description": "Primary Zone Configuration Details response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/secondary-dns_single_response_outgoing"}}}}, "4XX": {"description": "Primary Zone Configuration Details response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/secondary-dns_single_response_outgoing"}, {"$ref": "#/components/schemas/secondary-dns_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Secondary DNS (Primary Zone)"], "x-api-token-group": ["Zone Settings Write", "Zone Settings Read", "DNS Read", "Zone Write", "DNS Write"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "dns.zone-transfers.outgoing", "x-fern-sdk-method-name": "get"}
```
