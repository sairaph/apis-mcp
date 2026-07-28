---
title: List Request Assets
page_id: operation-post-accounts-account-id-cloudforce-one-requests-request-id-asset-1316aa7b
path: operations/request-for-information-rfi
description: Lists assets attached to a Cloudforce One intelligence request.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/requests/{request_id}/asset
operation_ids:
    - cloudforce-one-request-asset-list
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Request Assets

`POST /accounts/{account_id}/cloudforce-one/requests/{request_id}/asset`

Operation ID: `cloudforce-one-request-asset-list`

Lists assets attached to a Cloudforce One intelligence request.

## Definition

```yaml
{"operationId": "cloudforce-one-request-asset-list", "summary": "List Request Assets", "description": "Lists assets attached to a Cloudforce One intelligence request.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cloudforce-one-requests_identifier"}}, {"name": "request_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cloudforce-one-requests_uuid"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cloudforce-one-requests_request-asset-list"}}}}, "responses": {"200": {"description": "List request assets response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/cloudforce-one-requests_api-response-common"}, {"properties": {"result": {"type": "array", "items": {"$ref": "#/components/schemas/cloudforce-one-requests_request-asset-item"}}}, "type": "object"}]}}}}, "4XX": {"description": "List request assets response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cloudforce-one-requests_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Request for Information (RFI)"], "x-api-token-group": ["Cloudforce One Write"]}
```
