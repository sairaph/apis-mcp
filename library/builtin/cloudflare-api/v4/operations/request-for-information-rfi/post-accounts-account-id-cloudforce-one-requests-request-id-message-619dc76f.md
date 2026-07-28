---
title: List Request Messages
page_id: operation-post-accounts-account-id-cloudforce-one-requests-request-id-message-5980a00d
path: operations/request-for-information-rfi
description: Lists messages in a Cloudforce One intelligence request conversation.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/requests/{request_id}/message
operation_ids:
    - cloudforce-one-request-message-list
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Request Messages

`POST /accounts/{account_id}/cloudforce-one/requests/{request_id}/message`

Operation ID: `cloudforce-one-request-message-list`

Lists messages in a Cloudforce One intelligence request conversation.

## Definition

```yaml
{"operationId": "cloudforce-one-request-message-list", "summary": "List Request Messages", "description": "Lists messages in a Cloudforce One intelligence request conversation.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cloudforce-one-requests_identifier"}}, {"name": "request_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cloudforce-one-requests_uuid"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cloudforce-one-requests_request-message-list"}}}}, "responses": {"200": {"description": "List request messages response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/cloudforce-one-requests_api-response-common"}, {"properties": {"result": {"type": "array", "items": {"$ref": "#/components/schemas/cloudforce-one-requests_request-message-item"}}}, "type": "object"}]}}}}, "4XX": {"description": "List request messages response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cloudforce-one-requests_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Request for Information (RFI)"], "x-api-token-group": ["Cloudforce One Write"]}
```
