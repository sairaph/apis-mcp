---
title: Delete a custom page
page_id: operation-delete-accounts-account-id-access-custom-pages-custom-page-id-b7798c8b
path: operations/access-custom-pages
description: Delete a custom page
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/access/custom_pages/{custom_page_id}
operation_ids:
    - access-custom-pages-delete-a-custom-page
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete a custom page

`DELETE /accounts/{account_id}/access/custom_pages/{custom_page_id}`

Operation ID: `access-custom-pages-delete-a-custom-page`

Delete a custom page

## Definition

```yaml
{"operationId": "access-custom-pages-delete-a-custom-page", "summary": "Delete a custom page", "description": "Delete a custom page", "parameters": [{"name": "custom_page_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_uuid"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier"}}], "responses": {"202": {"description": "Delete a custom page response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_id_response-3"}}}}, "4XX": {"description": "Delete a custom page response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Access custom pages"], "x-api-token-group": ["Access: Custom Pages Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.access.custom-pages", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
