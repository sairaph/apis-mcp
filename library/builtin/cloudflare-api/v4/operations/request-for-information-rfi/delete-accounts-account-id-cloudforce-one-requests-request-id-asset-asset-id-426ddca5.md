---
title: Delete a Request Asset
page_id: operation-delete-accounts-account-id-cloudforce-one-requests-request-id-asset-asse-d9c208bd
path: operations/request-for-information-rfi
description: Removes an asset from a Cloudforce One intelligence request.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/requests/{request_id}/asset/{asset_id}
operation_ids:
    - cloudforce-one-request-asset-delete
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete a Request Asset

`DELETE /accounts/{account_id}/cloudforce-one/requests/{request_id}/asset/{asset_id}`

Operation ID: `cloudforce-one-request-asset-delete`

Removes an asset from a Cloudforce One intelligence request.

## Definition

```yaml
{"operationId": "cloudforce-one-request-asset-delete", "summary": "Delete a Request Asset", "description": "Removes an asset from a Cloudforce One intelligence request.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cloudforce-one-requests_identifier"}}, {"name": "request_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cloudforce-one-requests_uuid"}}, {"name": "asset_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cloudforce-one-requests_uuid"}}], "responses": {"200": {"description": "Delete request asset response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cloudforce-one-requests_api-response-common"}}}}, "4XX": {"description": "Delete request asset response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cloudforce-one-requests_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Request for Information (RFI)"], "x-api-token-group": ["Cloudforce One Write"]}
```
