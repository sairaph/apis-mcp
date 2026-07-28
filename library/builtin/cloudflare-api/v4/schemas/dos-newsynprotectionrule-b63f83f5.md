---
title: dos_NewSynProtectionRule
page_id: schema-dos-newsynprotectionrule-b63f83f5
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dos_NewSynProtectionRule

```yaml
{"type": "object", "properties": {"burst_sensitivity": {"description": "The burst sensitivity. Must be one of 'low', 'medium', 'high'.", "type": "string", "x-auditable": true}, "mitigation_type": {"description": "The type of mitigation. Must be one of 'challenge' or 'retransmit'. Optional. Defaults to 'challenge'.", "type": "string", "x-auditable": true}, "mode": {"description": "The mode for SYN Protection. Must be one of 'enabled', 'disabled', 'monitoring'.", "type": "string", "x-auditable": true}, "name": {"description": "The name of the SYN Protection rule. Value is relative to the 'scope' setting. For 'global' scope, name should be 'global'. For either the 'region' or 'datacenter' scope, name should be the actual name of the region or datacenter, e.g., 'wnam' or 'lax'.", "type": "string", "x-auditable": true}, "rate_sensitivity": {"description": "The rate sensitivity. Must be one of 'low', 'medium', 'high'.", "type": "string", "x-auditable": true}, "scope": {"description": "The scope for the SYN Protection rule. Must be one of 'global', 'region', or 'datacenter'.", "type": "string", "x-auditable": true}}, "additionalProperties": false, "required": ["scope", "name", "mode", "rate_sensitivity", "burst_sensitivity"]}
```
