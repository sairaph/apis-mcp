---
title: cloudforce-one_StatsResponse
page_id: schema-cloudforce-one-statsresponse-dc5c017f
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cloudforce-one_StatsResponse

```yaml
{"type": "object", "properties": {"pending_approvals": {"type": "number", "example": 5}, "rules_by_namespace": {"type": "object", "example": {"yara/dns_record": 12, "yara/workers": 30}, "additionalProperties": {"type": "number"}}, "total_rules": {"type": "number", "example": 42}}, "required": ["total_rules", "rules_by_namespace", "pending_approvals"]}
```
