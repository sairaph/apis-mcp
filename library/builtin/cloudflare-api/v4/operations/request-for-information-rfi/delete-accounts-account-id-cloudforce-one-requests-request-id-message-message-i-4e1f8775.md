---
title: Delete a Request Message
page_id: operation-delete-accounts-account-id-cloudforce-one-requests-request-id-message-me-a9f232c2
path: operations/request-for-information-rfi
description: Removes a message from a Cloudforce One intelligence request thread.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/requests/{request_id}/message/{message_id}
operation_ids:
    - cloudforce-one-request-message-delete
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete a Request Message

`DELETE /accounts/{account_id}/cloudforce-one/requests/{request_id}/message/{message_id}`

Operation ID: `cloudforce-one-request-message-delete`

Removes a message from a Cloudforce One intelligence request thread.

## Definition

```yaml
{"operationId": "cloudforce-one-request-message-delete", "summary": "Delete a Request Message", "description": "Removes a message from a Cloudforce One intelligence request thread.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cloudforce-one-requests_identifier"}}, {"name": "request_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cloudforce-one-requests_uuid"}}, {"name": "message_id", "in": "path", "required": true, "schema": {"type": "integer"}}], "responses": {"200": {"description": "Delete request message response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cloudforce-one-requests_api-response-common"}}}}, "4XX": {"description": "Delete request message response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cloudforce-one-requests_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Request for Information (RFI)"], "x-api-token-group": ["Cloudforce One Write"]}
```
