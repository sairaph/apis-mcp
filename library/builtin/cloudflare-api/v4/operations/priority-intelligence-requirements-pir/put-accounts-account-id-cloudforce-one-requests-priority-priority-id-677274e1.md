---
title: Update a Priority Intelligence Requirement
page_id: operation-put-accounts-account-id-cloudforce-one-requests-priority-priority-id-d24fc8bb
path: operations/priority-intelligence-requirements-pir
description: Updates a priority intelligence request in Cloudforce One.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/requests/priority/{priority_id}
operation_ids:
    - cloudforce-one-priority-update
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update a Priority Intelligence Requirement

`PUT /accounts/{account_id}/cloudforce-one/requests/priority/{priority_id}`

Operation ID: `cloudforce-one-priority-update`

Updates a priority intelligence request in Cloudforce One.

## Definition

```yaml
{"operationId": "cloudforce-one-priority-update", "summary": "Update a Priority Intelligence Requirement", "description": "Updates a priority intelligence request in Cloudforce One.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cloudforce-one-requests_identifier"}}, {"name": "priority_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cloudforce-one-requests_uuid"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cloudforce-one-requests_priority-edit"}}}}, "responses": {"200": {"description": "Update priority response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/cloudforce-one-requests_api-response-common"}, {"properties": {"result": {"allOf": [{"$ref": "#/components/schemas/cloudforce-one-requests_request-item"}]}}, "type": "object"}]}}}}, "4XX": {"description": "Update priority response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cloudforce-one-requests_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Priority Intelligence Requirements (PIR)"], "x-api-token-group": ["Cloudforce One Write"]}
```
