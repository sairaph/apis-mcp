---
title: Delete a Page Rule
page_id: operation-delete-zones-zone-id-pagerules-pagerule-id-56686e04
path: operations/page-rules
description: Deletes an existing Page Rule.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /zones/{zone_id}/pagerules/{pagerule_id}
operation_ids:
    - page-rules-delete-a-page-rule
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete a Page Rule

`DELETE /zones/{zone_id}/pagerules/{pagerule_id}`

Operation ID: `page-rules-delete-a-page-rule`

Deletes an existing Page Rule.

## Definition

```yaml
{"operationId": "page-rules-delete-a-page-rule", "summary": "Delete a Page Rule", "description": "Deletes an existing Page Rule.", "parameters": [{"name": "pagerule_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zones_identifier-2"}, "example": "023e105f4ecef8ad9ca31a8372d0c353"}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zones_identifier-2"}, "example": "023e105f4ecef8ad9ca31a8372d0c353"}], "responses": {"200": {"description": "Delete a Page Rule response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zones_api-response-single-id-2"}}}}, "4XX": {"description": "Delete a Page Rule response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zones_api-response-common-failure-2"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Page Rules"], "x-api-token-group": ["Zone Write", "Page Rules Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "page-rules", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
