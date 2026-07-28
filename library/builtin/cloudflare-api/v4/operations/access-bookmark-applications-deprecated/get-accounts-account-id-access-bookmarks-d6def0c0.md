---
title: List Bookmark applications
page_id: operation-get-accounts-account-id-access-bookmarks-5b52d859
path: operations/access-bookmark-applications-deprecated
description: Lists Bookmark applications.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/access/bookmarks
operation_ids:
    - access-bookmark-applications-(-deprecated)-list-bookmark-applications
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Bookmark applications

`GET /accounts/{account_id}/access/bookmarks`

Operation ID: `access-bookmark-applications-(-deprecated)-list-bookmark-applications`

Lists Bookmark applications.

## Definition

```yaml
{"operationId": "access-bookmark-applications-(-deprecated)-list-bookmark-applications", "summary": "List Bookmark applications", "description": "Lists Bookmark applications.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier-3"}}], "responses": {"200": {"description": "List Bookmark applications response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_response_collection-14"}}}}, "4XX": {"description": "List Bookmark applications response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "deprecated": true, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Access Bookmark applications (Deprecated)"], "x-api-token-group": ["Access: Apps and Policies Revoke", "Access: Apps and Policies Write", "Access: Apps and Policies Read"], "x-cfDeprecation": {"description": "This endpoint is deprecated in favor of using a specialized Access Application App Type.", "display": true, "eol": "2023-03-19", "id": "bookmarks_deprecation"}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.access.bookmarks", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
