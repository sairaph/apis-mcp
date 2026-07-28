---
title: Update an Access group
page_id: operation-put-zones-zone-id-access-groups-group-id-01461ff3
path: operations/zone-level-access-groups
description: Updates a configured Access group.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /zones/{zone_id}/access/groups/{group_id}
operation_ids:
    - zone-level-access-groups-update-an-access-group
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update an Access group

`PUT /zones/{zone_id}/access/groups/{group_id}`

Operation ID: `zone-level-access-groups-update-an-access-group`

Updates a configured Access group.

## Definition

```yaml
{"operationId": "zone-level-access-groups-update-an-access-group", "summary": "Update an Access group", "description": "Updates a configured Access group.", "parameters": [{"name": "group_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_uuid"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"properties": {"exclude": {"$ref": "#/components/schemas/access_exclude"}, "include": {"$ref": "#/components/schemas/access_include"}, "name": {"$ref": "#/components/schemas/access_name-16"}, "require": {"$ref": "#/components/schemas/access_require"}}, "required": ["name", "include"]}}}}, "responses": {"200": {"description": "Update an Access group response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_single_response-17"}}}}, "4XX": {"description": "Update an Access group response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Zone-Level Access groups"], "x-api-token-group": ["Access: Organizations, Identity Providers, and Groups Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.groups", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
