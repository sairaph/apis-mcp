---
title: Delete an origin cloud region mapping
page_id: operation-delete-zones-zone-id-origin-cloud-regions-origin-ip-b053c7e9
path: operations/origin-cloud-regions
description: Removes the cloud region mapping for a single origin IP address. The IP path parameter is normalized before lookup. Returns the deleted IP on success. Returns 404 if no mapping exists for the specified IP. When the last mapping for the zone is removed the underlying rule record is also deleted.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /zones/{zone_id}/origin/cloud_regions/{origin_ip}
operation_ids:
    - origin-cloud-regions-v2-delete
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete an origin cloud region mapping

`DELETE /zones/{zone_id}/origin/cloud_regions/{origin_ip}`

Operation ID: `origin-cloud-regions-v2-delete`

Removes the cloud region mapping for a single origin IP address. The IP path parameter is normalized before lookup. Returns the deleted IP on success. Returns 404 if no mapping exists for the specified IP. When the last mapping for the zone is removed the underlying rule record is also deleted.

## Definition

```yaml
{"operationId": "origin-cloud-regions-v2-delete", "summary": "Delete an origin cloud region mapping", "description": "Removes the cloud region mapping for a single origin IP address. The IP path parameter is normalized before lookup. Returns the deleted IP on success. Returns 404 if no mapping exists for the specified IP. When the last mapping for the zone is removed the underlying rule record is also deleted.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cache-rules_identifier"}}, {"name": "origin_ip", "in": "path", "description": "Origin IP address whose mapping should be deleted.", "required": true, "schema": {"type": "string", "example": "192.0.2.1"}}], "responses": {"200": {"description": "Delete origin cloud region mapping response.", "content": {"application/json": {"examples": {"deleted": {"$ref": "#/components/examples/cache-rules_origin_cloud_region_v2_delete"}}, "schema": {"allOf": [{"$ref": "#/components/schemas/cache-rules_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/cache-rules_origin_cloud_region_v2_delete_result"}}, "type": "object"}]}}}}, "4XX": {"description": "Delete origin cloud region mapping failure.", "content": {"application/json": {"examples": {"not_found": {"$ref": "#/components/examples/cache-rules_origin_cloud_region_v2_not_found"}}, "schema": {"$ref": "#/components/schemas/cache-rules_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Origin Cloud Regions"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
