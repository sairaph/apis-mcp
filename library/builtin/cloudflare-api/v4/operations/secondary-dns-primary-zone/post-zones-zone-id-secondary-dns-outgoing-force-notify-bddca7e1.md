---
title: Force DNS NOTIFY
page_id: operation-post-zones-zone-id-secondary-dns-outgoing-force-notify-01145937
path: operations/secondary-dns-primary-zone
description: Notifies the secondary nameserver(s) and clears IXFR backlog of primary zone.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /zones/{zone_id}/secondary_dns/outgoing/force_notify
operation_ids:
    - secondary-dns-(-primary-zone)-force-dns-notify
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Force DNS NOTIFY

`POST /zones/{zone_id}/secondary_dns/outgoing/force_notify`

Operation ID: `secondary-dns-(-primary-zone)-force-dns-notify`

Notifies the secondary nameserver(s) and clears IXFR backlog of primary zone.

## Definition

```yaml
{"operationId": "secondary-dns-(-primary-zone)-force-dns-notify", "summary": "Force DNS NOTIFY", "description": "Notifies the secondary nameserver(s) and clears IXFR backlog of primary zone.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/secondary-dns_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {}}}, "responses": {"200": {"description": "Force DNS NOTIFY response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/secondary-dns_force_response-2"}}}}, "4XX": {"description": "Force DNS NOTIFY response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/secondary-dns_force_response-2"}, {"$ref": "#/components/schemas/secondary-dns_api-response-common-failure"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Secondary DNS (Primary Zone)"], "x-api-token-group": ["Zone Settings Write", "Zone Write", "DNS Write"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "dns.zone-transfers.outgoing", "x-fern-sdk-method-name": "force-notify"}
```
