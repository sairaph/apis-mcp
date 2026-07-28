---
title: teams-devices_global_acceleration
page_id: schema-teams-devices-global-acceleration-9665193e
path: schemas
description: Global Acceleration settings for China. When configured, WARP clients connect to the Global Accelerator addresses instead of the default ones. Please contact your account representative to enable this feature on your account. See https://developers.cloudflare.com/china-network/concepts/global-acceleration/.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# teams-devices_global_acceleration

Global Acceleration settings for China. When configured, WARP clients connect to the Global Accelerator addresses instead of the default ones. Please contact your account representative to enable this feature on your account. See https://developers.cloudflare.com/china-network/concepts/global-acceleration/.

```yaml
{"description": "Global Acceleration settings for China. When configured, WARP clients connect to the Global Accelerator addresses instead of the default ones. Please contact your account representative to enable this feature on your account. See https://developers.cloudflare.com/china-network/concepts/global-acceleration/.", "type": "object", "properties": {"api_endpoints": {"description": "IP:port entries for the API endpoints.", "type": "array", "items": {"type": "string"}, "example": ["198.51.100.1:443"], "maxItems": 5, "x-auditable": true}, "enabled": {"description": "Global acceleration settings are used only when \"enabled\".", "type": "boolean", "example": true, "x-auditable": true}, "masque_endpoints": {"description": "IP:port entries for the MASQUE tunnel endpoints. Either wireguard_endpoints or masque_endpoints must be provided.", "type": "array", "items": {"type": "string"}, "example": ["198.51.100.1:443"], "maxItems": 5, "x-auditable": true}, "wireguard_endpoints": {"description": "IP:port entries for the WireGuard tunnel endpoints. Either wireguard_endpoints or masque_endpoints must be provided.", "type": "array", "items": {"type": "string"}, "example": ["198.51.100.1:2408"], "maxItems": 5, "x-auditable": true}}, "nullable": true, "required": ["enabled", "wireguard_endpoints", "masque_endpoints", "api_endpoints"], "x-auditable": true}
```
