---
title: magic_bgp_status_with_state
page_id: schema-magic-bgp-status-with-state-83707906
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# magic_bgp_status_with_state

```yaml
{"type": "object", "properties": {"bgp_state": {"type": "string"}, "cf_speaker_ip": {"type": "string", "format": "ipv4"}, "cf_speaker_port": {"type": "integer", "maximum": 65535, "minimum": 1}, "customer_speaker_ip": {"type": "string", "format": "ipv4"}, "customer_speaker_port": {"type": "integer", "maximum": 65535, "minimum": 1}, "state": {"type": "string", "enum": ["BGP_DOWN", "BGP_UP", "BGP_ESTABLISHING"]}, "tcp_established": {"type": "boolean"}, "updated_at": {"type": "string", "format": "date-time"}}, "required": ["state", "tcp_established", "updated_at"]}
```
