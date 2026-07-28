---
title: dos_SynProtectionRuleUpdate
page_id: schema-dos-synprotectionruleupdate-a5ba1bc8
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dos_SynProtectionRuleUpdate

```yaml
{"type": "object", "properties": {"burst_sensitivity": {"description": "The new burst sensitivity. Optional. Must be one of 'low', 'medium', 'high'.", "type": "string", "x-auditable": true}, "mitigation_type": {"description": "The new mitigation type. Optional. Must be one of 'challenge' or 'retransmit'.", "type": "string", "x-auditable": true}, "mode": {"description": "The new mode for SYN Protection. Optional. Must be one of 'enabled', 'disabled', 'monitoring'.", "type": "string", "x-auditable": true}, "rate_sensitivity": {"description": "The new rate sensitivity. Optional. Must be one of 'low', 'medium', 'high'.", "type": "string", "x-auditable": true}}, "additionalProperties": false}
```
