---
title: List account organizations
page_id: operation-get-accounts-account-id-organizations-6277a812
path: operations/accounts
description: |-
    Retrieve a list of the organizations that "contain" this account or are
    managing it.

    The returned list will be in order from "root" to "leaf", where the "leaf"
    will be the organization that _immediately_ contains the specified
    account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/organizations
operation_ids:
    - Accounts_listAccountOrganizations
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# List account organizations

`GET /accounts/{account_id}/organizations`

Operation ID: `Accounts_listAccountOrganizations`

Retrieve a list of the organizations that "contain" this account or are
managing it.

The returned list will be in order from "root" to "leaf", where the "leaf"
will be the organization that _immediately_ contains the specified
account.

## Definition

```yaml
{"operationId": "Accounts_listAccountOrganizations", "summary": "List account organizations", "description": "Retrieve a list of the organizations that \"contain\" this account or are\nmanaging it.\n\nThe returned list will be in order from \"root\" to \"leaf\", where the \"leaf\"\nwill be the organization that _immediately_ contains the specified\naccount.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"type": "string"}}], "responses": {"200": {"description": "The request has succeeded.", "content": {"application/json": {"schema": {"type": "object", "properties": {"errors": {"type": "array", "items": {"type": "object"}, "maxItems": 0}, "messages": {"type": "array", "items": {"$ref": "#/components/schemas/organizations-api_V4Message"}}, "result": {"type": "array", "items": {"$ref": "#/components/schemas/organizations-api_Organization"}}, "success": {"type": "boolean", "enum": [true]}}, "required": ["success", "errors", "messages", "result"]}}}}, "4XX": {"description": "An unexpected error response.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/organizations-api_V4ErrorResponse"}}}}}, "security": [{"api_email": [], "api_key": []}], "tags": ["Accounts"], "x-api-token-group": ["Trust and Safety Write", "Trust and Safety Read", "DNS View Write", "DNS View Read", "SCIM Provisioning", "Load Balancers Account Write", "Load Balancers Account Read", "Zero Trust: PII Read", "DDoS Botnet Feed Write", "DDoS Botnet Feed Read", "Workers R2 Storage Write", "Workers R2 Storage Read", "DDoS Protection Write", "DDoS Protection Read", "Workers Tail Read", "Workers KV Storage Write", "Workers KV Storage Read", "Workers Scripts Write", "Workers Scripts Read", "Load Balancing: Monitors and Pools Write", "Load Balancing: Monitors and Pools Read", "Account Firewall Access Rules Write", "Account Firewall Access Rules Read", "DNS Firewall Write", "DNS Firewall Read", "Billing Write", "Billing Read", "Account Settings Write", "Account Settings Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "organizations.account-organizations", "x-fern-sdk-method-name": "list", "x-forge-hidden": true}
```
