---
title: dos_NewDnsProtectionRule
page_id: schema-dos-newdnsprotectionrule-1549fcfa
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dos_NewDnsProtectionRule

```yaml
{"type": "object", "properties": {"block_any_queries": {"description": "Whether to block DNS ANY queries. Optional. Defaults to true.", "type": "boolean", "x-auditable": true}, "burst_sensitivity": {"description": "The burst sensitivity. Must be one of 'low', 'medium', 'high'.", "type": "string", "x-auditable": true}, "mode": {"description": "The mode for DNS Protection. Must be one of 'enabled', 'disabled', 'monitoring'.", "type": "string", "x-auditable": true}, "name": {"description": "The name of the DNS Protection rule. Value is relative to the 'scope' setting. For 'global' scope, name should be 'global'. For either the 'region' or 'datacenter' scope, name should be the actual name of the region or datacenter, e.g., 'wnam' or 'lax'.", "type": "string", "x-auditable": true}, "profile_sensitivity": {"description": "The profile sensitivity. Recommended setting is 'low'. Must be one of 'low', 'medium', 'high', or 'very_high'.", "type": "string", "x-auditable": true}, "rate_sensitivity": {"description": "The rate sensitivity. Must be one of 'low', 'medium', 'high'.", "type": "string", "x-auditable": true}, "scope": {"description": "The scope for the DNS Protection rule. Must be one of 'global', 'region', or 'datacenter'.", "type": "string", "x-auditable": true}}, "additionalProperties": false, "required": ["scope", "name", "mode", "profile_sensitivity", "rate_sensitivity", "burst_sensitivity"]}
```
