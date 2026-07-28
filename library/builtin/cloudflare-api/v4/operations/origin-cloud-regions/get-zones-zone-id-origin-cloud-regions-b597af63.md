---
title: List origin cloud region mappings
page_id: operation-get-zones-zone-id-origin-cloud-regions-220f91a8
path: operations/origin-cloud-regions
description: Returns all IP-to-cloud-region mappings configured for the zone with pagination support. Each mapping tells Cloudflare which cloud vendor and region hosts the origin at that IP, enabling the edge to route via the nearest Tiered Cache upper-tier co-located with that cloud provider. Returns an empty array when no mappings exist.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/origin/cloud_regions
operation_ids:
    - origin-cloud-regions-v2-list
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List origin cloud region mappings

`GET /zones/{zone_id}/origin/cloud_regions`

Operation ID: `origin-cloud-regions-v2-list`

Returns all IP-to-cloud-region mappings configured for the zone with pagination support. Each mapping tells Cloudflare which cloud vendor and region hosts the origin at that IP, enabling the edge to route via the nearest Tiered Cache upper-tier co-located with that cloud provider. Returns an empty array when no mappings exist.

## Definition

```yaml
{"operationId": "origin-cloud-regions-v2-list", "summary": "List origin cloud region mappings", "description": "Returns all IP-to-cloud-region mappings configured for the zone with pagination support. Each mapping tells Cloudflare which cloud vendor and region hosts the origin at that IP, enabling the edge to route via the nearest Tiered Cache upper-tier co-located with that cloud provider. Returns an empty array when no mappings exist.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cache-rules_identifier"}}, {"name": "page", "in": "query", "description": "Page number of paginated results.", "schema": {"type": "integer", "default": 1, "minimum": 1}}, {"name": "per_page", "in": "query", "description": "Number of items per page.", "schema": {"type": "integer", "default": 20, "maximum": 100, "minimum": 1}}], "responses": {"200": {"description": "List origin cloud region mappings response.", "content": {"application/json": {"examples": {"empty": {"$ref": "#/components/examples/cache-rules_origin_cloud_region_v2_list_empty"}, "two_mappings": {"$ref": "#/components/examples/cache-rules_origin_cloud_region_v2_list"}}, "schema": {"allOf": [{"$ref": "#/components/schemas/cache-rules_api-response-common"}, {"properties": {"result": {"type": "array", "items": {"$ref": "#/components/schemas/cache-rules_origin_cloud_region_v2_entry"}}, "result_info": {"$ref": "#/components/schemas/cache-rules_origin_cloud_region_v2_result_info"}}, "required": ["result", "result_info"], "type": "object"}]}}}}, "4XX": {"description": "List origin cloud region mappings failure.", "content": {"application/json": {"examples": {"invalid_pagination": {"$ref": "#/components/examples/cache-rules_origin_cloud_region_v2_invalid_pagination"}}, "schema": {"$ref": "#/components/schemas/cache-rules_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Origin Cloud Regions"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
