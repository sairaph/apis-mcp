---
title: Get project
page_id: operation-get-accounts-account-id-pages-projects-project-name-8e5d2988
path: operations/pages-project
description: Fetch a project by name.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/pages/projects/{project_name}
operation_ids:
    - pages-project-get-project
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get project

`GET /accounts/{account_id}/pages/projects/{project_name}`

Operation ID: `pages-project-get-project`

Fetch a project by name.

## Definition

```yaml
{"operationId": "pages-project-get-project", "summary": "Get project", "description": "Fetch a project by name.", "parameters": [{"name": "project_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/pages_project_name"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/pages_identifier"}}], "responses": {"200": {"description": "Get project response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/pages_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/pages_project"}}, "required": ["result"], "type": "object"}]}}}}, "4XX": {"description": "Get project response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/pages_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Pages Project"], "x-api-token-group": ["Pages Read", "Pages Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "pages.projects", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
