---
title: Create a New Priority Intelligence Requirement
page_id: operation-post-accounts-account-id-cloudforce-one-requests-priority-new-24ca6ab9
path: operations/priority-intelligence-requirements-pir
description: Creates a new priority intelligence request in Cloudforce One.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/requests/priority/new
operation_ids:
    - cloudforce-one-priority-new
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a New Priority Intelligence Requirement

`POST /accounts/{account_id}/cloudforce-one/requests/priority/new`

Operation ID: `cloudforce-one-priority-new`

Creates a new priority intelligence request in Cloudforce One.

## Definition

```yaml
{"operationId": "cloudforce-one-priority-new", "summary": "Create a New Priority Intelligence Requirement", "description": "Creates a new priority intelligence request in Cloudforce One.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cloudforce-one-requests_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cloudforce-one-requests_priority-edit"}}}}, "responses": {"200": {"description": "Create priority response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/cloudforce-one-requests_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/cloudforce-one-requests_priority-item"}}, "type": "object"}]}}}}, "4XX": {"description": "Create priority response  failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cloudforce-one-requests_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Priority Intelligence Requirements (PIR)"], "x-api-token-group": ["Cloudforce One Write"]}
```
