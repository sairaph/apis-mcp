---
title: Update a Request Asset
page_id: operation-put-accounts-account-id-cloudforce-one-requests-request-id-asset-asset-i-8f8c6bc3
path: operations/request-for-information-rfi
description: Updates an asset in a Cloudforce One intelligence request.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/requests/{request_id}/asset/{asset_id}
operation_ids:
    - cloudforce-one-request-asset-update
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update a Request Asset

`PUT /accounts/{account_id}/cloudforce-one/requests/{request_id}/asset/{asset_id}`

Operation ID: `cloudforce-one-request-asset-update`

Updates an asset in a Cloudforce One intelligence request.

## Definition

```yaml
{"operationId": "cloudforce-one-request-asset-update", "summary": "Update a Request Asset", "description": "Updates an asset in a Cloudforce One intelligence request.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cloudforce-one-requests_identifier"}}, {"name": "request_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cloudforce-one-requests_uuid"}}, {"name": "asset_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cloudforce-one-requests_uuid"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cloudforce-one-requests_request-asset-edit"}}}}, "responses": {"200": {"description": "Update request asset response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/cloudforce-one-requests_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/cloudforce-one-requests_request-asset-item"}}, "type": "object"}]}}}}, "4XX": {"description": "Update request asset response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cloudforce-one-requests_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Request for Information (RFI)"], "x-api-token-group": ["Cloudforce One Write"]}
```
