---
title: Edit a Page Rule
page_id: operation-patch-zones-zone-id-pagerules-pagerule-id-6503ab71
path: operations/page-rules
description: Updates one or more fields of an existing Page Rule.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /zones/{zone_id}/pagerules/{pagerule_id}
operation_ids:
    - page-rules-edit-a-page-rule
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Edit a Page Rule

`PATCH /zones/{zone_id}/pagerules/{pagerule_id}`

Operation ID: `page-rules-edit-a-page-rule`

Updates one or more fields of an existing Page Rule.

## Definition

```yaml
{"operationId": "page-rules-edit-a-page-rule", "summary": "Edit a Page Rule", "description": "Updates one or more fields of an existing Page Rule.", "parameters": [{"name": "pagerule_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zones_identifier-2"}, "example": "023e105f4ecef8ad9ca31a8372d0c353"}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zones_identifier-2"}, "example": "023e105f4ecef8ad9ca31a8372d0c353"}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"actions": {"$ref": "#/components/schemas/zones_actions"}, "priority": {"$ref": "#/components/schemas/zones_priority"}, "status": {"$ref": "#/components/schemas/zones_status"}, "targets": {"$ref": "#/components/schemas/zones_targets"}}}}}}, "responses": {"200": {"description": "Edit a Page Rule response", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/zones_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/zones_page_rule"}}}]}}}}, "4XX": {"description": "Edit a Page Rule response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zones_api-response-common-failure-2"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Page Rules"], "x-api-token-group": ["Zone Write", "Page Rules Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "page-rules", "x-fern-sdk-method-name": "edit", "x-forge-hidden": true}
```
