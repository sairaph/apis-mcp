---
title: Creates a new sensitivity group.
page_id: operation-post-accounts-account-id-dlp-sensitivity-groups-061278a3
path: operations/dlp-sensitivity-groups
description: Creates a sensitivity group, optionally from a template.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/dlp/sensitivity_groups
operation_ids:
    - dlp-sensitivity-groups-create
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Creates a new sensitivity group.

`POST /accounts/{account_id}/dlp/sensitivity_groups`

Operation ID: `dlp-sensitivity-groups-create`

Creates a sensitivity group, optionally from a template.

## Definition

```yaml
{"operationId": "dlp-sensitivity-groups-create", "summary": "Creates a new sensitivity group.", "description": "Creates a sensitivity group, optionally from a template.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}], "requestBody": {"description": "Attributes of the new sensitivity group.", "required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_NewSensitivityGroup"}}}}, "responses": {"200": {"description": "Sensitivity group created successfully.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dlp_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/dlp_SensitivityGroup"}}, "type": "object"}]}}}}, "4XX": {"description": "Sensitivity group creation failed.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["DLP Sensitivity Groups"], "x-api-token-group": ["Zero Trust Write"]}
```
