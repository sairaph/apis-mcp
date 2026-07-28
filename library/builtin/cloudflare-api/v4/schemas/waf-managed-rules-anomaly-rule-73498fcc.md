---
title: waf-managed-rules_anomaly_rule
page_id: schema-waf-managed-rules-anomaly-rule-73498fcc
path: schemas
description: When triggered, anomaly detection WAF rules contribute to an overall threat score that will determine if a request is considered malicious. You can configure the total scoring threshold through the 'sensitivity' property of the WAF package.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# waf-managed-rules_anomaly_rule

When triggered, anomaly detection WAF rules contribute to an overall threat score that will determine if a request is considered malicious. You can configure the total scoring threshold through the 'sensitivity' property of the WAF package.

```yaml
{"description": "When triggered, anomaly detection WAF rules contribute to an overall threat score that will determine if a request is considered malicious. You can configure the total scoring threshold through the 'sensitivity' property of the WAF package.", "type": "object", "allOf": [{"$ref": "#/components/schemas/waf-managed-rules_schemas-base"}, {"properties": {"allowed_modes": {"$ref": "#/components/schemas/waf-managed-rules_allowed_modes_anomaly"}, "mode": {"$ref": "#/components/schemas/waf-managed-rules_mode_anomaly"}}, "type": "object"}], "required": ["id", "description", "priority", "allowed_modes", "mode", "group", "package_id"], "title": "Anomaly detection WAF rule"}
```
