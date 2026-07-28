---
title: Get Redundancy Group Details
page_id: operation-get-accounts-account-id-magic-redundancy-groups-redundancy-group-id-386f3c4f
path: operations/magic-redundancy-groups
description: Gets details for a specific redundancy group, including full member tunnel data.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/magic/redundancy_groups/{redundancy_group_id}
operation_ids:
    - magic-redundancy-groups-get-redundancy-group
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Redundancy Group Details

`GET /accounts/{account_id}/magic/redundancy_groups/{redundancy_group_id}`

Operation ID: `magic-redundancy-groups-get-redundancy-group`

Gets details for a specific redundancy group, including full member tunnel data.

## Definition

```yaml
{"operationId": "magic-redundancy-groups-get-redundancy-group", "summary": "Get Redundancy Group Details", "description": "Gets details for a specific redundancy group, including full member tunnel data.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}, {"name": "redundancy_group_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}], "responses": {"200": {"description": "Get Redundancy Group response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_redundancy_group_with_members_response"}}}}, "4XX": {"description": "Get Redundancy Group response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/magic_redundancy_group_with_members_response"}, {"$ref": "#/components/schemas/magic_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Magic Redundancy Groups"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}}
```
