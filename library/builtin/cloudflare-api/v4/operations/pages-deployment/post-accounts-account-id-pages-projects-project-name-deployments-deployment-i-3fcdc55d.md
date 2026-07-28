---
title: Rollback deployment
page_id: operation-post-accounts-account-id-pages-projects-project-name-deployments-deploym-a134058f
path: operations/pages-deployment
description: Rollback the production deployment to a previous deployment. You can only rollback to succesful builds on production.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/pages/projects/{project_name}/deployments/{deployment_id}/rollback
operation_ids:
    - pages-deployment-rollback-deployment
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Rollback deployment

`POST /accounts/{account_id}/pages/projects/{project_name}/deployments/{deployment_id}/rollback`

Operation ID: `pages-deployment-rollback-deployment`

Rollback the production deployment to a previous deployment. You can only rollback to succesful builds on production.

## Definition

```yaml
{"operationId": "pages-deployment-rollback-deployment", "summary": "Rollback deployment", "description": "Rollback the production deployment to a previous deployment. You can only rollback to succesful builds on production.", "parameters": [{"name": "deployment_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/pages_identifier"}}, {"name": "project_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/pages_project_name"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/pages_identifier"}}], "responses": {"200": {"description": "Rollback deployment response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/pages_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/pages_deployment"}}, "required": ["result"], "type": "object"}]}}}}, "4XX": {"description": "Rollback deployment response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/pages_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Pages Deployment"], "x-api-token-group": ["Pages Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "pages.projects.deployments", "x-fern-sdk-method-name": "rollback", "x-forge-hidden": true}
```
