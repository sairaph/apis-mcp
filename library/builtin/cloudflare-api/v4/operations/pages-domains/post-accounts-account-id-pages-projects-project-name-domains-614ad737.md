---
title: Add domain
page_id: operation-post-accounts-account-id-pages-projects-project-name-domains-a234398e
path: operations/pages-domains
description: Add a new domain for the Pages project.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/pages/projects/{project_name}/domains
operation_ids:
    - pages-domains-add-domain
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Add domain

`POST /accounts/{account_id}/pages/projects/{project_name}/domains`

Operation ID: `pages-domains-add-domain`

Add a new domain for the Pages project.

## Definition

```yaml
{"operationId": "pages-domains-add-domain", "summary": "Add domain", "description": "Add a new domain for the Pages project.", "parameters": [{"name": "project_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/pages_project_name"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/pages_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"name": {"$ref": "#/components/schemas/pages_domain_name"}}, "required": ["name"]}}}}, "responses": {"200": {"description": "Add domain response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/pages_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/pages_domain"}}, "required": ["result"], "type": "object"}]}}}}, "4XX": {"description": "Add domain response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/pages_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Pages Domains"], "x-api-token-group": ["Pages Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "pages.projects.domains", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
