---
title: Get domains
page_id: operation-get-accounts-account-id-pages-projects-project-name-domains-dc098122
path: operations/pages-domains
description: Fetch a list of all domains associated with a Pages project.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/pages/projects/{project_name}/domains
operation_ids:
    - pages-domains-get-domains
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get domains

`GET /accounts/{account_id}/pages/projects/{project_name}/domains`

Operation ID: `pages-domains-get-domains`

Fetch a list of all domains associated with a Pages project.

## Definition

```yaml
{"operationId": "pages-domains-get-domains", "summary": "Get domains", "description": "Fetch a list of all domains associated with a Pages project.", "parameters": [{"name": "project_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/pages_project_name"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/pages_identifier"}}], "responses": {"200": {"description": "Get domains response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/pages_api-response-collection"}, {"properties": {"result": {"type": "array", "items": {"$ref": "#/components/schemas/pages_domain"}}}, "required": ["result"], "type": "object"}]}}}}, "4XX": {"description": "Get domains response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/pages_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Pages Domains"], "x-api-token-group": ["Pages Read", "Pages Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "pages.projects.domains", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
