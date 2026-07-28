---
title: Get Zone Hold by Zone Name
page_id: operation-get-zones-zone-id-hold-zone-name-c5a7475a
path: operations/zone-holds
description: |-
    Retrieve whether a given hostname is subject to a zone hold, and metadata about the hold.
    This endpoint checks whether the given hostname (or any of its ancestor domains) is blocked
    by an active zone hold. If a hold with `include_subdomains` is active on an ancestor domain,
    that hold is returned. This endpoint is used internally by SSL/COMS to check hold status
    during zone activation.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/hold/{zone_name}
operation_ids:
    - zones-0-hold-zone-name-get
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Zone Hold by Zone Name

`GET /zones/{zone_id}/hold/{zone_name}`

Operation ID: `zones-0-hold-zone-name-get`

Retrieve whether a given hostname is subject to a zone hold, and metadata about the hold.
This endpoint checks whether the given hostname (or any of its ancestor domains) is blocked
by an active zone hold. If a hold with `include_subdomains` is active on an ancestor domain,
that hold is returned. This endpoint is used internally by SSL/COMS to check hold status
during zone activation.

## Definition

```yaml
{"operationId": "zones-0-hold-zone-name-get", "summary": "Get Zone Hold by Zone Name", "description": "Retrieve whether a given hostname is subject to a zone hold, and metadata about the hold.\nThis endpoint checks whether the given hostname (or any of its ancestor domains) is blocked\nby an active zone hold. If a hold with `include_subdomains` is active on an ancestor domain,\nthat hold is returned. This endpoint is used internally by SSL/COMS to check hold status\nduring zone activation.", "parameters": [{"name": "zone_id", "in": "path", "description": "Zone identifier. Consumed by the API gateway for routing; the backend\nhandler does not use this value directly.", "required": true, "schema": {"$ref": "#/components/schemas/zones_identifier-2"}}, {"name": "zone_name", "in": "path", "description": "The hostname to check for a zone hold. May be a subdomain (e.g. `subdomain.example.com`)\nor an apex domain (e.g. `example.com`). The service checks the hostname and its ancestor\ndomains for active holds with `include_subdomains` enabled.", "required": true, "schema": {"type": "string", "example": "subdomain.example.com", "maxLength": 253, "pattern": "^([a-zA-Z0-9]([a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?\\.)+[a-zA-Z]{2,}$"}}], "responses": {"200": {"description": "Successful Response", "content": {"application/json": {"schema": {"type": "object", "properties": {"hold": {"description": "Whether the hostname is currently subject to a zone hold.", "type": "boolean", "example": true}, "hold_after": {"description": "The RFC3339-formatted timestamp at which the hold will be automatically\nre-enabled, if the hold was temporarily disabled. Null if the hold is\npermanently enabled or not set.", "type": "string", "format": "date-time", "example": "2023-01-31T15:56:36+00:00", "nullable": true}, "include_subdomains": {"description": "Whether the hold extends to block subdomains of the held zone.", "type": "boolean", "example": false}}}}}}, "4XX": {"description": "Client Error", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/zones_api-response-common-failure-2"}]}}}}}, "security": [], "tags": ["Zone Holds"], "x-api-token-group": ["Trust and Safety Write", "Trust and Safety Read", "Zero Trust: PII Read", "Zaraz Edit", "Zaraz Read", "Zaraz Admin", "Access: Apps and Policies Revoke", "Access: Apps and Policies Write", "Access: Apps and Policies Read", "Access: Apps and Policies Revoke", "Access: Mutual TLS Certificates Write", "Access: Organizations, Identity Providers, and Groups Write", "Zone Settings Write", "Zone Settings Read", "Zone Read", "DNS Read", "Workers Scripts Write", "Workers Scripts Read", "Zone Write", "Workers Routes Write", "Workers Routes Read", "Stream Write", "Stream Read", "SSL and Certificates Write", "SSL and Certificates Read", "Logs Write", "Logs Read", "Cache Purge", "Page Rules Write", "Page Rules Read", "Load Balancers Write", "Load Balancers Read", "Firewall Services Write", "Firewall Services Read", "DNS Write", "Apps Write", "Analytics Read", "Access: Apps and Policies Write", "Access: Apps and Policies Read"]}
```
