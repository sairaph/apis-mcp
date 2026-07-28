---
title: Batch create or update origin cloud region mappings
page_id: operation-patch-zones-zone-id-cache-origin-cloud-regions-batch-d5604a54
path: operations/origin-cloud-regions
description: Adds or updates up to 100 IP-to-cloud-region mappings in a single request. Each item is validated independently — valid items are applied and invalid items are returned in the `failed` array. The vendor and region for every item are validated against the list from `GET /zones/{zone_id}/cache/origin_cloud_regions/supported_regions`.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /zones/{zone_id}/cache/origin_cloud_regions/batch
operation_ids:
    - origin-cloud-regions-batch-upsert
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Batch create or update origin cloud region mappings

`PATCH /zones/{zone_id}/cache/origin_cloud_regions/batch`

Operation ID: `origin-cloud-regions-batch-upsert`

Adds or updates up to 100 IP-to-cloud-region mappings in a single request. Each item is validated independently — valid items are applied and invalid items are returned in the `failed` array. The vendor and region for every item are validated against the list from `GET /zones/{zone_id}/cache/origin_cloud_regions/supported_regions`.

## Definition

```yaml
{"operationId": "origin-cloud-regions-batch-upsert", "summary": "Batch create or update origin cloud region mappings", "description": "Adds or updates up to 100 IP-to-cloud-region mappings in a single request. Each item is validated independently — valid items are applied and invalid items are returned in the `failed` array. The vendor and region for every item are validated against the list from `GET /zones/{zone_id}/cache/origin_cloud_regions/supported_regions`.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cache-rules_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"examples": {"batch_patch": {"$ref": "#/components/examples/cache-rules_origin_cloud_region_batch_patch_request"}}, "schema": {"type": "array", "items": {"$ref": "#/components/schemas/cache-rules_origin_cloud_region_request"}, "maxItems": 100}}}}, "responses": {"200": {"description": "Batch create or update origin cloud region mappings successful.", "content": {"application/json": {"examples": {"batch_patched": {"$ref": "#/components/examples/cache-rules_origin_cloud_region_batch_patch_response"}}, "schema": {"allOf": [{"$ref": "#/components/schemas/cache-rules_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/cache-rules_origin_cloud_region_batch_result"}}, "type": "object"}]}}}}, "4XX": {"description": "Batch create or update origin cloud region mappings failure.", "content": {"application/json": {"examples": {"too_many_mappings": {"$ref": "#/components/examples/cache-rules_origin_cloud_region_too_many_mappings"}}, "schema": {"$ref": "#/components/schemas/cache-rules_api-response-common-failure"}}}}, "5XX": {"description": "Batch create or update origin cloud region mappings internal error.", "content": {"application/json": {"examples": {"internal_error": {"$ref": "#/components/examples/cache-rules_origin_cloud_region_batch_patch_internal_error"}, "region_validation_error": {"$ref": "#/components/examples/cache-rules_origin_cloud_region_region_validation_error"}}, "schema": {"$ref": "#/components/schemas/cache-rules_api-response-common-failure"}}}}}, "deprecated": true, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Origin Cloud Regions"], "x-cfDeprecation": {"description": "Use the v2 endpoints under `/zones/{zone_id}/origin/cloud_regions` instead. The v2 API uses idempotent PUT semantics for create+update, provides a flat response shape, and is the supported path for Terraform and SDK clients. The v1 endpoints under `/zones/{zone_id}/cache/origin_cloud_regions` will be removed on 2026-10-01.", "display": true, "eol": "2026-10-01", "id": "opcr_v1_deprecation"}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
