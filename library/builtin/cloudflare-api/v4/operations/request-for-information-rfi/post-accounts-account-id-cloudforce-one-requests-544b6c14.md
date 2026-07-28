---
title: List Requests
page_id: operation-post-accounts-account-id-cloudforce-one-requests-2320955e
path: operations/request-for-information-rfi
description: Lists Cloudforce One intelligence requests with filtering and pagination.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/requests
operation_ids:
    - cloudforce-one-request-list
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Requests

`POST /accounts/{account_id}/cloudforce-one/requests`

Operation ID: `cloudforce-one-request-list`

Lists Cloudforce One intelligence requests with filtering and pagination.

## Definition

```yaml
{"operationId": "cloudforce-one-request-list", "summary": "List Requests", "description": "Lists Cloudforce One intelligence requests with filtering and pagination.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cloudforce-one-requests_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cloudforce-one-requests_request-list"}}}}, "responses": {"200": {"description": "List requests response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/cloudforce-one-requests_api-response-common"}, {"properties": {"result": {"type": "array", "items": {"$ref": "#/components/schemas/cloudforce-one-requests_request-list-item"}}}, "type": "object"}]}}}}, "4XX": {"description": "Create response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cloudforce-one-requests_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Request for Information (RFI)"], "x-api-token-group": ["Cloudforce One Write"]}
```
