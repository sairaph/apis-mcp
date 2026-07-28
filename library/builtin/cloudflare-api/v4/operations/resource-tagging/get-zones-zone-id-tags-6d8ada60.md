---
title: Get tags for a zone-level resource
page_id: operation-get-zones-zone-id-tags-5333bc86
path: operations/resource-tagging
description: Retrieves tags for a specific zone-level resource.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/tags
operation_ids:
    - tags-zone-get
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get tags for a zone-level resource

`GET /zones/{zone_id}/tags`

Operation ID: `tags-zone-get`

Retrieves tags for a specific zone-level resource.

## Definition

```yaml
{"operationId": "tags-zone-get", "summary": "Get tags for a zone-level resource", "description": "Retrieves tags for a specific zone-level resource.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/resource-tagging_zone_id"}}, {"name": "resource_id", "in": "query", "description": "The ID of the resource to retrieve tags for.", "required": true, "schema": {"$ref": "#/components/schemas/resource-tagging_resource_id"}}, {"name": "resource_type", "in": "query", "description": "The type of the resource.", "required": true, "schema": {"$ref": "#/components/schemas/resource-tagging_zone_resource_type"}}, {"name": "access_application_id", "in": "query", "description": "Access application ID identifier. Required for access_application_policy resources.", "schema": {"$ref": "#/components/schemas/resource-tagging_access_application_id"}}], "responses": {"200": {"description": "Get tags for single resource response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/resource-tagging_tagged_resource_response_single"}}}}, "4XX": {"description": "Get tags for single resource response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/resource-tagging_api-response-common-failure"}}}}, "5XX": {"description": "Get tags for single resource response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/resource-tagging_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Resource Tagging"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}, "x-fern-availability": "beta", "x-fern-sdk-group-name": "zone-tags", "x-fern-sdk-method-name": "get", "x-stability": "beta"}
```
