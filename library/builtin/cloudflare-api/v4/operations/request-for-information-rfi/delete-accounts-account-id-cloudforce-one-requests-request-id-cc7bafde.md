---
title: Delete a Request
page_id: operation-delete-accounts-account-id-cloudforce-one-requests-request-id-01cd1050
path: operations/request-for-information-rfi
description: Deletes a Cloudforce One intelligence request and all associated data.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/requests/{request_id}
operation_ids:
    - cloudforce-one-request-delete
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete a Request

`DELETE /accounts/{account_id}/cloudforce-one/requests/{request_id}`

Operation ID: `cloudforce-one-request-delete`

Deletes a Cloudforce One intelligence request and all associated data.

## Definition

```yaml
{"operationId": "cloudforce-one-request-delete", "summary": "Delete a Request", "description": "Deletes a Cloudforce One intelligence request and all associated data.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cloudforce-one-requests_identifier"}}, {"name": "request_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cloudforce-one-requests_uuid"}}], "responses": {"200": {"description": "Delete request response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cloudforce-one-requests_api-response-common"}}}}, "4XX": {"description": "Delete request response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cloudforce-one-requests_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Request for Information (RFI)"], "x-api-token-group": ["Cloudforce One Write"]}
```
