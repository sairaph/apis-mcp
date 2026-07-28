---
title: List supported cloud vendors and regions
page_id: operation-get-zones-zone-id-origin-cloud-regions-supported-regions-c566a8d3
path: operations/origin-cloud-regions
description: Returns the cloud vendors and regions that are valid values for origin cloud region mappings. Each region includes the Tiered Cache upper-tier colocation codes that will be used for cache routing when a mapping targeting that region is active. Requires the zone to have Tiered Cache enabled.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/origin/cloud_regions/supported_regions
operation_ids:
    - origin-cloud-regions-v2-supported-regions
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List supported cloud vendors and regions

`GET /zones/{zone_id}/origin/cloud_regions/supported_regions`

Operation ID: `origin-cloud-regions-v2-supported-regions`

Returns the cloud vendors and regions that are valid values for origin cloud region mappings. Each region includes the Tiered Cache upper-tier colocation codes that will be used for cache routing when a mapping targeting that region is active. Requires the zone to have Tiered Cache enabled.

## Definition

```yaml
{"operationId": "origin-cloud-regions-v2-supported-regions", "summary": "List supported cloud vendors and regions", "description": "Returns the cloud vendors and regions that are valid values for origin cloud region mappings. Each region includes the Tiered Cache upper-tier colocation codes that will be used for cache routing when a mapping targeting that region is active. Requires the zone to have Tiered Cache enabled.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cache-rules_identifier"}}], "responses": {"200": {"description": "List supported cloud vendors and regions response.", "content": {"application/json": {"examples": {"vendors": {"$ref": "#/components/examples/cache-rules_supported_cloud_regions"}}, "schema": {"allOf": [{"$ref": "#/components/schemas/cache-rules_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/cache-rules_supported_cloud_regions_result"}}, "type": "object"}]}}}}, "4XX": {"description": "List supported cloud vendors and regions failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cache-rules_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Origin Cloud Regions"], "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
