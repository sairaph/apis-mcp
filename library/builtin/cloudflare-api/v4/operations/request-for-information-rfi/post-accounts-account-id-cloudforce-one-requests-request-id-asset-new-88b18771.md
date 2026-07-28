---
title: Create a New Request Asset
page_id: operation-post-accounts-account-id-cloudforce-one-requests-request-id-asset-new-bb87968a
path: operations/request-for-information-rfi
description: Uploads a new asset to a Cloudforce One intelligence request.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/requests/{request_id}/asset/new
operation_ids:
    - cloudforce-one-request-asset-new
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a New Request Asset

`POST /accounts/{account_id}/cloudforce-one/requests/{request_id}/asset/new`

Operation ID: `cloudforce-one-request-asset-new`

Uploads a new asset to a Cloudforce One intelligence request.

## Definition

```yaml
{"operationId": "cloudforce-one-request-asset-new", "summary": "Create a New Request Asset", "description": "Uploads a new asset to a Cloudforce One intelligence request.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cloudforce-one-requests_identifier"}}, {"name": "request_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cloudforce-one-requests_uuid"}}], "requestBody": {"required": true, "content": {"multipart/form-data": {"schema": {"$ref": "#/components/schemas/cloudforce-one-requests_request-asset-edit"}}}}, "responses": {"200": {"description": "Create request asset response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/cloudforce-one-requests_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/cloudforce-one-requests_request-asset-item"}}, "type": "object"}]}}}}, "4XX": {"description": "Create request asset response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cloudforce-one-requests_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Request for Information (RFI)"], "x-api-token-group": ["Cloudforce One Write"]}
```
