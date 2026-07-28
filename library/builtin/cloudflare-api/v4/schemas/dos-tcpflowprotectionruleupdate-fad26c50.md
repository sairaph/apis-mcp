---
title: dos_TcpFlowProtectionRuleUpdate
page_id: schema-dos-tcpflowprotectionruleupdate-fad26c50
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dos_TcpFlowProtectionRuleUpdate

```yaml
{"type": "object", "properties": {"burst_sensitivity": {"description": "The new burst sensitivity. Optional. Must be one of 'low', 'medium', 'high'.", "type": "string", "x-auditable": true}, "mode": {"description": "The new mode for TCP Flow Protection. Optional. Must be one of 'enabled', 'disabled', 'monitoring'.", "type": "string", "x-auditable": true}, "rate_sensitivity": {"description": "The new rate sensitivity. Optional. Must be one of 'low', 'medium', 'high'.", "type": "string", "x-auditable": true}}, "additionalProperties": false}
```
