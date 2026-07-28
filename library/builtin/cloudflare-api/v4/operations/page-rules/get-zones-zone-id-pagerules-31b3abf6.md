---
title: List Page Rules
page_id: operation-get-zones-zone-id-pagerules-ab66672f
path: operations/page-rules
description: Fetches Page Rules in a zone.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/pagerules
operation_ids:
    - page-rules-list-page-rules
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Page Rules

`GET /zones/{zone_id}/pagerules`

Operation ID: `page-rules-list-page-rules`

Fetches Page Rules in a zone.

## Definition

```yaml
{"operationId": "page-rules-list-page-rules", "summary": "List Page Rules", "description": "Fetches Page Rules in a zone.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zones_identifier-2"}, "example": "023e105f4ecef8ad9ca31a8372d0c353"}, {"name": "order", "in": "query", "schema": {"description": "The field used to sort returned Page Rules.", "type": "string", "example": "status", "default": "priority", "enum": ["status", "priority"]}}, {"name": "direction", "in": "query", "schema": {"description": "The direction used to sort returned Page Rules.", "type": "string", "example": "desc", "default": "desc", "enum": ["asc", "desc"]}}, {"name": "match", "in": "query", "schema": {"description": "When set to `all`, all the search requirements must match. When set to `any`, only one of the search requirements has to match.", "type": "string", "default": "all", "enum": ["any", "all"]}}, {"name": "status", "in": "query", "schema": {"description": "The status of the Page Rule.", "type": "string", "example": "active", "default": "disabled", "enum": ["active", "disabled"]}}], "responses": {"200": {"description": "List Page Rules response", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/zones_api-response-common-2"}, {"properties": {"result": {"type": "array", "items": {"$ref": "#/components/schemas/zones_page_rule"}}}}]}}}}, "4XX": {"description": "List Page Rules response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zones_api-response-common-failure-2"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Page Rules"], "x-api-token-group": ["Zone Read", "Zone Write", "Page Rules Write", "Page Rules Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "page-rules", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
