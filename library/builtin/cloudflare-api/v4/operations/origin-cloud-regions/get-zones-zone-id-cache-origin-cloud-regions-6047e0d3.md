---
title: List origin cloud region mappings
page_id: operation-get-zones-zone-id-cache-origin-cloud-regions-d8f966ac
path: operations/origin-cloud-regions
description: Returns all IP-to-cloud-region mappings configured for the zone. Each mapping tells Cloudflare which cloud vendor and region hosts the origin at that IP, enabling the edge to route via the nearest Tiered Cache upper-tier co-located with that cloud provider. Returns an empty array when no mappings exist.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/cache/origin_cloud_regions
operation_ids:
    - origin-cloud-regions-list
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List origin cloud region mappings

`GET /zones/{zone_id}/cache/origin_cloud_regions`

Operation ID: `origin-cloud-regions-list`

Returns all IP-to-cloud-region mappings configured for the zone. Each mapping tells Cloudflare which cloud vendor and region hosts the origin at that IP, enabling the edge to route via the nearest Tiered Cache upper-tier co-located with that cloud provider. Returns an empty array when no mappings exist.

## Definition

```yaml
{"operationId": "origin-cloud-regions-list", "summary": "List origin cloud region mappings", "description": "Returns all IP-to-cloud-region mappings configured for the zone. Each mapping tells Cloudflare which cloud vendor and region hosts the origin at that IP, enabling the edge to route via the nearest Tiered Cache upper-tier co-located with that cloud provider. Returns an empty array when no mappings exist.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cache-rules_identifier"}}], "responses": {"200": {"description": "List origin cloud region mappings response.", "content": {"application/json": {"examples": {"empty": {"$ref": "#/components/examples/cache-rules_origin_cloud_regions_empty"}, "two_mappings": {"$ref": "#/components/examples/cache-rules_origin_cloud_regions_list"}}, "schema": {"allOf": [{"$ref": "#/components/schemas/cache-rules_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/cache-rules_origin_cloud_regions_list_result"}}, "type": "object"}]}}}}, "4XX": {"description": "List origin cloud region mappings failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cache-rules_api-response-common-failure"}}}}}, "deprecated": true, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Origin Cloud Regions"], "x-cfDeprecation": {"description": "Use the v2 endpoints under `/zones/{zone_id}/origin/cloud_regions` instead. The v2 API uses idempotent PUT semantics for create+update, provides a flat response shape, and is the supported path for Terraform and SDK clients. The v1 endpoints under `/zones/{zone_id}/cache/origin_cloud_regions` will be removed on 2026-10-01.", "display": true, "eol": "2026-10-01", "id": "opcr_v1_deprecation"}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
