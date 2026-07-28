---
title: Update a Request Message
page_id: operation-put-accounts-account-id-cloudforce-one-requests-request-id-message-messa-15d444ec
path: operations/request-for-information-rfi
description: Updates a message in a Cloudforce One intelligence request thread.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/requests/{request_id}/message/{message_id}
operation_ids:
    - cloudforce-one-request-message-update
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update a Request Message

`PUT /accounts/{account_id}/cloudforce-one/requests/{request_id}/message/{message_id}`

Operation ID: `cloudforce-one-request-message-update`

Updates a message in a Cloudforce One intelligence request thread.

## Definition

```yaml
{"operationId": "cloudforce-one-request-message-update", "summary": "Update a Request Message", "description": "Updates a message in a Cloudforce One intelligence request thread.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cloudforce-one-requests_identifier"}}, {"name": "request_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cloudforce-one-requests_uuid"}}, {"name": "message_id", "in": "path", "required": true, "schema": {"type": "integer"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cloudforce-one-requests_request-message-edit"}}}}, "responses": {"200": {"description": "Update request message response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/cloudforce-one-requests_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/cloudforce-one-requests_request-message-item"}}, "type": "object"}]}}}}, "4XX": {"description": "Update request message response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cloudforce-one-requests_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Request for Information (RFI)"], "x-api-token-group": ["Cloudforce One Write"]}
```
