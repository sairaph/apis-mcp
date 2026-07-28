---
title: Creates a new sensitivity level.
page_id: operation-post-accounts-account-id-dlp-sensitivity-groups-sensitivity-group-id-lev-366248f9
path: operations/dlp-sensitivity-levels
description: Creates a sensitivity level in a group.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/dlp/sensitivity_groups/{sensitivity_group_id}/levels
operation_ids:
    - dlp-sensitivity-levels-create
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Creates a new sensitivity level.

`POST /accounts/{account_id}/dlp/sensitivity_groups/{sensitivity_group_id}/levels`

Operation ID: `dlp-sensitivity-levels-create`

Creates a sensitivity level in a group.

## Definition

```yaml
{"operationId": "dlp-sensitivity-levels-create", "summary": "Creates a new sensitivity level.", "description": "Creates a sensitivity level in a group.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "sensitivity_group_id", "in": "path", "required": true, "schema": {"type": "string", "format": "uuid"}}], "requestBody": {"description": "Attributes of the new sensitivity level.", "required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_NewSensitivityLevel"}}}}, "responses": {"200": {"description": "Sensitivity level created successfully.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dlp_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/dlp_SensitivityLevel"}}, "type": "object"}]}}}}, "4XX": {"description": "Sensitivity level creation failed.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["DLP Sensitivity Levels"], "x-api-token-group": ["Zero Trust Write"]}
```
