---
title: Update Cloudflare Source Subnet
page_id: operation-patch-accounts-account-id-zerotrust-subnets-cloudflare-source-address-fa-99c44c16
path: operations/zero-trust-subnets
description: Updates the Cloudflare Source subnet of the given address family
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/zerotrust/subnets/cloudflare_source/{address_family}
operation_ids:
    - zero-trust-networks-subnet-update-cloudflare-source
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update Cloudflare Source Subnet

`PATCH /accounts/{account_id}/zerotrust/subnets/cloudflare_source/{address_family}`

Operation ID: `zero-trust-networks-subnet-update-cloudflare-source`

Updates the Cloudflare Source subnet of the given address family

## Definition

```yaml
{"operationId": "zero-trust-networks-subnet-update-cloudflare-source", "summary": "Update Cloudflare Source Subnet", "description": "Updates the Cloudflare Source subnet of the given address family", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tunnel_account_id"}}, {"name": "address_family", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tunnel_address_family"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"comment": {"$ref": "#/components/schemas/tunnel_subnet_comment"}, "name": {"$ref": "#/components/schemas/tunnel_subnet_name"}, "network": {"$ref": "#/components/schemas/tunnel_subnet_ip_network"}}}}}}, "responses": {"200": {"description": "Update subnet response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tunnel_subnet_response_single"}}}}, "4XX": {"description": "Update subnet response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tunnel_subnet_response_single"}, {"$ref": "#/components/schemas/tunnel_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}, {"api_token": []}, {"user_service_key": []}], "tags": ["Zero Trust Subnets"], "x-api-token-group": ["Cloudflare One Networks Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.networks.subnets.cloudflare-source", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
