---
title: Update WARP Connector HA configuration
page_id: operation-put-accounts-account-id-warp-connector-tunnel-id-configurations-81d7ec61
path: operations/cloudflare-tunnel-configuration
description: Adds or updates the high-availability configuration for a WARP Connector tunnel.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - PUT
api_endpoints:
    - /accounts/{account_id}/warp_connector/{tunnel_id}/configurations
operation_ids:
    - cloudflare-tunnel-configuration-update-warp-connector-configuration
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Update WARP Connector HA configuration

`PUT /accounts/{account_id}/warp_connector/{tunnel_id}/configurations`

Operation ID: `cloudflare-tunnel-configuration-update-warp-connector-configuration`

Adds or updates the high-availability configuration for a WARP Connector tunnel.

## Definition

```yaml
{"operationId": "cloudflare-tunnel-configuration-update-warp-connector-configuration", "summary": "Update WARP Connector HA configuration", "description": "Adds or updates the high-availability configuration for a WARP Connector tunnel.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tunnel_identifier"}}, {"name": "tunnel_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tunnel_tunnel_id-2"}}], "requestBody": {"required": true, "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tunnel_mesh_configuration_request_body"}}}}, "responses": {"200": {"description": "Update WARP Connector HA configuration response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tunnel_mesh_configuration_response_single"}}}}, "4XX": {"description": "Update WARP Connector HA configuration response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tunnel_api-response-common-failure-2"}}}}}, "security": [{"api_email": [], "api_key": []}, {"api_token": []}, {"user_service_key": []}], "tags": ["Cloudflare Tunnel Configuration"], "x-api-token-group": ["Cloudflare One Connectors Write", "Cloudflare One Connector: WARP Write"]}
```
