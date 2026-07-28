---
title: dos_DnsProtectionRuleUpdate
page_id: schema-dos-dnsprotectionruleupdate-3c53f609
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dos_DnsProtectionRuleUpdate

```yaml
{"type": "object", "properties": {"block_any_queries": {"description": "The new value for whether to block DNS ANY queries. Optional.", "type": "boolean", "x-auditable": true}, "burst_sensitivity": {"description": "The new burst sensitivity. Optional. Must be one of 'low', 'medium', 'high'.", "type": "string", "x-auditable": true}, "mode": {"description": "The new mode for DNS Protection. Optional. Must be one of 'enabled', 'disabled', 'monitoring'.", "type": "string", "x-auditable": true}, "profile_sensitivity": {"description": "The new profile sensitivity. Optional. Recommended setting is 'low'. Must be one of 'low', 'medium', 'high', or 'very_high'.", "type": "string", "x-auditable": true}, "rate_sensitivity": {"description": "The new rate sensitivity. Optional. Must be one of 'low', 'medium', 'high'.", "type": "string", "x-auditable": true}}, "additionalProperties": false}
```
