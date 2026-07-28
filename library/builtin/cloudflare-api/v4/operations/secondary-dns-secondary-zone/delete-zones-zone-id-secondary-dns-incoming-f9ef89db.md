---
title: Delete Secondary Zone Configuration
page_id: operation-delete-zones-zone-id-secondary-dns-incoming-c367257d
path: operations/secondary-dns-secondary-zone
description: Delete secondary zone configuration for incoming zone transfers.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /zones/{zone_id}/secondary_dns/incoming
operation_ids:
    - secondary-dns-(-secondary-zone)-delete-secondary-zone-configuration
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete Secondary Zone Configuration

`DELETE /zones/{zone_id}/secondary_dns/incoming`

Operation ID: `secondary-dns-(-secondary-zone)-delete-secondary-zone-configuration`

Delete secondary zone configuration for incoming zone transfers.

## Definition

```yaml
{"operationId": "secondary-dns-(-secondary-zone)-delete-secondary-zone-configuration", "summary": "Delete Secondary Zone Configuration", "description": "Delete secondary zone configuration for incoming zone transfers.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/secondary-dns_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {}}}, "responses": {"200": {"description": "Delete Secondary Zone Configuration response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/secondary-dns_id_response"}}}}, "4XX": {"description": "Delete Secondary Zone Configuration response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/secondary-dns_id_response"}, {"$ref": "#/components/schemas/secondary-dns_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Secondary DNS (Secondary Zone)"], "x-api-token-group": ["Zone Settings Write", "Zone Write", "DNS Write"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "dns.zone-transfers.incoming", "x-fern-sdk-method-name": "delete"}
```
