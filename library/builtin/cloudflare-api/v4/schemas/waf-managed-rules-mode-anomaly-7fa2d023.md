---
title: waf-managed-rules_mode_anomaly
page_id: schema-waf-managed-rules-mode-anomaly-7fa2d023
path: schemas
description: Defines the mode anomaly. When set to `on`, the current WAF rule will be used when evaluating the request. Applies to anomaly detection WAF rules.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# waf-managed-rules_mode_anomaly

Defines the mode anomaly. When set to `on`, the current WAF rule will be used when evaluating the request. Applies to anomaly detection WAF rules.

```yaml
{"description": "Defines the mode anomaly. When set to `on`, the current WAF rule will be used when evaluating the request. Applies to anomaly detection WAF rules.", "type": "string", "example": "on", "enum": ["on", "off"], "x-auditable": true}
```
