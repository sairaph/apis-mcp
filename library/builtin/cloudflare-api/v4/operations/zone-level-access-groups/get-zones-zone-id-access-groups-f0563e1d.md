---
title: List Access groups
page_id: operation-get-zones-zone-id-access-groups-8f6625fd
path: operations/zone-level-access-groups
description: Lists all Access groups.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /zones/{zone_id}/access/groups
operation_ids:
    - zone-level-access-groups-list-access-groups
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Access groups

`GET /zones/{zone_id}/access/groups`

Operation ID: `zone-level-access-groups-list-access-groups`

Lists all Access groups.

## Definition

```yaml
{"operationId": "zone-level-access-groups-list-access-groups", "summary": "List Access groups", "description": "Lists all Access groups.", "parameters": [{"name": "zone_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}], "responses": {"200": {"description": "List Access groups response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_response_collection-18"}}}}, "4XX": {"description": "List Access groups response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Zone-Level Access groups"], "x-api-token-group": ["Access: Organizations, Identity Providers, and Groups Write", "Access: Organizations, Identity Providers, and Groups Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.groups", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
