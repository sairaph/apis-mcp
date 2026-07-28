---
title: Get WARP Connector HA configuration
page_id: operation-get-accounts-account-id-warp-connector-tunnel-id-configurations-45be3db3
path: operations/cloudflare-tunnel-configuration
description: Gets the high-availability configuration for a WARP Connector tunnel.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
http_methods:
    - GET
api_endpoints:
    - /accounts/{account_id}/warp_connector/{tunnel_id}/configurations
operation_ids:
    - cloudflare-tunnel-configuration-get-warp-connector-configuration
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# Get WARP Connector HA configuration

`GET /accounts/{account_id}/warp_connector/{tunnel_id}/configurations`

Operation ID: `cloudflare-tunnel-configuration-get-warp-connector-configuration`

Gets the high-availability configuration for a WARP Connector tunnel.

## Definition

```yaml
{"operationId": "cloudflare-tunnel-configuration-get-warp-connector-configuration", "summary": "Get WARP Connector HA configuration", "description": "Gets the high-availability configuration for a WARP Connector tunnel.", "parameters": [{"name": "account_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tunnel_identifier"}}, {"name": "tunnel_id", "in": "path", "required": true, "schema": {"$ref": "#/components/schemas/tunnel_tunnel_id-2"}}], "responses": {"200": {"description": "Get WARP Connector HA configuration response", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tunnel_mesh_configuration_response_single"}}}}, "4XX": {"description": "Get WARP Connector HA configuration response failure", "content": {"application/json": {"schema": {"$ref": "#/components/schemas/tunnel_api-response-common-failure-2"}}}}}, "security": [{"api_email": [], "api_key": []}, {"api_token": []}, {"user_service_key": []}], "tags": ["Cloudflare Tunnel Configuration"], "x-api-token-group": ["Cloudflare One Connectors Write", "Cloudflare One Connectors Read", "Cloudflare One Connector: WARP Write", "Cloudflare One Connector: WARP Read"]}
```
