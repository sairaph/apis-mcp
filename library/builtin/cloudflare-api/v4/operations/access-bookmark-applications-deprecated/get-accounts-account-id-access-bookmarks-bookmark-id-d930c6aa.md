---
title: Get a Bookmark application
page_id: operation-get-accounts-account-id-access-bookmarks-bookmark-id-0b0083bf
path: operations/access-bookmark-applications-deprecated
description: Fetches a single Bookmark application.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/access/bookmarks/{bookmark_id}
operation_ids:
    - access-bookmark-applications-(-deprecated)-get-a-bookmark-application
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get a Bookmark application

`GET /accounts/{account_id}/access/bookmarks/{bookmark_id}`

Operation ID: `access-bookmark-applications-(-deprecated)-get-a-bookmark-application`

Fetches a single Bookmark application.

## Definition

```yaml
{"operationId": "access-bookmark-applications-(-deprecated)-get-a-bookmark-application", "summary": "Get a Bookmark application", "description": "Fetches a single Bookmark application.", "parameters": [{"name": "bookmark_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_uuid"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/access_identifier-3"}}], "responses": {"200": {"description": "Get a Bookmark application response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_single_response-13"}}}}, "4XX": {"description": "Get a Bookmark application response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/access_api-response-common-failure"}}}}}, "deprecated": true, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Access Bookmark applications (Deprecated)"], "x-api-token-group": ["Access: Apps and Policies Write", "Access: Apps and Policies Read"], "x-cfDeprecation": {"description": "This endpoint is deprecated in favor of using a specialized Access Application App Type.", "display": true, "eol": "2023-03-19", "id": "bookmarks_deprecation"}, "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.access.bookmarks", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
