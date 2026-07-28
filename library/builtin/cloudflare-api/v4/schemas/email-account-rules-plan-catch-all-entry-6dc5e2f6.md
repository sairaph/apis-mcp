---
title: email_account_rules_plan_catch_all_entry
page_id: schema-email-account-rules-plan-catch-all-entry-6dc5e2f6
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# email_account_rules_plan_catch_all_entry

```yaml
{"type": "object", "properties": {"rule": {"$ref": "#/components/schemas/email_account_rules_plan_catch_all_rule"}, "target": {"description": "Catch-all target to plan for, using the `*@domain` shape.", "type": "string", "example": "*@example.com"}}, "required": ["target", "rule"]}
```
