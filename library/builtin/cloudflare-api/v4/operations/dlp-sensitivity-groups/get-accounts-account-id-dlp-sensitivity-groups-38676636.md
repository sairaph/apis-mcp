---
title: Retrieve all sensitivity groups in an account
page_id: operation-get-accounts-account-id-dlp-sensitivity-groups-6d339058
path: operations/dlp-sensitivity-groups
description: Lists sensitivity groups configured for the account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/dlp/sensitivity_groups
operation_ids:
    - dlp-sensitivity-groups-list
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Retrieve all sensitivity groups in an account

`GET /accounts/{account_id}/dlp/sensitivity_groups`

Operation ID: `dlp-sensitivity-groups-list`

Lists sensitivity groups configured for the account.

## Definition

```yaml
{"operationId": "dlp-sensitivity-groups-list", "summary": "Retrieve all sensitivity groups in an account", "description": "Lists sensitivity groups configured for the account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"description": "Sensitivity group read was successful.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dlp_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/dlp_SensitivityGroupArray"}}, "type": "object"}]}}}}, "4XX": {"description": "Sensitivity group read failed.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["DLP Sensitivity Groups"], "x-api-token-group": ["Zero Trust Read", "Zero Trust Write"]}
```
