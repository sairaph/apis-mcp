---
title: Create a Redundancy Group
page_id: operation-post-accounts-account-id-magic-redundancy-groups-2f175f87
path: operations/magic-redundancy-groups
description: Creates a new redundancy group, optionally with tunnel members.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/magic/redundancy_groups
operation_ids:
    - magic-redundancy-groups-create-redundancy-group
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a Redundancy Group

`POST /accounts/{account_id}/magic/redundancy_groups`

Operation ID: `magic-redundancy-groups-create-redundancy-group`

Creates a new redundancy group, optionally with tunnel members.

## Definition

```yaml
{"operationId": "magic-redundancy-groups-create-redundancy-group", "summary": "Create a Redundancy Group", "description": "Creates a new redundancy group, optionally with tunnel members.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/magic_identifier"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_create_redundancy_group_request"}}}}, "responses": {"201": {"description": "Create Redundancy Group response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/magic_redundancy_group_single_response"}}}}, "4XX": {"description": "Create Redundancy Group response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/magic_redundancy_group_single_response"}, {"$ref": "#/components/schemas/magic_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Magic Redundancy Groups"], "x-cfPlanAvailability": {"business": false, "enterprise": true, "free": false, "pro": false}}
```
