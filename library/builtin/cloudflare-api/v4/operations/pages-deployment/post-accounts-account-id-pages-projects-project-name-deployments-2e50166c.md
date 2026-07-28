---
title: Create deployment
page_id: operation-post-accounts-account-id-pages-projects-project-name-deployments-26af1d7d
path: operations/pages-deployment
description: Start a new deployment from production. The repository and account must have already been authorized on the Cloudflare Pages dashboard.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/pages/projects/{project_name}/deployments
operation_ids:
    - pages-deployment-create-deployment
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create deployment

`POST /accounts/{account_id}/pages/projects/{project_name}/deployments`

Operation ID: `pages-deployment-create-deployment`

Start a new deployment from production. The repository and account must have already been authorized on the Cloudflare Pages dashboard.

## Definition

```yaml
{"operationId": "pages-deployment-create-deployment", "summary": "Create deployment", "description": "Start a new deployment from production. The repository and account must have already been authorized on the Cloudflare Pages dashboard.", "parameters": [{"name": "project_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/pages_project_name"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/pages_identifier"}}], "requestBody": {"required": true, "content": {"multipart/form-data": {"encoding": {"_headers": {"contentType": "text/plain"}, "_redirects": {"contentType": "text/plain"}, "_routes.json": {"contentType": "application/json"}, "_worker.bundle": {"contentType": "multipart/form-data"}, "_worker.js": {"contentType": "application/javascript+module, text/javascript+module, application/javascript, text/javascript"}, "functions-filepath-routing-config.json": {"contentType": "application/json"}, "manifest": {"contentType": "application/json"}}, "schema": {"type": "object", "properties": {"_headers": {"description": "Headers configuration file for the deployment.", "type": "string", "format": "binary"}, "_redirects": {"description": "Redirects configuration file for the deployment.", "type": "string", "format": "binary"}, "_routes.json": {"description": "Routes configuration file defining routing rules.", "type": "string", "format": "binary"}, "_worker.bundle": {"description": "Worker bundle file in multipart/form-data format. Mutually exclusive with `_worker.js`.\nCannot specify both `_worker.js` and `_worker.bundle` in the same request.\nMaximum size: 25 MiB.\n", "type": "string", "format": "binary"}, "_worker.js": {"description": "Worker JavaScript file. Mutually exclusive with `_worker.bundle`.\nCannot specify both `_worker.js` and `_worker.bundle` in the same request.\n", "type": "string", "format": "binary"}, "branch": {"description": "The branch to build the new deployment from. The `HEAD` of the branch will be used. If omitted, the production branch will be used by default.", "type": "string", "example": "staging", "x-auditable": true}, "commit_dirty": {"description": "Boolean string indicating if the working directory has uncommitted changes.", "type": "string", "example": "false", "enum": ["true", "false"]}, "commit_hash": {"description": "Git commit SHA associated with this deployment.", "type": "string", "example": "a1b2c3d4e5f6"}, "commit_message": {"description": "Git commit message associated with this deployment.", "type": "string", "example": "Update homepage"}, "functions-filepath-routing-config.json": {"description": "Functions routing configuration file.", "type": "string", "format": "binary"}, "manifest": {"description": "JSON string containing a manifest of files to deploy. Maps file paths to their content hashes.\nRequired for direct upload deployments. Maximum 20,000 entries.\n", "type": "string", "example": "{\"index.html\": \"abc123\", \"style.css\": \"def456\"}"}, "pages_build_output_dir": {"description": "The build output directory path.", "type": "string", "example": "dist"}, "wrangler_config_hash": {"description": "Hash of the Wrangler configuration file used for this deployment.", "type": "string"}}}}}}, "responses": {"200": {"description": "Create deployment response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/pages_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/pages_deployment"}}, "required": ["result"], "type": "object"}]}}}}, "4XX": {"description": "Create deployment response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/pages_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Pages Deployment"], "x-api-token-group": ["Pages Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "pages.projects.deployments", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
