---
title: Delete Primary Zone Configuration
page_id: operation-delete-zones-zone-id-secondary-dns-outgoing-04f877b2
path: operations/secondary-dns-primary-zone
description: Delete primary zone configuration for outgoing zone transfers.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /zones/{zone_id}/secondary_dns/outgoing
operation_ids:
    - secondary-dns-(-primary-zone)-delete-primary-zone-configuration
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete Primary Zone Configuration

`DELETE /zones/{zone_id}/secondary_dns/outgoing`

Operation ID: `secondary-dns-(-primary-zone)-delete-primary-zone-configuration`

Delete primary zone configuration for outgoing zone transfers.

## Definition

```yaml
{"operationId": "secondary-dns-(-primary-zone)-delete-primary-zone-configuration", "summary": "Delete Primary Zone Configuration", "description": "Delete primary zone configuration for outgoing zone transfers.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/secondary-dns_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {}}}, "responses": {"200": {"description": "Delete Primary Zone Configuration response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/secondary-dns_id_response"}}}}, "4XX": {"description": "Delete Primary Zone Configuration response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/secondary-dns_id_response"}, {"$ref": "#/components/schemas/secondary-dns_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Secondary DNS (Primary Zone)"], "x-api-token-group": ["Zone Settings Write", "Zone Write", "DNS Write"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "dns.zone-transfers.outgoing", "x-fern-sdk-method-name": "delete"}
```
