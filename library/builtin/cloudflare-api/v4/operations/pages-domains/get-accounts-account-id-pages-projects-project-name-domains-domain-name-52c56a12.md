---
title: Get domain
page_id: operation-get-accounts-account-id-pages-projects-project-name-domains-domain-name-9bba1124
path: operations/pages-domains
description: Fetch a single domain.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/pages/projects/{project_name}/domains/{domain_name}
operation_ids:
    - pages-domains-get-domain
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get domain

`GET /accounts/{account_id}/pages/projects/{project_name}/domains/{domain_name}`

Operation ID: `pages-domains-get-domain`

Fetch a single domain.

## Definition

```yaml
{"operationId": "pages-domains-get-domain", "summary": "Get domain", "description": "Fetch a single domain.", "parameters": [{"name": "domain_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/pages_domain_name"}}, {"name": "project_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/pages_project_name"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/pages_identifier"}}], "responses": {"200": {"description": "Get domain response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/pages_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/pages_domain"}}, "required": ["result"], "type": "object"}]}}}}, "4XX": {"description": "Get domain response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/pages_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Pages Domains"], "x-api-token-group": ["Pages Read", "Pages Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "pages.projects.domains", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
