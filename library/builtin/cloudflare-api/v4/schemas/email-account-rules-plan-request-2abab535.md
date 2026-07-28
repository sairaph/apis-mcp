---
title: email_account_rules_plan_request
page_id: schema-email-account-rules-plan-request-2abab535
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# email_account_rules_plan_request

```yaml
{"type": "object", "properties": {"catch_all_rules": {"description": "Desired catch-all Email Routing rules managed by the deploying Worker.", "type": "array", "items": {"$ref": "#/components/schemas/email_account_rules_plan_catch_all_entry"}}, "owner_worker_tag": {"$ref": "#/components/schemas/email_rule_owner_worker_tag"}, "rules": {"description": "Desired normal Email Routing rules managed by the deploying Worker.", "type": "array", "items": {"$ref": "#/components/schemas/email_account_rules_plan_rule"}}}, "example": {"catch_all_rules": [{"rule": {"actions": [{"type": "worker", "value": ["my-worker"]}], "matchers": [{"type": "all"}]}, "target": "*@example.com"}], "owner_worker_tag": "a7e6fb77503c41d8a7f3113c6918f10c", "rules": [{"actions": [{"type": "worker", "value": ["my-worker"]}], "matchers": [{"field": "to", "type": "literal", "value": "support@example.com"}]}, {"actions": [{"type": "worker", "value": ["my-worker"]}], "matchers": [{"field": "to", "type": "literal", "value": "info@example.com"}]}, {"actions": [{"type": "worker", "value": ["my-worker"]}], "matchers": [{"field": "to", "type": "literal", "value": "billing@example.com"}]}]}, "required": ["owner_worker_tag"]}
```
