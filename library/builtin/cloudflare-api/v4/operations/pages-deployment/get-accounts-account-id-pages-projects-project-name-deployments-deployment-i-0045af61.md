---
title: Get deployment info
page_id: operation-get-accounts-account-id-pages-projects-project-name-deployments-deployme-87f6e4e1
path: operations/pages-deployment
description: Fetch information about a deployment.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/pages/projects/{project_name}/deployments/{deployment_id}
operation_ids:
    - pages-deployment-get-deployment-info
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get deployment info

`GET /accounts/{account_id}/pages/projects/{project_name}/deployments/{deployment_id}`

Operation ID: `pages-deployment-get-deployment-info`

Fetch information about a deployment.

## Definition

```yaml
{"operationId": "pages-deployment-get-deployment-info", "summary": "Get deployment info", "description": "Fetch information about a deployment.", "parameters": [{"name": "deployment_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/pages_identifier"}}, {"name": "project_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/pages_project_name"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/pages_identifier"}}], "responses": {"200": {"description": "Get deployment info response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/pages_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/pages_deployment"}}, "required": ["result"], "type": "object"}]}}}}, "4XX": {"description": "Get deployment info response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/pages_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Pages Deployment"], "x-api-token-group": ["Pages Read", "Pages Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "pages.projects.deployments", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
