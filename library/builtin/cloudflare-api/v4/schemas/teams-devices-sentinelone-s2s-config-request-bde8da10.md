---
title: teams-devices_sentinelone_s2s_config_request
page_id: schema-teams-devices-sentinelone-s2s-config-request-bde8da10
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# teams-devices_sentinelone_s2s_config_request

```yaml
{"type": "object", "properties": {"api_url": {"description": "The SentinelOne S2S API URL.", "type": "string", "example": "https://example.sentinelone.net"}, "client_secret": {"description": "The SentinelOne S2S client secret.", "type": "string", "example": "example client secret", "x-sensitive": true}}, "required": ["api_url", "client_secret"], "title": "SentinelOne S2S Config"}
```
