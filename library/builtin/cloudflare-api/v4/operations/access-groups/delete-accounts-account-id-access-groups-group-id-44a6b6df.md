---
title: Delete an Access group
page_id: operation-delete-accounts-account-id-access-groups-group-id-fe641854
path: operations/access-groups
description: Deletes an Access group.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/access/groups/{group_id}
operation_ids:
    - access-groups-delete-an-access-group
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete an Access group

`DELETE /accounts/{account_id}/access/groups/{group_id}`

Operation ID: `access-groups-delete-an-access-group`

Deletes an Access group.

## Definition

```yaml
{"operationId": "access-groups-delete-an-access-group", "summary": "Delete an Access group", "description": "Deletes an Access group.", "parameters": [{"name": "group_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_uuid"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}], "responses": {"202": {"description": "Delete an Access group response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_id_response"}}}}, "4XX": {"description": "Delete an Access group response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Access groups"], "x-api-token-group": ["Access: Organizations, Identity Providers, and Groups Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.access.groups", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
