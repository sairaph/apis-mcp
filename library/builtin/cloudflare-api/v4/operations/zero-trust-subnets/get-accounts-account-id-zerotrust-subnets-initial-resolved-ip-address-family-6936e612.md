---
title: Get Gateway Ephemeral Subnet
page_id: operation-get-accounts-account-id-zerotrust-subnets-initial-resolved-ip-address-fa-3c7fb491
path: operations/zero-trust-subnets
description: Returns the account's default gateway ephemeral subnet for the given address family.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/zerotrust/subnets/initial_resolved_ip/{address_family}
operation_ids:
    - zero-trust-networks-subnet-get-gateway-ephemeral
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get Gateway Ephemeral Subnet

`GET /accounts/{account_id}/zerotrust/subnets/initial_resolved_ip/{address_family}`

Operation ID: `zero-trust-networks-subnet-get-gateway-ephemeral`

Returns the account's default gateway ephemeral subnet for the given address family.

## Definition

```yaml
{"operationId": "zero-trust-networks-subnet-get-gateway-ephemeral", "summary": "Get Gateway Ephemeral Subnet", "description": "Returns the account's default gateway ephemeral subnet for the given address family.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tunnel_account_id"}}, {"name": "address_family", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tunnel_address_family"}}], "responses": {"200": {"description": "Get gateway ephemeral subnet response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tunnel_subnet_response_single"}}}}, "4XX": {"description": "Get gateway ephemeral subnet response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tunnel_subnet_response_single"}, {"$ref": "#/components/schemas/tunnel_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}, {"api_token": []}, {"user_service_key": []}], "tags": ["Zero Trust Subnets"], "x-api-token-group": ["Cloudflare One Networks Write", "Cloudflare One Networks Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.networks.subnets.initial-resolved-ip", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
