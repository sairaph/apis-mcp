---
title: Delete domain
page_id: operation-delete-accounts-account-id-pages-projects-project-name-domains-domain-na-e808dc79
path: operations/pages-domains
description: Delete a Pages project's domain.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/pages/projects/{project_name}/domains/{domain_name}
operation_ids:
    - pages-domains-delete-domain
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete domain

`DELETE /accounts/{account_id}/pages/projects/{project_name}/domains/{domain_name}`

Operation ID: `pages-domains-delete-domain`

Delete a Pages project's domain.

## Definition

```yaml
{"operationId": "pages-domains-delete-domain", "summary": "Delete domain", "description": "Delete a Pages project's domain.", "parameters": [{"name": "domain_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/pages_domain_name"}}, {"name": "project_name", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/pages_project_name"}}, {"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/pages_identifier"}}], "responses": {"200": {"description": "Delete domain response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/pages_api-response-common"}, {"properties": {"result": {"type": "object", "nullable": true}}, "required": ["result"], "type": "object"}]}}}}, "4XX": {"description": "Delete domain response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/pages_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Pages Domains"], "x-api-token-group": ["Pages Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "pages.projects.domains", "x-fern-sdk-method-name": "delete", "x-forge-hidden": true}
```
