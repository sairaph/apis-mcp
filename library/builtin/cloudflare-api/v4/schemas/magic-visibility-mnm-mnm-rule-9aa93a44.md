---
title: magic-visibility-mnm_mnm_rule
page_id: schema-magic-visibility-mnm-mnm-rule-9aa93a44
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# magic-visibility-mnm_mnm_rule

```yaml
{"type": "object", "properties": {"automatic_advertisement": {"$ref": "#/components/schemas/magic-visibility-mnm_mnm_rule_automatic_advertisement"}, "bandwidth_threshold": {"$ref": "#/components/schemas/magic-visibility-mnm_mnm_rule_bandwidth_threshold"}, "duration": {"$ref": "#/components/schemas/magic-visibility-mnm_mnm_rule_duration"}, "id": {"$ref": "#/components/schemas/magic-visibility-mnm_rule_identifier"}, "name": {"$ref": "#/components/schemas/magic-visibility-mnm_mnm_rule_name"}, "packet_threshold": {"$ref": "#/components/schemas/magic-visibility-mnm_mnm_rule_packet_threshold"}, "prefix_match": {"$ref": "#/components/schemas/magic-visibility-mnm_mnm_rule_prefix_match"}, "prefixes": {"$ref": "#/components/schemas/magic-visibility-mnm_mnm_rule_ip_prefixes"}, "type": {"$ref": "#/components/schemas/magic-visibility-mnm_mnm_rule_type"}, "zscore_sensitivity": {"$ref": "#/components/schemas/magic-visibility-mnm_mnm_rule_zscore_sensitivity"}, "zscore_target": {"$ref": "#/components/schemas/magic-visibility-mnm_mnm_rule_zscore_target"}}, "nullable": true, "required": ["id", "name", "prefixes", "automatic_advertisement", "type"]}
```
