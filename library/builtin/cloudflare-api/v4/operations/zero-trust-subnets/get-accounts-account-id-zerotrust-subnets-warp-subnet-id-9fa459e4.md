---
title: Get WARP IP subnet
page_id: operation-get-accounts-account-id-zerotrust-subnets-warp-subnet-id-ead669bf
path: operations/zero-trust-subnets
description: Get a WARP IP assignment subnet.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/zerotrust/subnets/warp/{subnet_id}
operation_ids:
    - zero-trust-networks-subnet-get-warp
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get WARP IP subnet

`GET /accounts/{account_id}/zerotrust/subnets/warp/{subnet_id}`

Operation ID: `zero-trust-networks-subnet-get-warp`

Get a WARP IP assignment subnet.

## Definition

```yaml
{"operationId": "zero-trust-networks-subnet-get-warp", "summary": "Get WARP IP subnet", "description": "Get a WARP IP assignment subnet.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tunnel_account_id"}}, {"name": "subnet_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tunnel_subnet_id"}}], "responses": {"200": {"description": "Get subnet response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tunnel_subnet_response_single"}}}}, "4XX": {"description": "Get subnet response failure", "content": {"application/json": {"schema": {"allOf": [{"$ref": "#/components/schemas/tunnel_subnet_response_single"}, {"$ref": "#/components/schemas/tunnel_api-response-common-failure"}]}}}}}, "security": [{"api_email": [], "api_key": []}, {"api_token": []}, {"user_service_key": []}], "tags": ["Zero Trust Subnets"], "x-api-token-group": ["Cloudflare One Networks Write", "Cloudflare One Networks Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.networks.subnets.warp", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
