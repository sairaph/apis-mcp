---
title: Set tags for a zone-level resource
page_id: operation-put-zones-zone-id-tags-a1220428
path: operations/resource-tagging
description: Creates or updates tags for a specific zone-level resource. Replaces all existing tags for the resource.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /zones/{zone_id}/tags
operation_ids:
    - tags-zone-set
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Set tags for a zone-level resource

`PUT /zones/{zone_id}/tags`

Operation ID: `tags-zone-set`

Creates or updates tags for a specific zone-level resource. Replaces all existing tags for the resource.

## Definition

```yaml
{"operationId": "tags-zone-set", "summary": "Set tags for a zone-level resource", "description": "Creates or updates tags for a specific zone-level resource. Replaces all existing tags for the resource.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/resource-tagging_zone_id"}}, {"$ref": "#/components/parameters/resource-tagging_if_match"}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/resource-tagging_set_tags_request_zone_level"}}}}, "responses": {"200": {"description": "Set tags response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/resource-tagging_tagged_resource_response_single"}}}}, "412": {"description": "Precondition failed. The resource has been modified since the provided ETag was obtained.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/resource-tagging_api-response-common-failure"}}}}, "4XX": {"description": "Set tags response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/resource-tagging_api-response-common-failure"}}}}, "5XX": {"description": "Set tags response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/resource-tagging_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Resource Tagging"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "beta", "x-fern-sdk-group-name": "zone-tags", "x-fern-sdk-method-name": "update", "x-stability": "beta"}
```
