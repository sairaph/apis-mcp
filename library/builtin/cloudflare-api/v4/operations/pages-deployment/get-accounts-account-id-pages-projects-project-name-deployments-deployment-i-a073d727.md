---
title: Get deployment logs
page_id: operation-get-accounts-account-id-pages-projects-project-name-deployments-deployme-ddd5a4a0
path: operations/pages-deployment
description: Fetch deployment logs for a project.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/pages/projects/{project_name}/deployments/{deployment_id}/history/logs
operation_ids:
    - pages-deployment-get-deployment-logs
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get deployment logs

`GET /accounts/{account_id}/pages/projects/{project_name}/deployments/{deployment_id}/history/logs`

Operation ID: `pages-deployment-get-deployment-logs`

Fetch deployment logs for a project.

## Definition

```yaml
{"operationId": "pages-deployment-get-deployment-logs", "summary": "Get deployment logs", "description": "Fetch deployment logs for a project.", "parameters": [{"name": "deployment_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/pages_identifier"}}, {"name": "project_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/pages_project_name"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/pages_identifier"}}], "responses": {"200": {"description": "Get deployment logs response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/pages_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/pages_deployment_log"}}, "required": ["result"], "type": "object"}]}}}}, "4XX": {"description": "Get deployment logs response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/pages_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Pages Deployment"], "x-api-token-group": ["Pages Read", "Pages Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "pages.projects.deployments.history.logs", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
