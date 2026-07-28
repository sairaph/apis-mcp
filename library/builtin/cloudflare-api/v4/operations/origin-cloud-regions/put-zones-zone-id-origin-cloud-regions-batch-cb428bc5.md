---
title: Batch create or replace origin cloud region mappings
page_id: operation-put-zones-zone-id-origin-cloud-regions-batch-8a07ef9c
path: operations/origin-cloud-regions
description: Upserts up to 100 IP-to-cloud-region mappings in a single request. Items in the request body are created or replaced; mappings not included in the request body are preserved unchanged (this is a merge operation, not a full collection replacement). Each item is validated independently — valid items are applied and invalid items are returned in the `failed` array. The vendor and region for every item are validated against the list from `GET /zones/{zone_id}/origin/cloud_regions/supported_regions`.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /zones/{zone_id}/origin/cloud_regions/batch
operation_ids:
    - origin-cloud-regions-v2-batch-upsert
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Batch create or replace origin cloud region mappings

`PUT /zones/{zone_id}/origin/cloud_regions/batch`

Operation ID: `origin-cloud-regions-v2-batch-upsert`

Upserts up to 100 IP-to-cloud-region mappings in a single request. Items in the request body are created or replaced; mappings not included in the request body are preserved unchanged (this is a merge operation, not a full collection replacement). Each item is validated independently — valid items are applied and invalid items are returned in the `failed` array. The vendor and region for every item are validated against the list from `GET /zones/{zone_id}/origin/cloud_regions/supported_regions`.

## Definition

```yaml
{"operationId": "origin-cloud-regions-v2-batch-upsert", "summary": "Batch create or replace origin cloud region mappings", "description": "Upserts up to 100 IP-to-cloud-region mappings in a single request. Items in the request body are created or replaced; mappings not included in the request body are preserved unchanged (this is a merge operation, not a full collection replacement). Each item is validated independently — valid items are applied and invalid items are returned in the `failed` array. The vendor and region for every item are validated against the list from `GET /zones/{zone_id}/origin/cloud_regions/supported_regions`.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cache-rules_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"examples": {"batch_put": {"$ref": "#/components/examples/cache-rules_origin_cloud_region_v2_batch_put_request"}}, "schema": {"type": "array", "items": {"$ref": "#/components/schemas/cache-rules_origin_cloud_region_v2_request"}, "maxItems": 100}}}}, "responses": {"200": {"description": "Batch create or replace origin cloud region mappings successful.", "content": {"application/json": {"examples": {"batch_upserted": {"$ref": "#/components/examples/cache-rules_origin_cloud_region_v2_batch_put_response"}}, "schema": {"allOf": [{"$ref": "#/components/schemas/cache-rules_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/cache-rules_origin_cloud_region_v2_batch_result"}}, "type": "object"}]}}}}, "4XX": {"description": "Batch create or replace origin cloud region mappings failure.", "content": {"application/json": {"examples": {"too_many_mappings": {"$ref": "#/components/examples/cache-rules_origin_cloud_region_v2_too_many_mappings"}}, "schema": {"$ref": "#/components/schemas/cache-rules_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Origin Cloud Regions"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
