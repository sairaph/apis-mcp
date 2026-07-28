---
title: google-tag-gateway_google-tag-gateway-config
page_id: schema-google-tag-gateway-google-tag-gateway-config-433d328e
path: schemas
description: Google Tag Gateway configuration for a zone.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# google-tag-gateway_google-tag-gateway-config

Google Tag Gateway configuration for a zone.

```yaml
{"description": "Google Tag Gateway configuration for a zone.", "type": "object", "properties": {"enabled": {"description": "Enables or disables Google Tag Gateway for this zone.", "type": "boolean", "example": true, "x-auditable": true}, "endpoint": {"description": "Specifies the endpoint path for proxying Google Tag Manager requests. Use an absolute path starting with '/', with no nested paths and alphanumeric characters only (e.g. /metrics).", "type": "string", "example": "/metrics", "x-auditable": true}, "hideOriginalIp": {"description": "Hides the original client IP address from Google when enabled.", "type": "boolean", "example": true, "x-auditable": true}, "measurementId": {"description": "Specify the Google Tag Manager container or measurement ID (e.g. GTM-XXXXXXX or G-XXXXXXXXXX).", "type": "string", "example": "GTM-P2F3N47Q", "x-auditable": true}, "setUpTag": {"description": "Set up the associated Google Tag on the zone automatically when enabled.", "type": "boolean", "example": true, "nullable": true, "x-auditable": true}}, "required": ["enabled", "endpoint", "hideOriginalIp", "measurementId"], "title": "Google Tag Gateway configuration"}
```
