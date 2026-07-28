---
title: Connect project source
page_id: operation-post-accounts-account-id-pages-projects-project-name-source-376d684e
path: operations/pages-project
description: Connect a Git repository source to an existing Pages project.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/pages/projects/{project_name}/source
operation_ids:
    - pages-project-connect-project-source
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Connect project source

`POST /accounts/{account_id}/pages/projects/{project_name}/source`

Operation ID: `pages-project-connect-project-source`

Connect a Git repository source to an existing Pages project.

## Definition

```yaml
{"operationId": "pages-project-connect-project-source", "summary": "Connect project source", "description": "Connect a Git repository source to an existing Pages project.", "parameters": [{"name": "project_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/pages_project_name"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/pages_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/pages_source"}}}}, "responses": {"200": {"description": "Connect project source response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/pages_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/pages_project"}}, "required": ["result"], "type": "object"}]}}}}, "4XX": {"description": "Connect project source response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/pages_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Pages Project"], "x-api-token-group": ["Pages Write"]}
```
