---
title: Get an Access group
page_id: operation-get-accounts-account-id-access-groups-group-id-f726e6ca
path: operations/access-groups
description: Fetches a single Access group.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/access/groups/{group_id}
operation_ids:
    - access-groups-get-an-access-group
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get an Access group

`GET /accounts/{account_id}/access/groups/{group_id}`

Operation ID: `access-groups-get-an-access-group`

Fetches a single Access group.

## Definition

```yaml
{"operationId": "access-groups-get-an-access-group", "summary": "Get an Access group", "description": "Fetches a single Access group.", "parameters": [{"name": "group_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_uuid"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}], "responses": {"200": {"description": "Get an Access group response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_single_response-4"}}}}, "4XX": {"description": "Get an Access group response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Access groups"], "x-api-token-group": ["Access: Organizations, Identity Providers, and Groups Write", "Access: Organizations, Identity Providers, and Groups Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.access.groups", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
