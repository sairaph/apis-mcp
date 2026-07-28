---
title: Disable Outgoing Zone Transfers
page_id: operation-post-zones-zone-id-secondary-dns-outgoing-disable-217020c8
path: operations/secondary-dns-primary-zone
description: Disable outgoing zone transfers for primary zone and clears IXFR backlog of primary zone.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /zones/{zone_id}/secondary_dns/outgoing/disable
operation_ids:
    - secondary-dns-(-primary-zone)-disable-outgoing-zone-transfers
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Disable Outgoing Zone Transfers

`POST /zones/{zone_id}/secondary_dns/outgoing/disable`

Operation ID: `secondary-dns-(-primary-zone)-disable-outgoing-zone-transfers`

Disable outgoing zone transfers for primary zone and clears IXFR backlog of primary zone.

## Definition

```yaml
{"operationId": "secondary-dns-(-primary-zone)-disable-outgoing-zone-transfers", "summary": "Disable Outgoing Zone Transfers", "description": "Disable outgoing zone transfers for primary zone and clears IXFR backlog of primary zone.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/secondary-dns_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {}}}, "responses": {"200": {"description": "Disable Outgoing Zone Transfers response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/secondary-dns_disable_transfer_response"}}}}, "4XX": {"description": "Disable Outgoing Zone Transfers response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/secondary-dns_disable_transfer_response"}, {"$ref": "#/components/schemas/secondary-dns_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Secondary DNS (Primary Zone)"], "x-api-token-group": ["Zone Settings Write", "Zone Write", "DNS Write"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "dns.zone-transfers.outgoing", "x-fern-sdk-method-name": "disable", "x-forge-require-confirmation": "This operation turns off outgoing zone transfers for the primary zone, which means secondary zones will stop receiving updates."}
```
