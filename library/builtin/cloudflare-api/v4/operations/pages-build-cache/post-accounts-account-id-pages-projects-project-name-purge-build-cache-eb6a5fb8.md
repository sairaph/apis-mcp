---
title: Purge build cache
page_id: operation-post-accounts-account-id-pages-projects-project-name-purge-build-cache-3db547ea
path: operations/pages-build-cache
description: Purge all cached build artifacts for a Pages project
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/pages/projects/{project_name}/purge_build_cache
operation_ids:
    - pages-purge-build-cache
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Purge build cache

`POST /accounts/{account_id}/pages/projects/{project_name}/purge_build_cache`

Operation ID: `pages-purge-build-cache`

Purge all cached build artifacts for a Pages project

## Definition

```yaml
{"operationId": "pages-purge-build-cache", "summary": "Purge build cache", "description": "Purge all cached build artifacts for a Pages project", "parameters": [{"name": "project_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/pages_project_name"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/pages_identifier"}}], "responses": {"200": {"description": "Purge build cache response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/pages_api-response-common"}, {"properties": {"result": {"type": "object", "nullable": true}}, "required": ["result"], "type": "object"}]}}}}, "4XX": {"description": "Purge build cache failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/pages_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Pages Build Cache"], "x-api-token-group": ["Pages Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "pages.projects", "x-fern-sdk-method-name": "purge-build-cache", "x-forge-hidden": true, "x-forge-require-confirmation": "This operation purges the build cache for a Pages project."}
```
