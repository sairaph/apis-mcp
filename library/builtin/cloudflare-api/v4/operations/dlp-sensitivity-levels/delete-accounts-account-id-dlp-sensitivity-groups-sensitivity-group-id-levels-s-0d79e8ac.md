---
title: Delete a single sensitivity level.
page_id: operation-delete-accounts-account-id-dlp-sensitivity-groups-sensitivity-group-id-l-c4c91aca
path: operations/dlp-sensitivity-levels
description: Deletes a sensitivity level from a group.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/dlp/sensitivity_groups/{sensitivity_group_id}/levels/{sensitivity_level_id}
operation_ids:
    - dlp-sensitivity-levels-delete
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete a single sensitivity level.

`DELETE /accounts/{account_id}/dlp/sensitivity_groups/{sensitivity_group_id}/levels/{sensitivity_level_id}`

Operation ID: `dlp-sensitivity-levels-delete`

Deletes a sensitivity level from a group.

## Definition

```yaml
{"operationId": "dlp-sensitivity-levels-delete", "summary": "Delete a single sensitivity level.", "description": "Deletes a sensitivity level from a group.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "sensitivity_group_id", "in": "path", "required": true, "schema": {"type": "string", "format": "uuid"}}, {"name": "sensitivity_level_id", "in": "path", "required": true, "schema": {"type": "string", "format": "uuid"}}], "responses": {"200": {"description": "Sensitivity level delete was successful.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dlp_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/dlp_Empty"}}, "type": "object"}]}}}}, "4XX": {"description": "Sensitivity level delete failed.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["DLP Sensitivity Levels"], "x-api-token-group": ["Zero Trust Write"]}
```
