---
title: Update a Bookmark application
page_id: operation-put-accounts-account-id-access-bookmarks-bookmark-id-91a391e6
path: operations/access-bookmark-applications-deprecated
description: Updates a configured Bookmark application.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/access/bookmarks/{bookmark_id}
operation_ids:
    - access-bookmark-applications-(-deprecated)-update-a-bookmark-application
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update a Bookmark application

`PUT /accounts/{account_id}/access/bookmarks/{bookmark_id}`

Operation ID: `access-bookmark-applications-(-deprecated)-update-a-bookmark-application`

Updates a configured Bookmark application.

## Definition

```yaml
{"operationId": "access-bookmark-applications-(-deprecated)-update-a-bookmark-application", "summary": "Update a Bookmark application", "description": "Updates a configured Bookmark application.", "parameters": [{"name": "bookmark_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_uuid"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier-3"}}], "requestBody": {"required": true, "content": {"application/json": {}}}, "responses": {"200": {"description": "Update a Bookmark application response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_single_response-13"}}}}, "4XX": {"description": "Update a Bookmark application response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "deprecated": true, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Access Bookmark applications (Deprecated)"], "x-api-token-group": ["Access: Apps and Policies Write"], "x-cfDeprecation": {"description": "This endpoint is deprecated in favor of using a specialized Access Application App Type.", "display": true, "eol": "2023-03-19", "id": "bookmarks_deprecation"}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.access.bookmarks", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
