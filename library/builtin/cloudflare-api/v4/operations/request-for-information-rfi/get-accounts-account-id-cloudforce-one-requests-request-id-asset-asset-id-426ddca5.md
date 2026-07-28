---
title: Get a Request Asset
page_id: operation-get-accounts-account-id-cloudforce-one-requests-request-id-asset-asset-i-03b8d8c5
path: operations/request-for-information-rfi
description: Retrieves an asset attached to a Cloudforce One intelligence request.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/requests/{request_id}/asset/{asset_id}
operation_ids:
    - cloudforce-one-request-asset-get
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get a Request Asset

`GET /accounts/{account_id}/cloudforce-one/requests/{request_id}/asset/{asset_id}`

Operation ID: `cloudforce-one-request-asset-get`

Retrieves an asset attached to a Cloudforce One intelligence request.

## Definition

```yaml
{"operationId": "cloudforce-one-request-asset-get", "summary": "Get a Request Asset", "description": "Retrieves an asset attached to a Cloudforce One intelligence request.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cloudforce-one-requests_identifier"}}, {"name": "request_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cloudforce-one-requests_uuid"}}, {"name": "asset_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cloudforce-one-requests_uuid"}}], "responses": {"200": {"description": "Get request asset response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/cloudforce-one-requests_api-response-common"}, {"properties": {"result": {"type": "array", "items": {"$ref": "#/components/schemas/cloudforce-one-requests_request-asset-item"}}}, "type": "object"}]}}}}, "4XX": {"description": "Get request asset response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cloudforce-one-requests_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Request for Information (RFI)"], "x-api-token-group": ["Cloudforce One Write", "Cloudforce One Read"]}
```
