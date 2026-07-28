---
title: Create a New Request Message
page_id: operation-post-accounts-account-id-cloudforce-one-requests-request-id-message-new-f5bb299b
path: operations/request-for-information-rfi
description: Adds a message to a Cloudforce One intelligence request conversation.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/requests/{request_id}/message/new
operation_ids:
    - cloudforce-one-request-message-new
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a New Request Message

`POST /accounts/{account_id}/cloudforce-one/requests/{request_id}/message/new`

Operation ID: `cloudforce-one-request-message-new`

Adds a message to a Cloudforce One intelligence request conversation.

## Definition

```yaml
{"operationId": "cloudforce-one-request-message-new", "summary": "Create a New Request Message", "description": "Adds a message to a Cloudforce One intelligence request conversation.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cloudforce-one-requests_identifier"}}, {"name": "request_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cloudforce-one-requests_uuid"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cloudforce-one-requests_request-message-edit"}}}}, "responses": {"200": {"description": "Create request message response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/cloudforce-one-requests_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/cloudforce-one-requests_request-message-item"}}, "type": "object"}]}}}}, "4XX": {"description": "Create request message response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cloudforce-one-requests_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Request for Information (RFI)"], "x-api-token-group": ["Cloudforce One Write"]}
```
