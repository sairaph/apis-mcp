---
title: Create WARP IP subnet
page_id: operation-post-accounts-account-id-zerotrust-subnets-warp-e67765db
path: operations/zero-trust-subnets
description: |-
    Create a WARP IP assignment subnet. Currently, only IPv4 subnets can be created.

    **Network constraints:**
    - The network must be within one of the following private IP ranges:
      - `10.0.0.0/8` (RFC 1918)
      - `172.16.0.0/12` (RFC 1918)
      - `192.168.0.0/16` (RFC 1918)
      - `100.64.0.0/10` (RFC 6598 - CGNAT)
    - The subnet must have a prefix length of `/24` or larger (e.g., `/16`, `/20`, `/24` are valid; `/25`, `/28` are not)
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - POST
api_endpoints:
    - /accounts/{account_id}/zerotrust/subnets/warp
operation_ids:
    - zero-trust-networks-subnet-create-warp
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Create WARP IP subnet

`POST /accounts/{account_id}/zerotrust/subnets/warp`

Operation ID: `zero-trust-networks-subnet-create-warp`

Create a WARP IP assignment subnet. Currently, only IPv4 subnets can be created.

**Network constraints:**
- The network must be within one of the following private IP ranges:
  - `10.0.0.0/8` (RFC 1918)
  - `172.16.0.0/12` (RFC 1918)
  - `192.168.0.0/16` (RFC 1918)
  - `100.64.0.0/10` (RFC 6598 - CGNAT)
- The subnet must have a prefix length of `/24` or larger (e.g., `/16`, `/20`, `/24` are valid; `/25`, `/28` are not)

## Definition

```yaml
{"operationId": "zero-trust-networks-subnet-create-warp", "summary": "Create WARP IP subnet", "description": "Create a WARP IP assignment subnet. Currently, only IPv4 subnets can be created.\n\n**Network constraints:**\n- The network must be within one of the following private IP ranges:\n  - `10.0.0.0/8` (RFC 1918)\n  - `172.16.0.0/12` (RFC 1918)\n  - `192.168.0.0/16` (RFC 1918)\n  - `100.64.0.0/10` (RFC 6598 - CGNAT)\n- The subnet must have a prefix length of `/24` or larger (e.g., `/16`, `/20`, `/24` are valid; `/25`, `/28` are not)\n", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tunnel_account_id"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"comment": {"$ref": "#/components/schemas/tunnel_subnet_comment"}, "is_default_network": {"$ref": "#/components/schemas/tunnel_subnet_is_default_network"}, "name": {"$ref": "#/components/schemas/tunnel_subnet_name"}, "network": {"$ref": "#/components/schemas/tunnel_subnet_ip_network"}}, "required": ["name", "network"]}}}}, "responses": {"200": {"description": "Create subnet response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tunnel_subnet_response_single"}}}}, "4XX": {"description": "Create subnet response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tunnel_subnet_response_single"}, {"$ref": "#/components/schemas/tunnel_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}, {"api_token": []}, {"user_service_key": []}], "tags": ["Zero Trust Subnets"], "x-api-token-group": ["Cloudflare One Networks Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.networks.subnets.warp", "x-fern-sdk-method-name": "create", "x-forge-hidden": true}
```
