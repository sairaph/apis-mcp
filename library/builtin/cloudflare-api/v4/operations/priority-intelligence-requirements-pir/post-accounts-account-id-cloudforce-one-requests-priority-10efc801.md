---
title: List Priority Intelligence Requirements
page_id: operation-post-accounts-account-id-cloudforce-one-requests-priority-9c474035
path: operations/priority-intelligence-requirements-pir
description: Lists priority intelligence requests in Cloudforce One.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/requests/priority
operation_ids:
    - cloudforce-one-priority-list
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Priority Intelligence Requirements

`POST /accounts/{account_id}/cloudforce-one/requests/priority`

Operation ID: `cloudforce-one-priority-list`

Lists priority intelligence requests in Cloudforce One.

## Definition

```yaml
{"operationId": "cloudforce-one-priority-list", "summary": "List Priority Intelligence Requirements", "description": "Lists priority intelligence requests in Cloudforce One.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cloudforce-one-requests_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cloudforce-one-requests_priority-list"}}}}, "responses": {"200": {"description": "List priorities response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/cloudforce-one-requests_api-response-common"}, {"properties": {"result": {"type": "array", "items": {"$ref": "#/components/schemas/cloudforce-one-requests_priority-item"}}}, "type": "object"}]}}}}, "4XX": {"description": "List priorities response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cloudforce-one-requests_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Priority Intelligence Requirements (PIR)"], "x-api-token-group": ["Cloudforce One Write"]}
```
