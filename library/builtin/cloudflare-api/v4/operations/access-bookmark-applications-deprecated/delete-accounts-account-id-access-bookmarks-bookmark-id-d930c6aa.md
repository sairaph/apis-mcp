---
title: Delete a Bookmark application
page_id: operation-delete-accounts-account-id-access-bookmarks-bookmark-id-2cccc839
path: operations/access-bookmark-applications-deprecated
description: Deletes a Bookmark application.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/access/bookmarks/{bookmark_id}
operation_ids:
    - access-bookmark-applications-(-deprecated)-delete-a-bookmark-application
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete a Bookmark application

`DELETE /accounts/{account_id}/access/bookmarks/{bookmark_id}`

Operation ID: `access-bookmark-applications-(-deprecated)-delete-a-bookmark-application`

Deletes a Bookmark application.

## Definition

```yaml
{"operationId": "access-bookmark-applications-(-deprecated)-delete-a-bookmark-application", "summary": "Delete a Bookmark application", "description": "Deletes a Bookmark application.", "parameters": [{"name": "bookmark_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_uuid"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier-3"}}], "requestBody": {"required": true, "content": {"application/json": {}}}, "responses": {"200": {"description": "Delete a Bookmark application response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_id_response"}}}}, "4XX": {"description": "Delete a Bookmark application response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "deprecated": true, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Access Bookmark applications (Deprecated)"], "x-api-token-group": ["Access: Apps and Policies Write"], "x-cfDeprecation": {"description": "This endpoint is deprecated in favor of using a specialized Access Application App Type.", "display": true, "eol": "2023-03-19", "id": "bookmarks_deprecation"}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.access.bookmarks", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
