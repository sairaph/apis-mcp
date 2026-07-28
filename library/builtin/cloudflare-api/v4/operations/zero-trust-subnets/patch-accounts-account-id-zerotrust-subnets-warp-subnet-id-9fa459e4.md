---
title: Update WARP IP subnet
page_id: operation-patch-accounts-account-id-zerotrust-subnets-warp-subnet-id-73b3ba4c
path: operations/zero-trust-subnets
description: |-
    Updates a WARP IP assignment subnet.

    **Update constraints:**
    - The `network` field cannot be modified for WARP subnets. Only `name`, `comment`, and `is_default_network` can be updated.
    - IPv6 subnets cannot be updated
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PATCH
api_endpoints:
    - /accounts/{account_id}/zerotrust/subnets/warp/{subnet_id}
operation_ids:
    - zero-trust-networks-subnet-update-warp
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update WARP IP subnet

`PATCH /accounts/{account_id}/zerotrust/subnets/warp/{subnet_id}`

Operation ID: `zero-trust-networks-subnet-update-warp`

Updates a WARP IP assignment subnet.

**Update constraints:**
- The `network` field cannot be modified for WARP subnets. Only `name`, `comment`, and `is_default_network` can be updated.
- IPv6 subnets cannot be updated

## Definition

```yaml
{"operationId": "zero-trust-networks-subnet-update-warp", "summary": "Update WARP IP subnet", "description": "Updates a WARP IP assignment subnet.\n\n**Update constraints:**\n- The `network` field cannot be modified for WARP subnets. Only `name`, `comment`, and `is_default_network` can be updated.\n- IPv6 subnets cannot be updated\n", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tunnel_account_id"}}, {"name": "subnet_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tunnel_subnet_id"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"comment": {"$ref": "#/components/schemas/tunnel_subnet_comment"}, "is_default_network": {"$ref": "#/components/schemas/tunnel_subnet_is_default_network"}, "name": {"$ref": "#/components/schemas/tunnel_subnet_name"}, "network": {"$ref": "#/components/schemas/tunnel_subnet_ip_network"}}}}}}, "responses": {"200": {"description": "Update subnet response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tunnel_subnet_response_single"}}}}, "4XX": {"description": "Update subnet response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tunnel_subnet_response_single"}, {"$ref": "#/components/schemas/tunnel_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}, {"api_token": []}, {"user_service_key": []}], "tags": ["Zero Trust Subnets"], "x-api-token-group": ["Cloudflare One Networks Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.networks.subnets.warp", "x-fern-sdk-method-name": "edit", "x-forge-hidden": true}
```
