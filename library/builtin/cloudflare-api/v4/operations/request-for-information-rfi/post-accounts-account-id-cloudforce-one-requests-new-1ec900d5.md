---
title: Create a New Request.
page_id: operation-post-accounts-account-id-cloudforce-one-requests-new-e828cb07
path: operations/request-for-information-rfi
description: Creating a request adds the request into the Cloudforce One queue for analysis. In addition to the content, a short title, type, priority, and releasability should be provided. If one is not provided, a default will be assigned.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/cloudforce-one/requests/new
operation_ids:
    - cloudforce-one-request-new
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a New Request.

`POST /accounts/{account_id}/cloudforce-one/requests/new`

Operation ID: `cloudforce-one-request-new`

Creating a request adds the request into the Cloudforce One queue for analysis. In addition to the content, a short title, type, priority, and releasability should be provided. If one is not provided, a default will be assigned.

## Definition

```yaml
{"operationId": "cloudforce-one-request-new", "summary": "Create a New Request.", "description": "Creating a request adds the request into the Cloudforce One queue for analysis. In addition to the content, a short title, type, priority, and releasability should be provided. If one is not provided, a default will be assigned.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/cloudforce-one-requests_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cloudforce-one-requests_request-edit"}}}}, "responses": {"200": {"description": "Create request response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/cloudforce-one-requests_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/cloudforce-one-requests_request-item"}}, "type": "object"}]}}}}, "4XX": {"description": "Create response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/cloudforce-one-requests_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Request for Information (RFI)"], "x-api-token-group": ["Cloudforce One Write"]}
```
