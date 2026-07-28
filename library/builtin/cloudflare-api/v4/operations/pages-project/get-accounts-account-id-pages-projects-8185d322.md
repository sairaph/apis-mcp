---
title: Get projects
page_id: operation-get-accounts-account-id-pages-projects-b64d33ac
path: operations/pages-project
description: Fetch a list of all user projects.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/pages/projects
operation_ids:
    - pages-project-get-projects
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get projects

`GET /accounts/{account_id}/pages/projects`

Operation ID: `pages-project-get-projects`

Fetch a list of all user projects.

## Definition

```yaml
{"operationId": "pages-project-get-projects", "summary": "Get projects", "description": "Fetch a list of all user projects.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/pages_identifier"}}, {"name": "page", "in": "query", "schema": {"description": "Which page of projects to fetch.", "type": "integer", "example": 1}}, {"name": "per_page", "in": "query", "schema": {"description": "How many projects to return per page.", "type": "integer", "example": 10}}], "responses": {"200": {"description": "Get projects response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/pages_api-response-collection"}, {"properties": {"result": {"type": "array", "items": {"$ref": "#/components/schemas/pages_project"}}}, "required": ["result"], "type": "object"}]}}}}, "4XX": {"description": "Get projects response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/pages_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Pages Project"], "x-api-token-group": ["Pages Read", "Pages Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "pages.projects", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
