---
title: Update Gateway Ephemeral Subnet
page_id: operation-put-accounts-account-id-zerotrust-subnets-initial-resolved-ip-address-fa-5e4869ac
path: operations/zero-trust-subnets
description: |-
    Updates the CIDR for the account's default gateway ephemeral subnet of the given address
    family. The new CIDR must not conflict with existing private routes in the account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/zerotrust/subnets/initial_resolved_ip/{address_family}
operation_ids:
    - zero-trust-networks-subnet-update-gateway-ephemeral
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update Gateway Ephemeral Subnet

`PUT /accounts/{account_id}/zerotrust/subnets/initial_resolved_ip/{address_family}`

Operation ID: `zero-trust-networks-subnet-update-gateway-ephemeral`

Updates the CIDR for the account's default gateway ephemeral subnet of the given address
family. The new CIDR must not conflict with existing private routes in the account.

## Definition

```yaml
{"operationId": "zero-trust-networks-subnet-update-gateway-ephemeral", "summary": "Update Gateway Ephemeral Subnet", "description": "Updates the CIDR for the account's default gateway ephemeral subnet of the given address\nfamily. The new CIDR must not conflict with existing private routes in the account.\n", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tunnel_account_id"}}, {"name": "address_family", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tunnel_address_family"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"comment": {"$ref": "#/components/schemas/tunnel_subnet_comment"}, "name": {"$ref": "#/components/schemas/tunnel_subnet_name"}, "network": {"$ref": "#/components/schemas/tunnel_subnet_ip_network"}}}}}}, "responses": {"200": {"description": "Update gateway ephemeral subnet response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tunnel_subnet_response_single"}}}}, "4XX": {"description": "Update gateway ephemeral subnet response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tunnel_subnet_response_single"}, {"$ref": "#/components/schemas/tunnel_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}, {"api_token": []}, {"user_service_key": []}], "tags": ["Zero Trust Subnets"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.networks.subnets.initial-resolved-ip", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
