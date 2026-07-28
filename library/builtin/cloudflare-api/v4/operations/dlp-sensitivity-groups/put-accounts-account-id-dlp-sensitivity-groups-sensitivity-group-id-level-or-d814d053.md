---
title: Set the ordering of levels within a sensitivity group.
page_id: operation-put-accounts-account-id-dlp-sensitivity-groups-sensitivity-group-id-leve-981bf7d6
path: operations/dlp-sensitivity-groups
description: Updates the order of sensitivity levels in a group.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/dlp/sensitivity_groups/{sensitivity_group_id}/level_order
operation_ids:
    - dlp-sensitivity-groups-put-level-order
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Set the ordering of levels within a sensitivity group.

`PUT /accounts/{account_id}/dlp/sensitivity_groups/{sensitivity_group_id}/level_order`

Operation ID: `dlp-sensitivity-groups-put-level-order`

Updates the order of sensitivity levels in a group.

## Definition

```yaml
{"operationId": "dlp-sensitivity-groups-put-level-order", "summary": "Set the ordering of levels within a sensitivity group.", "description": "Updates the order of sensitivity levels in a group.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "sensitivity_group_id", "in": "path", "required": true, "schema": {"type": "string", "format": "uuid"}}], "requestBody": {"description": "Desired ordered list of level IDs.", "required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_SensitivityLevelOrder"}}}}, "responses": {"200": {"description": "Level order update was successful.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dlp_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/dlp_SensitivityLevelOrder"}}, "type": "object"}]}}}}, "4XX": {"description": "Level order update failed.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["DLP Sensitivity Groups"], "x-api-token-group": ["Zero Trust Write"]}
```
