---
title: Disconnect project source
page_id: operation-delete-accounts-account-id-pages-projects-project-name-source-f9e5f29f
path: operations/pages-project
description: Disconnect the Git repository source from an existing Pages project.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/pages/projects/{project_name}/source
operation_ids:
    - pages-project-disconnect-project-source
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Disconnect project source

`DELETE /accounts/{account_id}/pages/projects/{project_name}/source`

Operation ID: `pages-project-disconnect-project-source`

Disconnect the Git repository source from an existing Pages project.

## Definition

```yaml
{"operationId": "pages-project-disconnect-project-source", "summary": "Disconnect project source", "description": "Disconnect the Git repository source from an existing Pages project.", "parameters": [{"name": "project_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/pages_project_name"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/pages_identifier"}}], "responses": {"200": {"description": "Disconnect project source response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/pages_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/pages_project"}}, "required": ["result"], "type": "object"}]}}}}, "4XX": {"description": "Disconnect project source response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/pages_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Pages Project"], "x-api-token-group": ["Pages Write"]}
```
