---
title: List Access groups
page_id: operation-get-accounts-account-id-access-groups-dd46558b
path: operations/access-groups
description: Lists all Access groups.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/access/groups
operation_ids:
    - access-groups-list-access-groups
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Access groups

`GET /accounts/{account_id}/access/groups`

Operation ID: `access-groups-list-access-groups`

Lists all Access groups.

## Definition

```yaml
{"operationId": "access-groups-list-access-groups", "summary": "List Access groups", "description": "Lists all Access groups.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}, {"name": "name", "in": "query", "schema": {"description": "The name of the group.", "type": "string"}}, {"name": "search", "in": "query", "schema": {"description": "Search for groups by other listed query parameters.", "type": "string"}}, {"$ref": "#/components/parameters/access_page"}, {"$ref": "#/components/parameters/access_per_page"}], "responses": {"200": {"description": "List Access groups response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_response_collection-2"}}}}, "4XX": {"description": "List Access groups response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Access groups"], "x-api-token-group": ["Access: Organizations, Identity Providers, and Groups Write", "Access: Organizations, Identity Providers, and Groups Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.access.groups", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
