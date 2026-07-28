---
title: Fetch Worker Account Settings
page_id: operation-get-accounts-account-id-workers-account-settings-3b673d3e
path: operations/worker-account-settings
description: Fetches Worker account settings for an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/workers/account-settings
operation_ids:
    - worker-account-settings-fetch-worker-account-settings
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Fetch Worker Account Settings

`GET /accounts/{account_id}/workers/account-settings`

Operation ID: `worker-account-settings-fetch-worker-account-settings`

Fetches Worker account settings for an account.

## Definition

```yaml
{"operationId": "worker-account-settings-fetch-worker-account-settings", "summary": "Fetch Worker Account Settings", "description": "Fetches Worker account settings for an account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/workers_identifier"}}], "responses": {"200": {"description": "Fetch Worker Account Settings response.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/workers_api-response-common"}, {"properties": {"result": {"$ref": "#/components/schemas/workers_account-settings"}}, "required": ["result"], "type": "object"}]}}}}, "4XX": {"description": "Fetch Worker Account Settings response failure.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/workers_api-response-common-failure"}}}}}, "security": [{"api_token": []}, {"api_email": [], "api_key": []}], "tags": ["Worker Account Settings"], "x-api-token-group": ["Trust and Safety Write", "Trust and Safety Read", "DNS View Write", "DNS View Read", "SCIM Provisioning", "Load Balancers Account Write", "Load Balancers Account Read", "Zero Trust: PII Read", "DDoS Botnet Feed Write", "DDoS Botnet Feed Read", "Workers R2 Storage Write", "Workers R2 Storage Read", "DDoS Protection Write", "DDoS Protection Read", "Workers Tail Read", "Workers KV Storage Write", "Workers KV Storage Read", "Workers Scripts Write", "Workers Scripts Read", "Load Balancing: Monitors and Pools Write", "Load Balancing: Monitors and Pools Read", "Account Firewall Access Rules Write", "Account Firewall Access Rules Read", "DNS Firewall Write", "DNS Firewall Read", "Billing Write", "Billing Read", "Account Settings Write", "Account Settings Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "workers.account-settings", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
