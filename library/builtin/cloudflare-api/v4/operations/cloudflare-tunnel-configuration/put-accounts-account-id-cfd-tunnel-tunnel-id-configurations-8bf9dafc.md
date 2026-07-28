---
title: Put configuration
page_id: operation-put-accounts-account-id-cfd-tunnel-tunnel-id-configurations-1ae35233
path: operations/cloudflare-tunnel-configuration
description: Adds or updates the configuration for a remotely-managed tunnel.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/cfd_tunnel/{tunnel_id}/configurations
operation_ids:
    - cloudflare-tunnel-configuration-put-configuration
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Put configuration

`PUT /accounts/{account_id}/cfd_tunnel/{tunnel_id}/configurations`

Operation ID: `cloudflare-tunnel-configuration-put-configuration`

Adds or updates the configuration for a remotely-managed tunnel.

## Definition

```yaml
{"operationId": "cloudflare-tunnel-configuration-put-configuration", "summary": "Put configuration", "description": "Adds or updates the configuration for a remotely-managed tunnel.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tunnel_identifier"}}, {"name": "tunnel_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tunnel_tunnel_id-2"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"type": "object", "properties": {"config": {"$ref": "#/components/schemas/tunnel_config"}}}}}}, "responses": {"200": {"description": "Put configuration response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tunnel_configuration_response"}}}}, "4XX": {"description": "Put configuration response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tunnel_api-response-common-failure-2"}}}}}, "security": [{"api_email": [], "api_key": []}, {"api_token": []}, {"user_service_key": []}], "tags": ["Cloudflare Tunnel Configuration"], "x-api-token-group": ["Cloudflare One Connectors Write", "Cloudflare One Connector: cloudflared Write", "Cloudflare Tunnel Write"], "x-fern-availability": "generally-available", "x-fern-sdk-group-name": "zero-trust.tunnels.cloudflared.configurations", "x-fern-sdk-method-name": "update", "x-forge-hidden": true}
```
