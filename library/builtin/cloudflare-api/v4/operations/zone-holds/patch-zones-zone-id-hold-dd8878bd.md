---
title: Update Zone Hold
page_id: operation-patch-zones-zone-id-hold-4170b69d
path: operations/zone-holds
description: |-
    Update the `hold_after` and/or `include_subdomains` values on an existing zone hold.
    The hold is enabled if the `hold_after` date-time value is in the past.
    Existing zone holds can be removed from CDN-only zones by setting `hold_after` to `null`.
    Other zone hold updates cannot be made on CDN-only zones.
    Active holds are automatically disabled when a zone transitions to CDN-only mode.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /zones/{zone_id}/hold
operation_ids:
    - zones-0-hold-patch
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update Zone Hold

`PATCH /zones/{zone_id}/hold`

Operation ID: `zones-0-hold-patch`

Update the `hold_after` and/or `include_subdomains` values on an existing zone hold.
The hold is enabled if the `hold_after` date-time value is in the past.
Existing zone holds can be removed from CDN-only zones by setting `hold_after` to `null`.
Other zone hold updates cannot be made on CDN-only zones.
Active holds are automatically disabled when a zone transitions to CDN-only mode.

## Definition

```yaml
{"operationId": "zones-0-hold-patch", "summary": "Update Zone Hold", "description": "Update the `hold_after` and/or `include_subdomains` values on an existing zone hold.\nThe hold is enabled if the `hold_after` date-time value is in the past.\nExisting zone holds can be removed from CDN-only zones by setting `hold_after` to `null`.\nOther zone hold updates cannot be made on CDN-only zones.\nActive holds are automatically disabled when a zone transitions to CDN-only mode.", "parameters": [{"name": "zone_id", "in": "path", "description": "Zone ID", "required": true, "schema": {"$ref": "#/components/schemas/zones_identifier-2"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"hold_after": {"description": "If `hold_after` is provided and future-dated, the hold will be temporarily disabled,\nthen automatically re-enabled by the system at the time specified\nin this RFC3339-formatted timestamp. A past-dated `hold_after` value will have\nno effect on an existing, enabled hold. Providing an empty string will set its value\nto the current time. Providing `null` will disable the hold indefinitely.", "type": "string", "example": "2023-01-31T15:56:36+00:00", "default": "", "nullable": true}, "include_subdomains": {"description": "If `true`, the zone hold will extend to block any subdomain of the given zone, as well\nas SSL4SaaS Custom Hostnames. For example, a zone hold on a zone with the hostname\n'example.com' and include_subdomains=true will block 'example.com',\n'staging.example.com', 'api.staging.example.com', etc.", "type": "boolean", "example": true, "default": false}}}}}}, "responses": {"200": {"description": "Successful Response", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/zones_api-response-single-2"}, {"properties": {"result": {"type": "object", "properties": {"hold": {"type": "boolean", "example": false}, "hold_after": {"type": "string"}, "include_subdomains": {"type": "string", "example": false}}}}, "type": "object"}]}}}}, "4XX": {"description": "Client Error", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/zones_api-response-common-failure-2"}]}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Zone Holds"], "x-api-token-group": ["Zone Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zones.holds", "x-fern-sdk-method-name": "edit", "x-forge-params": {"hold_after": {"default": null}, "include_subdomains": {"default": null}}}
```
