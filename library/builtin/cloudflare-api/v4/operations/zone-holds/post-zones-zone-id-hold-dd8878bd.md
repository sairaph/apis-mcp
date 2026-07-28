---
title: Create Zone Hold
page_id: operation-post-zones-zone-id-hold-77c49deb
path: operations/zone-holds
description: |-
    Enforce a zone hold on the zone, blocking the creation and activation of zones with this zone's hostname.
    Zone holds cannot be enabled on CDN-only zones.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /zones/{zone_id}/hold
operation_ids:
    - zones-0-hold-post
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create Zone Hold

`POST /zones/{zone_id}/hold`

Operation ID: `zones-0-hold-post`

Enforce a zone hold on the zone, blocking the creation and activation of zones with this zone's hostname.
Zone holds cannot be enabled on CDN-only zones.

## Definition

```yaml
{"operationId": "zones-0-hold-post", "summary": "Create Zone Hold", "description": "Enforce a zone hold on the zone, blocking the creation and activation of zones with this zone's hostname.\nZone holds cannot be enabled on CDN-only zones.", "parameters": [{"name": "zone_id", "in": "path", "description": "Zone ID", "required": true, "schema": {"$ref": "#/components/schemas/zones_identifier-2"}}, {"name": "include_subdomains", "in": "query", "description": "If provided, the zone hold will extend to block any subdomain of the given zone, as well\nas SSL4SaaS Custom Hostnames. For example, a zone hold on a zone with the hostname\n'example.com' and include_subdomains=true will block 'example.com',\n'staging.example.com', 'api.staging.example.com', etc.", "schema": {"type": "boolean"}, "example": true}], "responses": {"200": {"description": "Successful Response", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/zones_api-response-single-2"}, {"properties": {"result": {"type": "object", "properties": {"hold": {"type": "boolean", "example": true}, "hold_after": {"type": "string", "example": "2023-01-31T15:56:36+00:00"}, "include_subdomains": {"type": "string", "example": true}}}}, "type": "object"}]}}}}, "4XX": {"description": "Client Error", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/zones_api-response-common-failure-2"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Zone Holds"], "x-api-token-group": ["Zone Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zones.holds", "x-fern-sdk-method-name": "create"}
```
