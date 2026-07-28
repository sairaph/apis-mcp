---
title: dns-firewall_attack_mitigation
page_id: schema-dns-firewall-attack-mitigation-21803325
path: schemas
description: Attack mitigation settings
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dns-firewall_attack_mitigation

Attack mitigation settings

```yaml
{"description": "Attack mitigation settings", "type": "object", "properties": {"enabled": {"description": "When enabled, automatically mitigate random-prefix attacks to protect upstream DNS servers", "type": "boolean", "example": true, "x-auditable": true}, "only_when_upstream_unhealthy": {"description": "Only mitigate attacks when upstream servers seem unhealthy", "type": "boolean", "example": false, "default": true, "x-auditable": true}}, "nullable": true}
```
