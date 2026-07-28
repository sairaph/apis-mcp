---
title: Create a Page Rule
page_id: operation-post-zones-zone-id-pagerules-df51e8b0
path: operations/page-rules
description: Creates a new Page Rule.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /zones/{zone_id}/pagerules
operation_ids:
    - page-rules-create-a-page-rule
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a Page Rule

`POST /zones/{zone_id}/pagerules`

Operation ID: `page-rules-create-a-page-rule`

Creates a new Page Rule.

## Definition

```yaml
{"operationId": "page-rules-create-a-page-rule", "summary": "Create a Page Rule", "description": "Creates a new Page Rule.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zones_identifier-2"}, "example": "023e105f4ecef8ad9ca31a8372d0c353"}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"actions": {"$ref": "#/components/schemas/zones_actions"}, "priority": {"$ref": "#/components/schemas/zones_priority"}, "status": {"$ref": "#/components/schemas/zones_status"}, "targets": {"$ref": "#/components/schemas/zones_targets"}}, "required": ["targets", "actions"]}}}}, "responses": {"200": {"description": "Create a Page Rule response", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/zones_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/zones_page_rule"}}}]}}}}, "4XX": {"description": "Create a Page Rule response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zones_api-response-common-failure-2"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Page Rules"], "x-api-token-group": ["Zone Write", "Page Rules Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "page-rules", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
