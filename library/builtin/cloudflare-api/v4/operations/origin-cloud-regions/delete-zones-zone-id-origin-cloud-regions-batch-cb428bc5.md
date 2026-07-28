---
title: Batch delete origin cloud region mappings
page_id: operation-delete-zones-zone-id-origin-cloud-regions-batch-65b5bc43
path: operations/origin-cloud-regions
description: Removes up to 100 IP-to-cloud-region mappings in a single request. Each IP is validated independently — successfully deleted items are returned in the `succeeded` array and IPs that could not be found or are invalid are returned in the `failed` array.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /zones/{zone_id}/origin/cloud_regions/batch
operation_ids:
    - origin-cloud-regions-v2-batch-delete
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Batch delete origin cloud region mappings

`DELETE /zones/{zone_id}/origin/cloud_regions/batch`

Operation ID: `origin-cloud-regions-v2-batch-delete`

Removes up to 100 IP-to-cloud-region mappings in a single request. Each IP is validated independently — successfully deleted items are returned in the `succeeded` array and IPs that could not be found or are invalid are returned in the `failed` array.

## Definition

```yaml
{"operationId": "origin-cloud-regions-v2-batch-delete", "summary": "Batch delete origin cloud region mappings", "description": "Removes up to 100 IP-to-cloud-region mappings in a single request. Each IP is validated independently — successfully deleted items are returned in the `succeeded` array and IPs that could not be found or are invalid are returned in the `failed` array.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cache-rules_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"examples": {"batch_delete": {"$ref": "#/components/examples/cache-rules_origin_cloud_region_v2_batch_delete_request"}}, "schema": {"type": "array", "items": {"description": "Origin IP address to delete.", "example": "192.0.2.1", "type": "string"}, "maxItems": 100}}}}, "responses": {"200": {"description": "Batch delete origin cloud region mappings successful.", "content": {"application/json": {"examples": {"batch_deleted": {"$ref": "#/components/examples/cache-rules_origin_cloud_region_v2_batch_delete_response"}}, "schema": {"allOf": [{"$ref": "#/components/schemas/cache-rules_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/cache-rules_origin_cloud_region_v2_batch_result"}}, "type": "object"}]}}}}, "4XX": {"description": "Batch delete origin cloud region mappings failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cache-rules_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Origin Cloud Regions"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
