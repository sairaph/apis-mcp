---
title: Update a Page Rule
page_id: operation-put-zones-zone-id-pagerules-pagerule-id-3b46a796
path: operations/page-rules
description: Replaces the configuration of an existing Page Rule. The configuration of the updated Page Rule will exactly match the data passed in the API request.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /zones/{zone_id}/pagerules/{pagerule_id}
operation_ids:
    - page-rules-update-a-page-rule
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update a Page Rule

`PUT /zones/{zone_id}/pagerules/{pagerule_id}`

Operation ID: `page-rules-update-a-page-rule`

Replaces the configuration of an existing Page Rule. The configuration of the updated Page Rule will exactly match the data passed in the API request.

## Definition

```yaml
{"operationId": "page-rules-update-a-page-rule", "summary": "Update a Page Rule", "description": "Replaces the configuration of an existing Page Rule. The configuration of the updated Page Rule will exactly match the data passed in the API request.", "parameters": [{"name": "pagerule_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zones_identifier-2"}, "example": "023e105f4ecef8ad9ca31a8372d0c353"}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zones_identifier-2"}, "example": "023e105f4ecef8ad9ca31a8372d0c353"}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"actions": {"$ref": "#/components/schemas/zones_actions"}, "priority": {"$ref": "#/components/schemas/zones_priority"}, "status": {"$ref": "#/components/schemas/zones_status"}, "targets": {"$ref": "#/components/schemas/zones_targets"}}, "required": ["targets", "actions"]}}}}, "responses": {"200": {"description": "Update a Page Rule response", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/zones_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/zones_page_rule"}}}]}}}}, "4XX": {"description": "Update a Page Rule response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zones_api-response-common-failure-2"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Page Rules"], "x-api-token-group": ["Zone Write", "Page Rules Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "page-rules", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
