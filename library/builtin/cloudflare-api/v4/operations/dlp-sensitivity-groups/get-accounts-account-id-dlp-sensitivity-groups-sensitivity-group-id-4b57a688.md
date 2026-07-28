---
title: Retrieve a specific sensitivity group.
page_id: operation-get-accounts-account-id-dlp-sensitivity-groups-sensitivity-group-id-23e6f36b
path: operations/dlp-sensitivity-groups
description: Gets a sensitivity group and its levels.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/dlp/sensitivity_groups/{sensitivity_group_id}
operation_ids:
    - dlp-sensitivity-groups-read
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Retrieve a specific sensitivity group.

`GET /accounts/{account_id}/dlp/sensitivity_groups/{sensitivity_group_id}`

Operation ID: `dlp-sensitivity-groups-read`

Gets a sensitivity group and its levels.

## Definition

```yaml
{"operationId": "dlp-sensitivity-groups-read", "summary": "Retrieve a specific sensitivity group.", "description": "Gets a sensitivity group and its levels.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "sensitivity_group_id", "in": "path", "required": true, "schema": {"type": "string", "format": "uuid"}}], "responses": {"200": {"description": "Sensitivity group read was successful.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dlp_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/dlp_SensitivityGroup"}}, "type": "object"}]}}}}, "4XX": {"description": "Sensitivity group read failed.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["DLP Sensitivity Groups"], "x-api-token-group": ["Zero Trust Read", "Zero Trust Write"]}
```
