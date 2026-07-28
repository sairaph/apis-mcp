---
title: email_account_rules_plan_change
page_id: schema-email-account-rules-plan-change-9d33793e
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# email_account_rules_plan_change

```yaml
{"type": "object", "properties": {"remote": {"$ref": "#/components/schemas/email_account_rules_plan_remote_rule"}, "target": {"description": "Canonical recipient address or catch-all target.", "type": "string", "example": "support@example.com"}, "type": {"description": "Planned change type.", "type": "string", "example": "added", "enum": ["added", "updated", "deleted", "conflict"]}}}
```
