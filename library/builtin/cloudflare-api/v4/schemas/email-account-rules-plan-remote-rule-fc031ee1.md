---
title: email_account_rules_plan_remote_rule
page_id: schema-email-account-rules-plan-remote-rule-fc031ee1
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# email_account_rules_plan_remote_rule

```yaml
{"type": "object", "properties": {"actions": {"$ref": "#/components/schemas/email_rule_actions"}, "enabled": {"$ref": "#/components/schemas/email_rule_enabled"}, "id": {"$ref": "#/components/schemas/email_rule_identifier"}, "matchers": {"$ref": "#/components/schemas/email_rule_matchers"}, "owner_worker_name": {"description": "Name of the Wrangler-owned Worker that owns the remote rule, when available for display.", "type": "string", "example": "email-processor"}, "source": {"$ref": "#/components/schemas/email_rule_source"}}}
```
