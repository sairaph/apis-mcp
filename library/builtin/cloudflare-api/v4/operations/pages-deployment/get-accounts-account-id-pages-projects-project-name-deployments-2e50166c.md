---
title: Get deployments
page_id: operation-get-accounts-account-id-pages-projects-project-name-deployments-ab76ab26
path: operations/pages-deployment
description: Fetch a list of project deployments.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/pages/projects/{project_name}/deployments
operation_ids:
    - pages-deployment-get-deployments
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get deployments

`GET /accounts/{account_id}/pages/projects/{project_name}/deployments`

Operation ID: `pages-deployment-get-deployments`

Fetch a list of project deployments.

## Definition

```yaml
{"operationId": "pages-deployment-get-deployments", "summary": "Get deployments", "description": "Fetch a list of project deployments.", "parameters": [{"name": "project_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/pages_project_name"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/pages_identifier"}}, {"name": "env", "in": "query", "schema": {"description": "What type of deployments to fetch.", "type": "string", "example": "preview", "enum": ["production", "preview"]}}, {"name": "page", "in": "query", "schema": {"description": "Which page of deployments to fetch.", "type": "integer", "example": 1}}, {"name": "per_page", "in": "query", "schema": {"description": "How many deployments to return per page.", "type": "integer", "example": 10}}], "responses": {"200": {"description": "Get deployments response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/pages_api-response-collection"}, {"properties": {"result": {"type": "array", "items": {"$ref": "#/components/schemas/pages_deployment"}}}, "required": ["result"], "type": "object"}]}}}}, "4XX": {"description": "Get deployments response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/pages_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Pages Deployment"], "x-api-token-group": ["Pages Read", "Pages Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "pages.projects.deployments", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
