---
title: Get an Access group
page_id: operation-get-zones-zone-id-access-groups-group-id-1528c9a5
path: operations/zone-level-access-groups
description: Fetches a single Access group.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/access/groups/{group_id}
operation_ids:
    - zone-level-access-groups-get-an-access-group
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get an Access group

`GET /zones/{zone_id}/access/groups/{group_id}`

Operation ID: `zone-level-access-groups-get-an-access-group`

Fetches a single Access group.

## Definition

```yaml
{"operationId": "zone-level-access-groups-get-an-access-group", "summary": "Get an Access group", "description": "Fetches a single Access group.", "parameters": [{"name": "group_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_uuid"}}, {"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}], "responses": {"200": {"description": "Get an Access group response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_single_response-17"}}}}, "4XX": {"description": "Get an Access group response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Zone-Level Access groups"], "x-api-token-group": ["Access: Organizations, Identity Providers, and Groups Write", "Access: Organizations, Identity Providers, and Groups Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.groups", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
