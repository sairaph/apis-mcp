---
title: Create a Bookmark application
page_id: operation-post-accounts-account-id-access-bookmarks-bookmark-id-367c2afc
path: operations/access-bookmark-applications-deprecated
description: Create a new Bookmark application.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/access/bookmarks/{bookmark_id}
operation_ids:
    - access-bookmark-applications-(-deprecated)-create-a-bookmark-application
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a Bookmark application

`POST /accounts/{account_id}/access/bookmarks/{bookmark_id}`

Operation ID: `access-bookmark-applications-(-deprecated)-create-a-bookmark-application`

Create a new Bookmark application.

## Definition

```yaml
{"operationId": "access-bookmark-applications-(-deprecated)-create-a-bookmark-application", "summary": "Create a Bookmark application", "description": "Create a new Bookmark application.", "parameters": [{"name": "bookmark_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_uuid"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier-3"}}], "requestBody": {"required": true, "content": {"application/json": {}}}, "responses": {"200": {"description": "Create a Bookmark application response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_single_response-13"}}}}, "4XX": {"description": "Create a Bookmark application response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "deprecated": true, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Access Bookmark applications (Deprecated)"], "x-cfDeprecation": {"description": "This endpoint is deprecated in favor of using a specialized Access Application App Type.", "display": true, "eol": "2023-03-19", "id": "bookmarks_deprecation"}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.access.bookmarks", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
