---
title: Get upload token
page_id: operation-get-accounts-account-id-pages-projects-project-name-upload-token-718243be
path: operations/pages-project
description: Get a short-lived JWT for Pages Direct Upload asset operations.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/pages/projects/{project_name}/upload-token
operation_ids:
    - pages-project-get-upload-token
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get upload token

`GET /accounts/{account_id}/pages/projects/{project_name}/upload-token`

Operation ID: `pages-project-get-upload-token`

Get a short-lived JWT for Pages Direct Upload asset operations.

## Definition

```yaml
{"operationId": "pages-project-get-upload-token", "summary": "Get upload token", "description": "Get a short-lived JWT for Pages Direct Upload asset operations.", "parameters": [{"name": "project_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/pages_project_name"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/pages_identifier"}}], "responses": {"200": {"description": "Upload token response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/pages_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/pages_pages_upload_token_response"}}, "required": ["result"], "type": "object"}]}}}}, "4XX": {"description": "Upload token failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/pages_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Pages Project"], "x-api-token-group": ["Pages Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "pages.projects", "x-fern-sdk-method-name": "get-upload-token", "x-forge-hidden": true}
```
