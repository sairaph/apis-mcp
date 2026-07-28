---
title: Get Request Quota
page_id: operation-get-accounts-account-id-cloudforce-one-requests-quota-879e4e75
path: operations/request-for-information-rfi
description: Retrieves quota usage for Cloudforce One standard requests.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/requests/quota
operation_ids:
    - cloudforce-one-request-quota
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Request Quota

`GET /accounts/{account_id}/cloudforce-one/requests/quota`

Operation ID: `cloudforce-one-request-quota`

Retrieves quota usage for Cloudforce One standard requests.

## Definition

```yaml
{"operationId": "cloudforce-one-request-quota", "summary": "Get Request Quota", "description": "Retrieves quota usage for Cloudforce One standard requests.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cloudforce-one-requests_identifier"}}], "responses": {"200": {"description": "Get request quota response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/cloudforce-one-requests_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/cloudforce-one-requests_quota"}}, "type": "object"}]}}}}, "4XX": {"description": "Get request quota response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cloudforce-one-requests_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Request for Information (RFI)"], "x-api-token-group": ["Cloudforce One Write", "Cloudforce One Read"]}
```
