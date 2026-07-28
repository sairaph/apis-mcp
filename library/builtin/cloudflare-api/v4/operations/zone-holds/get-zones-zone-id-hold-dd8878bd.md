---
title: Get Zone Hold
page_id: operation-get-zones-zone-id-hold-9ad5b5d5
path: operations/zone-holds
description: Retrieve whether the zone is subject to a zone hold, and metadata about the hold.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/hold
operation_ids:
    - zones-0-hold-get
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Zone Hold

`GET /zones/{zone_id}/hold`

Operation ID: `zones-0-hold-get`

Retrieve whether the zone is subject to a zone hold, and metadata about the hold.

## Definition

```yaml
{"operationId": "zones-0-hold-get", "summary": "Get Zone Hold", "description": "Retrieve whether the zone is subject to a zone hold, and metadata about the hold.", "parameters": [{"name": "zone_id", "in": "path", "description": "Zone ID", "required": true, "schema": {"$ref": "#/components/schemas/zones_identifier-2"}}], "responses": {"200": {"description": "Successful Response", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/zones_api-response-single-2"}, {"properties": {"result": {"type": "object", "properties": {"hold": {"type": "boolean", "example": true}, "hold_after": {"type": "string", "example": "2023-01-31T15:56:36+00:00"}, "include_subdomains": {"type": "string", "example": false}}}}, "type": "object"}]}}}}, "4XX": {"description": "Client Error", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/zones_api-response-common-failure-2"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Zone Holds"], "x-api-token-group": ["Trust and Safety Write", "Trust and Safety Read", "Zero Trust: PII Read", "Zaraz Edit", "Zaraz Read", "Zaraz Admin", "Access: Apps and Policies Revoke", "Access: Apps and Policies Write", "Access: Apps and Policies Read", "Access: Apps and Policies Revoke", "Access: Mutual TLS Certificates Write", "Access: Organizations, Identity Providers, and Groups Write", "Zone Settings Write", "Zone Settings Read", "Zone Read", "DNS Read", "Workers Scripts Write", "Workers Scripts Read", "Zone Write", "Workers Routes Write", "Workers Routes Read", "Stream Write", "Stream Read", "SSL and Certificates Write", "SSL and Certificates Read", "Logs Write", "Logs Read", "Cache Purge", "Page Rules Write", "Page Rules Read", "Load Balancers Write", "Load Balancers Read", "Firewall Services Write", "Firewall Services Read", "DNS Write", "Apps Write", "Analytics Read", "Access: Apps and Policies Write", "Access: Apps and Policies Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zones.holds", "x-fern-sdk-method-name": "get"}
```
