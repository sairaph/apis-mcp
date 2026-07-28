---
title: Get Outgoing Zone Transfer Status
page_id: operation-get-zones-zone-id-secondary-dns-outgoing-status-4195a37c
path: operations/secondary-dns-primary-zone
description: Get primary zone transfer status.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/secondary_dns/outgoing/status
operation_ids:
    - secondary-dns-(-primary-zone)-get-outgoing-zone-transfer-status
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Outgoing Zone Transfer Status

`GET /zones/{zone_id}/secondary_dns/outgoing/status`

Operation ID: `secondary-dns-(-primary-zone)-get-outgoing-zone-transfer-status`

Get primary zone transfer status.

## Definition

```yaml
{"operationId": "secondary-dns-(-primary-zone)-get-outgoing-zone-transfer-status", "summary": "Get Outgoing Zone Transfer Status", "description": "Get primary zone transfer status.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/secondary-dns_identifier"}}], "responses": {"200": {"description": "Get Outgoing Zone Transfer Status response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/secondary-dns_enable_transfer_response"}}}}, "4XX": {"description": "Get Outgoing Zone Transfer Status response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/secondary-dns_enable_transfer_response"}, {"$ref": "#/components/schemas/secondary-dns_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Secondary DNS (Primary Zone)"], "x-api-token-group": ["Zone Settings Write", "Zone Settings Read", "DNS Read", "Zone Write", "DNS Write"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "dns.zone-transfers.outgoing.status", "x-fern-sdk-method-name": "get"}
```
