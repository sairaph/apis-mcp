---
title: Delete a Redundancy Group
page_id: operation-delete-accounts-account-id-magic-redundancy-groups-redundancy-group-id-a20f2037
path: operations/magic-redundancy-groups
description: Deletes a redundancy group. Member tunnels are not deleted — their redundancy_group_id is cleared.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - DELETE
api_endpoints:
    - /accounts/{account_id}/magic/redundancy_groups/{redundancy_group_id}
operation_ids:
    - magic-redundancy-groups-delete-redundancy-group
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Delete a Redundancy Group

`DELETE /accounts/{account_id}/magic/redundancy_groups/{redundancy_group_id}`

Operation ID: `magic-redundancy-groups-delete-redundancy-group`

Deletes a redundancy group. Member tunnels are not deleted — their redundancy_group_id is cleared.

## Definition

```yaml
{"operationId": "magic-redundancy-groups-delete-redundancy-group", "summary": "Delete a Redundancy Group", "description": "Deletes a redundancy group. Member tunnels are not deleted — their redundancy_group_id is cleared.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}, {"name": "redundancy_group_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}], "responses": {"200": {"description": "Delete Redundancy Group response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_delete_redundancy_group_response"}}}}, "4XX": {"description": "Delete Redundancy Group response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/magic_delete_redundancy_group_response"}, {"$ref": "#/components/schemas/magic_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Magic Redundancy Groups"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}}
```
