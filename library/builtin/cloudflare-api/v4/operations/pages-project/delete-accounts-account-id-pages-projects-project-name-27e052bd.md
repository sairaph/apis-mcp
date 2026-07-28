---
title: Delete project
page_id: operation-delete-accounts-account-id-pages-projects-project-name-7d9c3957
path: operations/pages-project
description: Delete a project by name.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/pages/projects/{project_name}
operation_ids:
    - pages-project-delete-project
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete project

`DELETE /accounts/{account_id}/pages/projects/{project_name}`

Operation ID: `pages-project-delete-project`

Delete a project by name.

## Definition

```yaml
{"operationId": "pages-project-delete-project", "summary": "Delete project", "description": "Delete a project by name.", "parameters": [{"name": "project_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/pages_project_name"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/pages_identifier"}}], "responses": {"200": {"description": "Delete project response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/pages_api-response-common"}, {"properties": {"result": {"type": "object", "nullable": true}}, "required": ["result"], "type": "object"}]}}}}, "4XX": {"description": "Delete project response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/pages_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Pages Project"], "x-api-token-group": ["Pages Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "pages.projects", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
