---
title: Retry deployment
page_id: operation-post-accounts-account-id-pages-projects-project-name-deployments-deploym-3d74655d
path: operations/pages-deployment
description: Retry a previous deployment.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/pages/projects/{project_name}/deployments/{deployment_id}/retry
operation_ids:
    - pages-deployment-retry-deployment
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Retry deployment

`POST /accounts/{account_id}/pages/projects/{project_name}/deployments/{deployment_id}/retry`

Operation ID: `pages-deployment-retry-deployment`

Retry a previous deployment.

## Definition

```yaml
{"operationId": "pages-deployment-retry-deployment", "summary": "Retry deployment", "description": "Retry a previous deployment.", "parameters": [{"name": "deployment_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/pages_identifier"}}, {"name": "project_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/pages_project_name"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/pages_identifier"}}], "responses": {"200": {"description": "Retry deployment response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/pages_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/pages_deployment"}}, "required": ["result"], "type": "object"}]}}}}, "4XX": {"description": "Retry deployment response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/pages_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Pages Deployment"], "x-api-token-group": ["Pages Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "pages.projects.deployments", "x-fern-sdk-method-name": "retry", "x-forge-hidden": true}
```
