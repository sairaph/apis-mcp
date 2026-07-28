---
title: Update the attributes of a single sensitivity level.
page_id: operation-put-accounts-account-id-dlp-sensitivity-groups-sensitivity-group-id-leve-04003216
path: operations/dlp-sensitivity-levels
description: Updates a sensitivity level in a group.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/dlp/sensitivity_groups/{sensitivity_group_id}/levels/{sensitivity_level_id}
operation_ids:
    - dlp-sensitivity-levels-update
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update the attributes of a single sensitivity level.

`PUT /accounts/{account_id}/dlp/sensitivity_groups/{sensitivity_group_id}/levels/{sensitivity_level_id}`

Operation ID: `dlp-sensitivity-levels-update`

Updates a sensitivity level in a group.

## Definition

```yaml
{"operationId": "dlp-sensitivity-levels-update", "summary": "Update the attributes of a single sensitivity level.", "description": "Updates a sensitivity level in a group.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "sensitivity_group_id", "in": "path", "required": true, "schema": {"type": "string", "format": "uuid"}}, {"name": "sensitivity_level_id", "in": "path", "required": true, "schema": {"type": "string", "format": "uuid"}}], "requestBody": {"description": "Attributes of the sensitivity level to update.", "required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_SensitivityLevelUpdate"}}}}, "responses": {"200": {"description": "Sensitivity level update was successful.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dlp_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/dlp_SensitivityLevel"}}, "type": "object"}]}}}}, "4XX": {"description": "Sensitivity level update failed.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["DLP Sensitivity Levels"], "x-api-token-group": ["Zero Trust Write"]}
```
