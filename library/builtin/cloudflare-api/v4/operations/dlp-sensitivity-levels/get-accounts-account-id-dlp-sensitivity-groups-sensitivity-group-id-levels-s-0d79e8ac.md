---
title: Retrieve a specific sensitivity level.
page_id: operation-get-accounts-account-id-dlp-sensitivity-groups-sensitivity-group-id-leve-fc603ab2
path: operations/dlp-sensitivity-levels
description: Gets a sensitivity level from a group.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/dlp/sensitivity_groups/{sensitivity_group_id}/levels/{sensitivity_level_id}
operation_ids:
    - dlp-sensitivity-levels-read
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Retrieve a specific sensitivity level.

`GET /accounts/{account_id}/dlp/sensitivity_groups/{sensitivity_group_id}/levels/{sensitivity_level_id}`

Operation ID: `dlp-sensitivity-levels-read`

Gets a sensitivity level from a group.

## Definition

```yaml
{"operationId": "dlp-sensitivity-levels-read", "summary": "Retrieve a specific sensitivity level.", "description": "Gets a sensitivity level from a group.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "sensitivity_group_id", "in": "path", "required": true, "schema": {"type": "string", "format": "uuid"}}, {"name": "sensitivity_level_id", "in": "path", "required": true, "schema": {"type": "string", "format": "uuid"}}], "responses": {"200": {"description": "Sensitivity level read was successful.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dlp_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/dlp_SensitivityLevel"}}, "type": "object"}]}}}}, "4XX": {"description": "Sensitivity level read failed.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["DLP Sensitivity Levels"], "x-api-token-group": ["Zero Trust Read", "Zero Trust Write"]}
```
