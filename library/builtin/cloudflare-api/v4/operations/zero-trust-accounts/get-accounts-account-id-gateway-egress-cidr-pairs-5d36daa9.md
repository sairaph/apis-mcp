---
title: Get gateway egress CIDRs pairs assigned to this account
page_id: operation-get-accounts-account-id-gateway-egress-cidr-pairs-e511d927
path: operations/zero-trust-accounts
description: Retrieve the list of egress CIDRs allocated to this Zero Trust account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/gateway/egress_cidr_pairs
operation_ids:
    - zero-trust-accounts-get-egress-cidr-pairs
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get gateway egress CIDRs pairs assigned to this account

`GET /accounts/{account_id}/gateway/egress_cidr_pairs`

Operation ID: `zero-trust-accounts-get-egress-cidr-pairs`

Retrieve the list of egress CIDRs allocated to this Zero Trust account.

## Definition

```yaml
{"operationId": "zero-trust-accounts-get-egress-cidr-pairs", "summary": "Get gateway egress CIDRs pairs assigned to this account", "description": "Retrieve the list of egress CIDRs allocated to this Zero Trust account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/zero-trust-gateway_identifier"}}], "responses": {"200": {"description": "List of egress CIDR pairs.", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/zero-trust-gateway_gateway-account-egress-cidr-pair-list-response"}}}}, "4XX": {"description": "List of egress CIDR pairs response failure.", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/zero-trust-gateway_gateway-account-egress-cidr-pair-list-response"}, {"$ref": "#/components/schemas/zero-trust-gateway_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": [], "api_token": []}], "tags": ["Zero Trust accounts"], "x-api-token-group": ["Zero Trust Read", "Zero Trust Write"]}
```
