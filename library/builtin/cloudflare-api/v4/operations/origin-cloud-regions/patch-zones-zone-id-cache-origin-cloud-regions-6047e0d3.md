---
title: Create or update an origin cloud region mapping
page_id: operation-patch-zones-zone-id-cache-origin-cloud-regions-081284f5
path: operations/origin-cloud-regions
description: Adds or updates a single IP-to-cloud-region mapping for the zone. Unlike POST, this operation is idempotent — if a mapping for the IP already exists it is overwritten. Returns the complete updated list of all mappings for the zone. Returns 403 (code 1164) when the zone has reached the limit of 3,500 IP mappings.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /zones/{zone_id}/cache/origin_cloud_regions
operation_ids:
    - origin-cloud-regions-upsert
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create or update an origin cloud region mapping

`PATCH /zones/{zone_id}/cache/origin_cloud_regions`

Operation ID: `origin-cloud-regions-upsert`

Adds or updates a single IP-to-cloud-region mapping for the zone. Unlike POST, this operation is idempotent — if a mapping for the IP already exists it is overwritten. Returns the complete updated list of all mappings for the zone. Returns 403 (code 1164) when the zone has reached the limit of 3,500 IP mappings.

## Definition

```yaml
{"operationId": "origin-cloud-regions-upsert", "summary": "Create or update an origin cloud region mapping", "description": "Adds or updates a single IP-to-cloud-region mapping for the zone. Unlike POST, this operation is idempotent — if a mapping for the IP already exists it is overwritten. Returns the complete updated list of all mappings for the zone. Returns 403 (code 1164) when the zone has reached the limit of 3,500 IP mappings.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cache-rules_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"examples": {"gcp": {"$ref": "#/components/examples/cache-rules_origin_cloud_region_request_gcp"}}, "schema": {"$ref": "#/components/schemas/cache-rules_origin_cloud_region_request"}}}}, "responses": {"200": {"description": "Create or update origin cloud region mapping response.", "content": {"application/json": {"examples": {"patched": {"$ref": "#/components/examples/cache-rules_origin_cloud_regions_list"}}, "schema": {"allOf": [{"$ref": "#/components/schemas/cache-rules_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/cache-rules_origin_cloud_regions_list_result"}}, "type": "object"}]}}}}, "4XX": {"description": "Create or update origin cloud region mapping failure.", "content": {"application/json": {"examples": {"invalid_vendor_region": {"$ref": "#/components/examples/cache-rules_origin_cloud_region_invalid_vendor_region"}, "too_many_mappings": {"$ref": "#/components/examples/cache-rules_origin_cloud_region_too_many_mappings"}}, "schema": {"$ref": "#/components/schemas/cache-rules_api-response-common-failure"}}}}}, "deprecated": true, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Origin Cloud Regions"], "x-cfDeprecation": {"description": "Use the v2 endpoints under `/zones/{zone_id}/origin/cloud_regions` instead. The v2 API uses idempotent PUT semantics for create+update, provides a flat response shape, and is the supported path for Terraform and SDK clients. The v1 endpoints under `/zones/{zone_id}/cache/origin_cloud_regions` will be removed on 2026-10-01.", "display": true, "eol": "2026-10-01", "id": "opcr_v1_deprecation"}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
