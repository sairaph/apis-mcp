---
title: Create an origin cloud region mapping
page_id: operation-post-zones-zone-id-cache-origin-cloud-regions-85fd9512
path: operations/origin-cloud-regions
description: Adds a single IP-to-cloud-region mapping for the zone. The IP must be a valid IPv4 or IPv6 address and is normalized to canonical form before storage (RFC 5952 for IPv6). Returns 400 (code 1145) if a mapping for that IP already exists — use PATCH to update an existing entry. The vendor and region are validated against the list from `GET /zones/{zone_id}/cache/origin_cloud_regions/supported_regions`.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /zones/{zone_id}/cache/origin_cloud_regions
operation_ids:
    - origin-cloud-regions-create
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create an origin cloud region mapping

`POST /zones/{zone_id}/cache/origin_cloud_regions`

Operation ID: `origin-cloud-regions-create`

Adds a single IP-to-cloud-region mapping for the zone. The IP must be a valid IPv4 or IPv6 address and is normalized to canonical form before storage (RFC 5952 for IPv6). Returns 400 (code 1145) if a mapping for that IP already exists — use PATCH to update an existing entry. The vendor and region are validated against the list from `GET /zones/{zone_id}/cache/origin_cloud_regions/supported_regions`.

## Definition

```yaml
{"operationId": "origin-cloud-regions-create", "summary": "Create an origin cloud region mapping", "description": "Adds a single IP-to-cloud-region mapping for the zone. The IP must be a valid IPv4 or IPv6 address and is normalized to canonical form before storage (RFC 5952 for IPv6). Returns 400 (code 1145) if a mapping for that IP already exists — use PATCH to update an existing entry. The vendor and region are validated against the list from `GET /zones/{zone_id}/cache/origin_cloud_regions/supported_regions`.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cache-rules_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"examples": {"aws": {"$ref": "#/components/examples/cache-rules_origin_cloud_region_request_aws"}}, "schema": {"$ref": "#/components/schemas/cache-rules_origin_cloud_region_request"}}}}, "responses": {"200": {"description": "Create origin cloud region mapping response.", "content": {"application/json": {"examples": {"created": {"$ref": "#/components/examples/cache-rules_origin_cloud_region_single"}}, "schema": {"allOf": [{"$ref": "#/components/schemas/cache-rules_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/cache-rules_origin_cloud_region_single_result"}}, "type": "object"}]}}}}, "4XX": {"description": "Create origin cloud region mapping failure.", "content": {"application/json": {"examples": {"already_exists": {"$ref": "#/components/examples/cache-rules_origin_cloud_region_already_exists"}, "invalid_ip": {"$ref": "#/components/examples/cache-rules_origin_cloud_region_invalid_ip"}, "invalid_vendor_region": {"$ref": "#/components/examples/cache-rules_origin_cloud_region_invalid_vendor_region"}}, "schema": {"$ref": "#/components/schemas/cache-rules_api-response-common-failure"}}}}}, "deprecated": true, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Origin Cloud Regions"], "x-cfDeprecation": {"description": "Use the v2 endpoints under `/zones/{zone_id}/origin/cloud_regions` instead. The v2 API uses idempotent PUT semantics for create+update, provides a flat response shape, and is the supported path for Terraform and SDK clients. The v1 endpoints under `/zones/{zone_id}/cache/origin_cloud_regions` will be removed on 2026-10-01.", "display": true, "eol": "2026-10-01", "id": "opcr_v1_deprecation"}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
