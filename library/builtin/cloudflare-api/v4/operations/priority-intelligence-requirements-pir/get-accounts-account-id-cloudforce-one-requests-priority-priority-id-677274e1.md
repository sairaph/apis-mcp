---
title: Get a Priority Intelligence Requirement
page_id: operation-get-accounts-account-id-cloudforce-one-requests-priority-priority-id-5b09fd8c
path: operations/priority-intelligence-requirements-pir
description: Retrieves a specific priority intelligence request from Cloudforce One.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/requests/priority/{priority_id}
operation_ids:
    - cloudforce-one-priority-get
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get a Priority Intelligence Requirement

`GET /accounts/{account_id}/cloudforce-one/requests/priority/{priority_id}`

Operation ID: `cloudforce-one-priority-get`

Retrieves a specific priority intelligence request from Cloudforce One.

## Definition

```yaml
{"operationId": "cloudforce-one-priority-get", "summary": "Get a Priority Intelligence Requirement", "description": "Retrieves a specific priority intelligence request from Cloudforce One.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cloudforce-one-requests_identifier"}}, {"name": "priority_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cloudforce-one-requests_uuid"}}], "responses": {"200": {"description": "Get priority response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/cloudforce-one-requests_api-response-common"}, {"properties": {"result": {"allOf": [{"$ref": "#/components/schemas/cloudforce-one-requests_request-item"}]}}, "type": "object"}]}}}}, "4XX": {"description": "Get priority response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cloudforce-one-requests_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Priority Intelligence Requirements (PIR)"], "x-api-token-group": ["Cloudforce One Write", "Cloudforce One Read"]}
```
