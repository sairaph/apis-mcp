---
title: Remove Zone Hold
page_id: operation-delete-zones-zone-id-hold-7e346b78
path: operations/zone-holds
description: |-
    Stop enforcement of a zone hold on the zone, permanently or temporarily, allowing the
    creation and activation of zones with this zone's hostname.
    Existing zone holds can be removed from CDN-only zones when `hold_after` is not provided.
    Active holds are automatically disabled when a zone transitions to CDN-only mode.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /zones/{zone_id}/hold
operation_ids:
    - zones-0-hold-delete
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Remove Zone Hold

`DELETE /zones/{zone_id}/hold`

Operation ID: `zones-0-hold-delete`

Stop enforcement of a zone hold on the zone, permanently or temporarily, allowing the
creation and activation of zones with this zone's hostname.
Existing zone holds can be removed from CDN-only zones when `hold_after` is not provided.
Active holds are automatically disabled when a zone transitions to CDN-only mode.

## Definition

```yaml
{"operationId": "zones-0-hold-delete", "summary": "Remove Zone Hold", "description": "Stop enforcement of a zone hold on the zone, permanently or temporarily, allowing the\ncreation and activation of zones with this zone's hostname.\nExisting zone holds can be removed from CDN-only zones when `hold_after` is not provided.\nActive holds are automatically disabled when a zone transitions to CDN-only mode.", "parameters": [{"name": "zone_id", "in": "path", "description": "Zone ID", "required": true, "schema": {"$ref": "#/components/schemas/zones_identifier-2"}}, {"name": "hold_after", "in": "query", "description": "If `hold_after` is provided, the hold will be temporarily disabled,\nthen automatically re-enabled by the system at the time specified\nin this RFC3339-formatted timestamp. Otherwise, the hold will be\ndisabled indefinitely. `hold_after` cannot be provided for CDN-only zones.", "schema": {"type": "string"}, "example": "2023-01-31T15:56:36+00:00"}], "responses": {"200": {"description": "Successful Response", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/zones_api-response-single-2"}, {"properties": {"result": {"type": "object", "properties": {"hold": {"type": "boolean", "example": false}, "hold_after": {"type": "string"}, "include_subdomains": {"type": "string", "example": false}}}}, "type": "object"}]}}}}, "4XX": {"description": "Client Error", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/zones_api-response-common-failure-2"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Zone Holds"], "x-api-token-group": ["Zone Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zones.holds", "x-fern-sdk-method-name": "delete"}
```
