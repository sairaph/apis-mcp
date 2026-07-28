---
title: Get Priority Intelligence Requirement Quota
page_id: operation-get-accounts-account-id-cloudforce-one-requests-priority-quota-b9e3eab1
path: operations/priority-intelligence-requirements-pir
description: Retrieves quota usage for Cloudforce One priority requests.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/requests/priority/quota
operation_ids:
    - cloudforce-one-priority-quota
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Priority Intelligence Requirement Quota

`GET /accounts/{account_id}/cloudforce-one/requests/priority/quota`

Operation ID: `cloudforce-one-priority-quota`

Retrieves quota usage for Cloudforce One priority requests.

## Definition

```yaml
{"operationId": "cloudforce-one-priority-quota", "summary": "Get Priority Intelligence Requirement Quota", "description": "Retrieves quota usage for Cloudforce One priority requests.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cloudforce-one-requests_identifier"}}], "responses": {"200": {"description": "Get priority quota response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/cloudforce-one-requests_api-response-common"}, {"properties": {"result": {"allOf": [{"$ref": "#/components/schemas/cloudforce-one-requests_quota"}]}}, "type": "object"}]}}}}, "4XX": {"description": "Get priority quota response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cloudforce-one-requests_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Priority Intelligence Requirements (PIR)"], "x-api-token-group": ["Cloudforce One Write", "Cloudforce One Read"]}
```
