---
title: Create Primary Zone Configuration
page_id: operation-post-zones-zone-id-secondary-dns-outgoing-af97ffd0
path: operations/secondary-dns-primary-zone
description: Create primary zone configuration for outgoing zone transfers.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /zones/{zone_id}/secondary_dns/outgoing
operation_ids:
    - secondary-dns-(-primary-zone)-create-primary-zone-configuration
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create Primary Zone Configuration

`POST /zones/{zone_id}/secondary_dns/outgoing`

Operation ID: `secondary-dns-(-primary-zone)-create-primary-zone-configuration`

Create primary zone configuration for outgoing zone transfers.

## Definition

```yaml
{"operationId": "secondary-dns-(-primary-zone)-create-primary-zone-configuration", "summary": "Create Primary Zone Configuration", "description": "Create primary zone configuration for outgoing zone transfers.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/secondary-dns_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/secondary-dns_single_request_outgoing"}}}}, "responses": {"200": {"description": "Create Primary Zone Configuration response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/secondary-dns_single_response_outgoing"}}}}, "4XX": {"description": "Create Primary Zone Configuration response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/secondary-dns_single_response_outgoing"}, {"$ref": "#/components/schemas/secondary-dns_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Secondary DNS (Primary Zone)"], "x-api-token-group": ["Zone Settings Write", "Zone Write", "DNS Write"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "dns.zone-transfers.outgoing", "x-fern-sdk-method-name": "create"}
```
