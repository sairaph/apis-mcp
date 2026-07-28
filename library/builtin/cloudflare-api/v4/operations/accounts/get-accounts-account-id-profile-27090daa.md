---
title: Get account profile
page_id: operation-get-accounts-account-id-profile-bb210c12
path: operations/accounts
description: Retrieves the profile information for a specific Cloudflare account, including organization details, settings, and metadata. This endpoint is commonly used to verify account access and retrieve account-level configuration.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/profile
operation_ids:
    - Accounts_getAccountProfile
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get account profile

`GET /accounts/{account_id}/profile`

Operation ID: `Accounts_getAccountProfile`

Retrieves the profile information for a specific Cloudflare account, including organization details, settings, and metadata. This endpoint is commonly used to verify account access and retrieve account-level configuration.

## Definition

```yaml
{"operationId": "Accounts_getAccountProfile", "summary": "Get account profile", "description": "Retrieves the profile information for a specific Cloudflare account, including organization details, settings, and metadata. This endpoint is commonly used to verify account access and retrieve account-level configuration.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"description": "The request has succeeded.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"type": "object"}, "maxItems": 0}, "messages": {"type": "array", "items": {"$ref": "#/components/schemas/organizations-api_V4Message"}}, "result": {"$ref": "#/components/schemas/organizations-api_Profile"}, "success": {"type": "boolean", "enum": [true]}}, "required": ["success", "errors", "messages", "result"]}}}}, "4XX": {"description": "An unexpected error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/organizations-api_V4ErrorResponse"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Accounts"], "x-api-token-group": ["Trust and Safety Write", "Trust and Safety Read", "DNS View Write", "DNS View Read", "SCIM Provisioning", "Load Balancers Account Write", "Load Balancers Account Read", "Zero Trust: PII Read", "DDoS Botnet Feed Write", "DDoS Botnet Feed Read", "Workers R2 Storage Write", "Workers R2 Storage Read", "DDoS Protection Write", "DDoS Protection Read", "Workers Tail Read", "Workers KV Storage Write", "Workers KV Storage Read", "Workers Scripts Write", "Workers Scripts Read", "Load Balancing: Monitors and Pools Write", "Load Balancing: Monitors and Pools Read", "Account Firewall Access Rules Write", "Account Firewall Access Rules Read", "DNS Firewall Write", "DNS Firewall Read", "Billing Write", "Billing Read", "Account Settings Write", "Account Settings Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "accounts.account-profile", "x-fern-sdk-method-name": "get"}
```
