---
title: Get an origin cloud region mapping
page_id: operation-get-zones-zone-id-origin-cloud-regions-origin-ip-26c22dce
path: operations/origin-cloud-regions
description: Returns the cloud region mapping for a single origin IP address. The IP path parameter is normalized before lookup (RFC 5952 for IPv6). Returns 404 if the zone has no mappings or if the specified IP has no mapping.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/origin/cloud_regions/{origin_ip}
operation_ids:
    - origin-cloud-regions-v2-get
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get an origin cloud region mapping

`GET /zones/{zone_id}/origin/cloud_regions/{origin_ip}`

Operation ID: `origin-cloud-regions-v2-get`

Returns the cloud region mapping for a single origin IP address. The IP path parameter is normalized before lookup (RFC 5952 for IPv6). Returns 404 if the zone has no mappings or if the specified IP has no mapping.

## Definition

```yaml
{"operationId": "origin-cloud-regions-v2-get", "summary": "Get an origin cloud region mapping", "description": "Returns the cloud region mapping for a single origin IP address. The IP path parameter is normalized before lookup (RFC 5952 for IPv6). Returns 404 if the zone has no mappings or if the specified IP has no mapping.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cache-rules_identifier"}}, {"name": "origin_ip", "in": "path", "description": "Origin IP address to look up. IPv4 and IPv6 are supported.", "required": true, "schema": {"type": "string", "example": "192.0.2.1"}}], "responses": {"200": {"description": "Get origin cloud region mapping response.", "content": {"application/json": {"examples": {"found": {"$ref": "#/components/examples/cache-rules_origin_cloud_region_v2_single"}}, "schema": {"allOf": [{"$ref": "#/components/schemas/cache-rules_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/cache-rules_origin_cloud_region_v2_entry"}}, "type": "object"}]}}}}, "4XX": {"description": "Get origin cloud region mapping failure.", "content": {"application/json": {"examples": {"not_found": {"$ref": "#/components/examples/cache-rules_origin_cloud_region_v2_not_found"}}, "schema": {"$ref": "#/components/schemas/cache-rules_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Origin Cloud Regions"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
