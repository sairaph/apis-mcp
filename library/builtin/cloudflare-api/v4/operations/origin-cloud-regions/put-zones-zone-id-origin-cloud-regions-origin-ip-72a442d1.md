---
title: Create or replace an origin cloud region mapping
page_id: operation-put-zones-zone-id-origin-cloud-regions-origin-ip-84571e2d
path: operations/origin-cloud-regions
description: Creates a new IP-to-cloud-region mapping or replaces the existing mapping for the specified IP. PUT is idempotent — calling it repeatedly with the same body produces the same result. The IP path parameter is normalized to canonical form (RFC 5952 for IPv6) before storage. The vendor and region are validated against the list from `GET /zones/{zone_id}/origin/cloud_regions/supported_regions`. Returns 400 if the `origin_ip` in the body does not match the URL path parameter. Returns 403 (code 1164) when the zone has reached the limit of 3,500 IP mappings.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /zones/{zone_id}/origin/cloud_regions/{origin_ip}
operation_ids:
    - origin-cloud-regions-v2-upsert
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create or replace an origin cloud region mapping

`PUT /zones/{zone_id}/origin/cloud_regions/{origin_ip}`

Operation ID: `origin-cloud-regions-v2-upsert`

Creates a new IP-to-cloud-region mapping or replaces the existing mapping for the specified IP. PUT is idempotent — calling it repeatedly with the same body produces the same result. The IP path parameter is normalized to canonical form (RFC 5952 for IPv6) before storage. The vendor and region are validated against the list from `GET /zones/{zone_id}/origin/cloud_regions/supported_regions`. Returns 400 if the `origin_ip` in the body does not match the URL path parameter. Returns 403 (code 1164) when the zone has reached the limit of 3,500 IP mappings.

## Definition

```yaml
{"operationId": "origin-cloud-regions-v2-upsert", "summary": "Create or replace an origin cloud region mapping", "description": "Creates a new IP-to-cloud-region mapping or replaces the existing mapping for the specified IP. PUT is idempotent — calling it repeatedly with the same body produces the same result. The IP path parameter is normalized to canonical form (RFC 5952 for IPv6) before storage. The vendor and region are validated against the list from `GET /zones/{zone_id}/origin/cloud_regions/supported_regions`. Returns 400 if the `origin_ip` in the body does not match the URL path parameter. Returns 403 (code 1164) when the zone has reached the limit of 3,500 IP mappings.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cache-rules_identifier"}}, {"name": "origin_ip", "in": "path", "description": "Origin IP address to create or replace.", "required": true, "schema": {"type": "string", "example": "192.0.2.1"}}], "requestBody": {"required": true, "content": {"application/json": {"examples": {"aws": {"$ref": "#/components/examples/cache-rules_origin_cloud_region_v2_request_aws"}}, "schema": {"$ref": "#/components/schemas/cache-rules_origin_cloud_region_v2_request"}}}}, "responses": {"200": {"description": "Create or replace origin cloud region mapping response.", "content": {"application/json": {"examples": {"upserted": {"$ref": "#/components/examples/cache-rules_origin_cloud_region_v2_single"}}, "schema": {"allOf": [{"$ref": "#/components/schemas/cache-rules_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/cache-rules_origin_cloud_region_v2_entry"}}, "type": "object"}]}}}}, "4XX": {"description": "Create or replace origin cloud region mapping failure.", "content": {"application/json": {"examples": {"invalid_vendor_region": {"$ref": "#/components/examples/cache-rules_origin_cloud_region_v2_invalid_vendor_region"}, "ip_mismatch": {"$ref": "#/components/examples/cache-rules_origin_cloud_region_v2_ip_mismatch"}, "too_many_mappings": {"$ref": "#/components/examples/cache-rules_origin_cloud_region_v2_too_many_mappings"}}, "schema": {"$ref": "#/components/schemas/cache-rules_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Origin Cloud Regions"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
