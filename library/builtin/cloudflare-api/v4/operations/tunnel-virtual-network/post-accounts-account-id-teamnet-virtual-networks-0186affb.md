---
title: Create a virtual network
page_id: operation-post-accounts-account-id-teamnet-virtual-networks-7a0ed079
path: operations/tunnel-virtual-network
description: Adds a new virtual network to an account.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/teamnet/virtual_networks
operation_ids:
    - tunnel-virtual-network-create-a-virtual-network
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create a virtual network

`POST /accounts/{account_id}/teamnet/virtual_networks`

Operation ID: `tunnel-virtual-network-create-a-virtual-network`

Adds a new virtual network to an account.

## Definition

```yaml
{"operationId": "tunnel-virtual-network-create-a-virtual-network", "summary": "Create a virtual network", "description": "Adds a new virtual network to an account.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tunnel_account_id"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"comment": {"$ref": "#/components/schemas/tunnel_virtual_network_comment"}, "is_default": {"description": "If `true`, this virtual network is the default for the account.", "type": "boolean", "example": true, "deprecated": true, "x-auditable": true, "x-stainless-deprecation-message": "Use the is_default_network property instead.", "x-stainless-ignore": true}, "is_default_network": {"$ref": "#/components/schemas/tunnel_is_default_network_optional"}, "name": {"$ref": "#/components/schemas/tunnel_virtual_network_name"}}, "required": ["name"]}}}}, "responses": {"200": {"description": "Create a virtual network response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tunnel_vnet_response_single"}}}}, "4XX": {"description": "Create a virtual network response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tunnel_vnet_response_single"}, {"$ref": "#/components/schemas/tunnel_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}, {"api_token": []}, {"user_service_key": []}], "tags": ["Tunnel Virtual Network"], "x-api-token-group": ["Cloudflare One Networks Write", "Cloudflare Tunnel Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.networks.virtual-networks", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
