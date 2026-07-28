---
title: Retrieve all sensitivity levels in a sensitivity group
page_id: operation-get-accounts-account-id-dlp-sensitivity-groups-sensitivity-group-id-leve-f136ec44
path: operations/dlp-sensitivity-levels
description: Lists sensitivity levels in a sensitivity group.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/dlp/sensitivity_groups/{sensitivity_group_id}/levels
operation_ids:
    - dlp-sensitivity-levels-list
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Retrieve all sensitivity levels in a sensitivity group

`GET /accounts/{account_id}/dlp/sensitivity_groups/{sensitivity_group_id}/levels`

Operation ID: `dlp-sensitivity-levels-list`

Lists sensitivity levels in a sensitivity group.

## Definition

```yaml
{"operationId": "dlp-sensitivity-levels-list", "summary": "Retrieve all sensitivity levels in a sensitivity group", "description": "Lists sensitivity levels in a sensitivity group.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "sensitivity_group_id", "in": "path", "required": true, "schema": {"type": "string", "format": "uuid"}}], "responses": {"200": {"description": "Sensitivity levels read was successful.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dlp_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/dlp_SensitivityLevelArray"}}, "type": "object"}]}}}}, "4XX": {"description": "Sensitivity levels read failed.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["DLP Sensitivity Levels"], "x-api-token-group": ["Zero Trust Read", "Zero Trust Write"]}
```
