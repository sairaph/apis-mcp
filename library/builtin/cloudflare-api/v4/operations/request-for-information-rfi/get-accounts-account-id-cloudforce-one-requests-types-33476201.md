---
title: Get Request Types
page_id: operation-get-accounts-account-id-cloudforce-one-requests-types-68f9158b
path: operations/request-for-information-rfi
description: Lists available request types for Cloudforce One intelligence requests.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/requests/types
operation_ids:
    - cloudforce-one-request-types
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Request Types

`GET /accounts/{account_id}/cloudforce-one/requests/types`

Operation ID: `cloudforce-one-request-types`

Lists available request types for Cloudforce One intelligence requests.

## Definition

```yaml
{"operationId": "cloudforce-one-request-types", "summary": "Get Request Types", "description": "Lists available request types for Cloudforce One intelligence requests.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cloudforce-one-requests_identifier"}}], "responses": {"200": {"description": "Get request types response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/cloudforce-one-requests_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/cloudforce-one-requests_request-types"}}, "type": "object"}]}}}}, "4XX": {"description": "Get request types response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cloudforce-one-requests_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Request for Information (RFI)"], "x-api-token-group": ["Cloudforce One Write", "Cloudforce One Read"]}
```
