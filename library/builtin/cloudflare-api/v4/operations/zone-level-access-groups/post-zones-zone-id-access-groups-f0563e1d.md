---
title: Create an Access group
page_id: operation-post-zones-zone-id-access-groups-16aeffde
path: operations/zone-level-access-groups
description: Creates a new Access group.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /zones/{zone_id}/access/groups
operation_ids:
    - zone-level-access-groups-create-an-access-group
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create an Access group

`POST /zones/{zone_id}/access/groups`

Operation ID: `zone-level-access-groups-create-an-access-group`

Creates a new Access group.

## Definition

```yaml
{"operationId": "zone-level-access-groups-create-an-access-group", "summary": "Create an Access group", "description": "Creates a new Access group.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"properties": {"exclude": {"$ref": "#/components/schemas/access_exclude"}, "include": {"$ref": "#/components/schemas/access_include"}, "name": {"$ref": "#/components/schemas/access_name-16"}, "require": {"$ref": "#/components/schemas/access_require"}}, "required": ["name", "include"]}}}}, "responses": {"201": {"description": "Create an Access group response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_single_response-17"}}}}, "4XX": {"description": "Create an Access group response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Zone-Level Access groups"], "x-api-token-group": ["Access: Organizations, Identity Providers, and Groups Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.groups", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
