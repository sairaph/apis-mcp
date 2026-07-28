---
title: Create deployment tail
page_id: operation-post-accounts-account-id-pages-projects-project-name-deployments-deploym-27a15b3c
path: operations/pages-deployment
description: Start a tail that receives logs and exception data.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/pages/projects/{project_name}/deployments/{deployment_id}/tails
operation_ids:
    - pages-deployment-create-tail
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create deployment tail

`POST /accounts/{account_id}/pages/projects/{project_name}/deployments/{deployment_id}/tails`

Operation ID: `pages-deployment-create-tail`

Start a tail that receives logs and exception data.

## Definition

```yaml
{"operationId": "pages-deployment-create-tail", "summary": "Create deployment tail", "description": "Start a tail that receives logs and exception data.", "parameters": [{"name": "deployment_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/pages_identifier"}}, {"name": "project_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/pages_project_name"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/pages_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"filters": {"description": "Filters to apply to the tail session.", "type": "array", "items": {"additionalProperties": true, "type": "object"}}}, "example": {"filters": [{"outcome": ["exception"]}]}}}}}, "responses": {"200": {"description": "Create deployment tail response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/pages_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/pages_tail"}}, "required": ["result"], "type": "object"}]}}}}, "4XX": {"description": "Create deployment tail response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/pages_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Pages Deployment"], "x-api-token-group": ["Pages Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "pages.projects.deployments.tails", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
