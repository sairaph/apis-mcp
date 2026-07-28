---
title: Delete a tag
page_id: operation-delete-accounts-account-id-access-tags-tag-name-90e5bc71
path: operations/access-tags
description: Delete a tag
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/access/tags/{tag_name}
operation_ids:
    - access-tags-delete-a-tag
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete a tag

`DELETE /accounts/{account_id}/access/tags/{tag_name}`

Operation ID: `access-tags-delete-a-tag`

Delete a tag

## Definition

```yaml
{"operationId": "access-tags-delete-a-tag", "summary": "Delete a tag", "description": "Delete a tag", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}, {"name": "tag_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_name-13"}}], "responses": {"202": {"description": "Delete a tag response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_name_response"}}}}, "4XX": {"description": "Delete a tag response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Access tags"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.access.tags", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
