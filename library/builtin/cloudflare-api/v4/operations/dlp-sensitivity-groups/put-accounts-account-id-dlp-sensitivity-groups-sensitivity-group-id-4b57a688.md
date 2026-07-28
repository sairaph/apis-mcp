---
title: Update the attributes of a single sensitivity group.
page_id: operation-put-accounts-account-id-dlp-sensitivity-groups-sensitivity-group-id-30b33e89
path: operations/dlp-sensitivity-groups
description: Updates a sensitivity group and its levels.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/dlp/sensitivity_groups/{sensitivity_group_id}
operation_ids:
    - dlp-sensitivity-groups-update
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update the attributes of a single sensitivity group.

`PUT /accounts/{account_id}/dlp/sensitivity_groups/{sensitivity_group_id}`

Operation ID: `dlp-sensitivity-groups-update`

Updates a sensitivity group and its levels.

## Definition

```yaml
{"operationId": "dlp-sensitivity-groups-update", "summary": "Update the attributes of a single sensitivity group.", "description": "Updates a sensitivity group and its levels.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}, {"name": "sensitivity_group_id", "in": "path", "required": true, "schema": {"type": "string", "format": "uuid"}}], "requestBody": {"description": "Attributes of the sensitivity group to update.", "required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_SensitivityGroupUpdate"}}}}, "responses": {"200": {"description": "Sensitivity group update was successful.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/dlp_api-response-single"}, {"properties": {"result": {"$ref": "#/components/schemas/dlp_SensitivityGroup"}}, "type": "object"}]}}}}, "4XX": {"description": "Sensitivity group update failed.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/dlp_api-response-common-failure"}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["DLP Sensitivity Groups"], "x-api-token-group": ["Zero Trust Write"]}
```
