---
title: waf-managed-rules_allowed_modes_anomaly
page_id: schema-waf-managed-rules-allowed-modes-anomaly-533cd139
path: schemas
description: Defines the available modes for the current WAF rule. Applies to anomaly detection WAF rules.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# waf-managed-rules_allowed_modes_anomaly

Defines the available modes for the current WAF rule. Applies to anomaly detection WAF rules.

```yaml
{"description": "Defines the available modes for the current WAF rule. Applies to anomaly detection WAF rules.", "type": "array", "items": {"$ref": "#/components/schemas/waf-managed-rules_mode_anomaly"}, "example": ["on", "off"], "readOnly": true}
```
