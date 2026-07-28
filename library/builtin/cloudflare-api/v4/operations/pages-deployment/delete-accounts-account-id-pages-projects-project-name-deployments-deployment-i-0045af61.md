---
title: Delete deployment
page_id: operation-delete-accounts-account-id-pages-projects-project-name-deployments-deplo-b682b02a
path: operations/pages-deployment
description: Delete a deployment.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/pages/projects/{project_name}/deployments/{deployment_id}
operation_ids:
    - pages-deployment-delete-deployment
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete deployment

`DELETE /accounts/{account_id}/pages/projects/{project_name}/deployments/{deployment_id}`

Operation ID: `pages-deployment-delete-deployment`

Delete a deployment.

## Definition

```yaml
{"operationId": "pages-deployment-delete-deployment", "summary": "Delete deployment", "description": "Delete a deployment.", "parameters": [{"name": "deployment_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/pages_identifier"}}, {"name": "project_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/pages_project_name"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/pages_identifier"}}, {"name": "force", "in": "query", "schema": {"description": "Allow deletion of aliased non-production deployments when a normal delete would be rejected.", "type": "boolean", "example": true}}], "responses": {"200": {"description": "Delete deployment response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/pages_api-response-common"}, {"properties": {"result": {"type": "object", "nullable": true}}, "required": ["result"], "type": "object"}]}}}}, "4XX": {"description": "Delete deployment response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/pages_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Pages Deployment"], "x-api-token-group": ["Pages Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "pages.projects.deployments", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
