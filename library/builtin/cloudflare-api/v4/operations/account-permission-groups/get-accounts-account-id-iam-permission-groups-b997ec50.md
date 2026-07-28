---
title: List Account Permission Groups
page_id: operation-get-accounts-account-id-iam-permission-groups-4b7f0770
path: operations/account-permission-groups
description: List all the permissions groups for an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/iam/permission_groups
operation_ids:
    - account-permission-group-list
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List Account Permission Groups

`GET /accounts/{account_id}/iam/permission_groups`

Operation ID: `account-permission-group-list`

List all the permissions groups for an account.

## Definition

```yaml
{"operationId": "account-permission-group-list", "summary": "List Account Permission Groups", "description": "List all the permissions groups for an account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_account_identifier"}}, {"name": "id", "in": "query", "schema": {"description": "ID of the permission group to be fetched.", "type": "string", "example": "6d7f2f5f5b1d4a0e9081fdc98d432fd1", "maxLength": 32, "minLength": 32}}, {"name": "name", "in": "query", "schema": {"description": "Name of the permission group to be fetched.", "type": "string", "example": "NameOfThePermissionGroup"}}, {"name": "label", "in": "query", "schema": {"description": "Label of the permission group to be fetched.", "type": "string", "example": "labelOfThePermissionGroup"}}, {"name": "page", "in": "query", "schema": {"description": "Page number of paginated results.", "type": "number", "default": 1, "minimum": 1}}, {"name": "per_page", "in": "query", "schema": {"description": "Maximum number of results per page.", "type": "number", "default": 20, "maximum": 50, "minimum": 5}}], "responses": {"200": {"description": "List Permission Groups response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_collection_permission_groups_response"}}}}, "4XX": {"description": "List Permission Groups response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Account Permission Groups"], "x-api-token-group": ["Trust and Safety Write", "Trust and Safety Read", "DNS View Write", "DNS View Read", "SCIM Provisioning", "Load Balancers Account Write", "Load Balancers Account Read", "Zero Trust: PII Read", "DDoS Botnet Feed Write", "DDoS Botnet Feed Read", "Workers R2 Storage Write", "Workers R2 Storage Read", "DDoS Protection Write", "DDoS Protection Read", "Workers Tail Read", "Workers KV Storage Write", "Workers KV Storage Read", "Workers Scripts Write", "Workers Scripts Read", "Load Balancing: Monitors and Pools Write", "Load Balancing: Monitors and Pools Read", "Account Firewall Access Rules Write", "Account Firewall Access Rules Read", "DNS Firewall Write", "DNS Firewall Read", "Billing Write", "Billing Read", "Account Settings Write", "Account Settings Read"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.iam.permission-group.list"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
