---
title: Retrieve all data classes in an account
page_id: operation-get-accounts-account-id-dlp-data-classes-5d051ee6
path: operations/dlp-data-classes
description: Lists data classes configured for the account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/dlp/data_classes
operation_ids:
    - dlp-data-classes-list
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Retrieve all data classes in an account

`GET /accounts/{account_id}/dlp/data_classes`

Operation ID: `dlp-data-classes-list`

Lists data classes configured for the account.

## Definition

```yaml
{"operationId": "dlp-data-classes-list", "summary": "Retrieve all data classes in an account", "description": "Lists data classes configured for the account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"description": "Data class list was successful.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dlp_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/dlp_DataClassArray"}}, "type": "object"}]}}}}, "4XX": {"description": "Data class list failed.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["DLP Data Classes"], "x-api-token-group": ["Zero Trust Read", "Zero Trust Write"]}
```
