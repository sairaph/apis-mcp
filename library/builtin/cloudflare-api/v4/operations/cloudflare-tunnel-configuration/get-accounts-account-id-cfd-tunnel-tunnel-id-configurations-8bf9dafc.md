---
title: Get configuration
page_id: operation-get-accounts-account-id-cfd-tunnel-tunnel-id-configurations-6d0983af
path: operations/cloudflare-tunnel-configuration
description: Gets the configuration for a remotely-managed tunnel
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/cfd_tunnel/{tunnel_id}/configurations
operation_ids:
    - cloudflare-tunnel-configuration-get-configuration
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get configuration

`GET /accounts/{account_id}/cfd_tunnel/{tunnel_id}/configurations`

Operation ID: `cloudflare-tunnel-configuration-get-configuration`

Gets the configuration for a remotely-managed tunnel

## Definition

```yaml
{"operationId": "cloudflare-tunnel-configuration-get-configuration", "summary": "Get configuration", "description": "Gets the configuration for a remotely-managed tunnel", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tunnel_identifier"}}, {"name": "tunnel_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tunnel_tunnel_id-2"}}], "responses": {"200": {"description": "Get configuration response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tunnel_configuration_response"}}}}, "4XX": {"description": "Get configuration response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tunnel_api-response-common-failure-2"}}}}}, "security": [{"api_email": [], "api_key": []}, {"api_token": []}, {"user_service_key": []}], "tags": ["Cloudflare Tunnel Configuration"], "x-api-token-group": ["Cloudflare One Connectors Write", "Cloudflare One Connectors Read", "Cloudflare One Connector: cloudflared Write", "Cloudflare One Connector: cloudflared Read", "Cloudflare Tunnel Write", "Cloudflare Tunnel Read"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.tunnels.cloudflared.configurations", "x-fern-sdk-method-name": "get", "x-forge-hidden": true}
```
