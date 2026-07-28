---
title: Get Request Priority, Status, and TLP constants
page_id: operation-get-accounts-account-id-cloudforce-one-requests-constants-d7cb72c5
path: operations/request-for-information-rfi
description: Retrieves constant values used in Cloudforce One requests, including valid statuses and types.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/requests/constants
operation_ids:
    - cloudforce-one-request-constants
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Request Priority, Status, and TLP constants

`GET /accounts/{account_id}/cloudforce-one/requests/constants`

Operation ID: `cloudforce-one-request-constants`

Retrieves constant values used in Cloudforce One requests, including valid statuses and types.

## Definition

```yaml
{"operationId": "cloudforce-one-request-constants", "summary": "Get Request Priority, Status, and TLP constants", "description": "Retrieves constant values used in Cloudforce One requests, including valid statuses and types.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cloudforce-one-requests_identifier"}}], "responses": {"200": {"description": "Get request constants response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/cloudforce-one-requests_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/cloudforce-one-requests_request-constants"}}, "type": "object"}]}}}}, "4XX": {"description": "Get request constants response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cloudforce-one-requests_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Request for Information (RFI)"], "x-api-token-group": ["Cloudforce One Write", "Cloudforce One Read"]}
```
