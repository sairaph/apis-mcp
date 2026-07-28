---
title: Zone Details
page_id: operation-get-zones-zone-id-1d67a29f
path: operations/zone
description: |-
    Retrieves detailed information about a specific zone identified by its zone ID.

    Returns zone configuration, status, nameservers, and associated metadata.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}
operation_ids:
    - zones-0-get
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Zone Details

`GET /zones/{zone_id}`

Operation ID: `zones-0-get`

Retrieves detailed information about a specific zone identified by its zone ID.

Returns zone configuration, status, nameservers, and associated metadata.

## Definition

```yaml
{"operationId": "zones-0-get", "summary": "Zone Details", "description": "Retrieves detailed information about a specific zone identified by its zone ID.\n\nReturns zone configuration, status, nameservers, and associated metadata.\n", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zones_identifier"}}], "responses": {"200": {"description": "Zone Details response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/zones_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/zones_zone"}}, "type": "object"}]}}}}, "4XX": {"description": "Zone Details response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zones_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Zone"], "x-api-token-group": ["Trust and Safety Write", "Trust and Safety Read", "Zero Trust: PII Read", "Zaraz Edit", "Zaraz Read", "Zaraz Admin", "Access: Apps and Policies Revoke", "Access: Apps and Policies Write", "Access: Apps and Policies Read", "Access: Apps and Policies Revoke", "Access: Mutual TLS Certificates Write", "Access: Organizations, Identity Providers, and Groups Write", "Zone Settings Write", "Zone Settings Read", "Zone Read", "DNS Read", "Workers Scripts Write", "Workers Scripts Read", "Zone Write", "Workers Routes Write", "Workers Routes Read", "Stream Write", "Stream Read", "SSL and Certificates Write", "SSL and Certificates Read", "Logs Write", "Logs Read", "Cache Purge", "Page Rules Write", "Page Rules Read", "Load Balancers Write", "Load Balancers Read", "Firewall Services Write", "Firewall Services Read", "DNS Write", "Apps Write", "Analytics Read", "Access: Apps and Policies Write", "Access: Apps and Policies Read"], "x-cfPermissionsRequired": {"enum": ["#zone:read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zones", "x-fern-sdk-method-name": "get"}
```
