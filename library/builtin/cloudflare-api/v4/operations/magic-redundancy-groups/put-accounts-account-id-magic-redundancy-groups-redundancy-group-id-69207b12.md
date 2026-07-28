---
title: Update a Redundancy Group
page_id: operation-put-accounts-account-id-magic-redundancy-groups-redundancy-group-id-8666c242
path: operations/magic-redundancy-groups
description: Replaces the name, description, and full set of members for an existing redundancy group.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/magic/redundancy_groups/{redundancy_group_id}
operation_ids:
    - magic-redundancy-groups-update-redundancy-group
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update a Redundancy Group

`PUT /accounts/{account_id}/magic/redundancy_groups/{redundancy_group_id}`

Operation ID: `magic-redundancy-groups-update-redundancy-group`

Replaces the name, description, and full set of members for an existing redundancy group.

## Definition

```yaml
{"operationId": "magic-redundancy-groups-update-redundancy-group", "summary": "Update a Redundancy Group", "description": "Replaces the name, description, and full set of members for an existing redundancy group.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}, {"name": "redundancy_group_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_create_redundancy_group_request"}}}}, "responses": {"200": {"description": "Update Redundancy Group response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_redundancy_group_single_response"}}}}, "4XX": {"description": "Update Redundancy Group response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/magic_redundancy_group_single_response"}, {"$ref": "#/components/schemas/magic_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Magic Redundancy Groups"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}}
```
