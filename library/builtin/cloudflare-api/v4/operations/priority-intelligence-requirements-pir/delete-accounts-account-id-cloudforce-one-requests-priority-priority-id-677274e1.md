---
title: Delete a Priority Intelligence Requirement
page_id: operation-delete-accounts-account-id-cloudforce-one-requests-priority-priority-id-c38e88fe
path: operations/priority-intelligence-requirements-pir
description: Deletes a priority intelligence request from Cloudforce One.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/requests/priority/{priority_id}
operation_ids:
    - cloudforce-one-priority-delete
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete a Priority Intelligence Requirement

`DELETE /accounts/{account_id}/cloudforce-one/requests/priority/{priority_id}`

Operation ID: `cloudforce-one-priority-delete`

Deletes a priority intelligence request from Cloudforce One.

## Definition

```yaml
{"operationId": "cloudforce-one-priority-delete", "summary": "Delete a Priority Intelligence Requirement", "description": "Deletes a priority intelligence request from Cloudforce One.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cloudforce-one-requests_identifier"}}, {"name": "priority_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cloudforce-one-requests_uuid"}}], "responses": {"200": {"description": "Delete priority response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cloudforce-one-requests_api-response-common"}}}}, "4XX": {"description": "Delete priority response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cloudforce-one-requests_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Priority Intelligence Requirements (PIR)"], "x-api-token-group": ["Cloudforce One Write"]}
```
