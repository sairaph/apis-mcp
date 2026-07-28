---
title: dos_TcpFlowProtectionRule
page_id: schema-dos-tcpflowprotectionrule-e0850961
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dos_TcpFlowProtectionRule

```yaml
{"type": "object", "properties": {"burst_sensitivity": {"description": "The burst sensitivity. Must be one of 'low', 'medium', 'high'.", "type": "string", "x-auditable": true}, "created_on": {"description": "The creation timestamp of the TCP Flow Protection rule.", "type": "string", "format": "date-time", "x-auditable": true}, "id": {"description": "The unique ID of the TCP Flow Protection rule.", "type": "string", "x-auditable": true}, "mode": {"description": "The mode for TCP Flow Protection. Must be one of 'enabled', 'disabled', 'monitoring'.", "type": "string", "x-auditable": true}, "modified_on": {"description": "The last modification timestamp of the TCP Flow Protection rule.", "type": "string", "format": "date-time", "x-auditable": true}, "name": {"description": "The name of the TCP Flow Protection rule. Value is relative to the 'scope' setting. For 'global' scope, name should be 'global'. For either the 'region' or 'datacenter' scope, name should be the actual name of the region or datacenter, e.g., 'wnam' or 'lax'.", "type": "string", "x-auditable": true}, "rate_sensitivity": {"description": "The rate sensitivity. Must be one of 'low', 'medium', 'high'.", "type": "string", "x-auditable": true}, "scope": {"description": "The scope for the TCP Flow Protection rule. Must be one of 'global', 'region', or 'datacenter'.", "type": "string", "x-auditable": true}}, "additionalProperties": false, "required": ["id", "scope", "name", "mode", "rate_sensitivity", "burst_sensitivity", "created_on", "modified_on"]}
```
