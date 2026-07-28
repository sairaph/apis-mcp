---
title: List Redundancy Groups
page_id: operation-get-accounts-account-id-magic-redundancy-groups-85be64e5
path: operations/magic-redundancy-groups
description: Lists redundancy groups associated with an account, including full member tunnel data.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/magic/redundancy_groups
operation_ids:
    - magic-redundancy-groups-list-redundancy-groups
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Redundancy Groups

`GET /accounts/{account_id}/magic/redundancy_groups`

Operation ID: `magic-redundancy-groups-list-redundancy-groups`

Lists redundancy groups associated with an account, including full member tunnel data.

## Definition

```yaml
{"operationId": "magic-redundancy-groups-list-redundancy-groups", "summary": "List Redundancy Groups", "description": "Lists redundancy groups associated with an account, including full member tunnel data.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}], "responses": {"200": {"description": "List Redundancy Groups response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_redundancy_groups_collection_response"}}}}, "4XX": {"description": "List Redundancy Groups response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/magic_redundancy_groups_collection_response"}, {"$ref": "#/components/schemas/magic_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Magic Redundancy Groups"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}}
```
