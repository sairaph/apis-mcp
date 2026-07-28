---
title: Permission Group Details
page_id: operation-get-accounts-account-id-iam-permission-groups-permission-group-id-057330f8
path: operations/account-permission-groups
description: Get information about a specific permission group in an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/iam/permission_groups/{permission_group_id}
operation_ids:
    - account-permission-group-details
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Permission Group Details

`GET /accounts/{account_id}/iam/permission_groups/{permission_group_id}`

Operation ID: `account-permission-group-details`

Get information about a specific permission group in an account.

## Definition

```yaml
{"operationId": "account-permission-group-details", "summary": "Permission Group Details", "description": "Get information about a specific permission group in an account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_account_identifier"}}, {"name": "permission_group_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/iam_permission_group_identifier"}}], "responses": {"200": {"description": "Permission Group Details response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_single_permission_groups_response"}}}}, "4XX": {"description": "Permission Group Details response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/iam_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Account Permission Groups"], "x-api-token-group": ["Trust and Safety Write", "Trust and Safety Read", "DNS View Write", "DNS View Read", "SCIM Provisioning", "Load Balancers Account Write", "Load Balancers Account Read", "Zero Trust: PII Read", "DDoS Botnet Feed Write", "DDoS Botnet Feed Read", "Workers R2 Storage Write", "Workers R2 Storage Read", "DDoS Protection Write", "DDoS Protection Read", "Workers Tail Read", "Workers KV Storage Write", "Workers KV Storage Read", "Workers Scripts Write", "Workers Scripts Read", "Load Balancing: Monitors and Pools Write", "Load Balancing: Monitors and Pools Read", "Account Firewall Access Rules Write", "Account Firewall Access Rules Read", "DNS Firewall Write", "DNS Firewall Read", "Billing Write", "Billing Read", "Account Settings Write", "Account Settings Read"], "x-cfPermissionsRequired": {"enum": ["com.cloudflare.api.iam.permission-group.read"]}, "x-cfPlanAvailability": {"business": true, "enterprise": true, "free": true, "pro": true}}
```
