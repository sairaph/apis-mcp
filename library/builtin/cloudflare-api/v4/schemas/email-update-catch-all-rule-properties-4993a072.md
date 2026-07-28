---
title: email_update_catch_all_rule_properties
page_id: schema-email-update-catch-all-rule-properties-4993a072
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# email_update_catch_all_rule_properties

```yaml
{"type": "object", "properties": {"actions": {"$ref": "#/components/schemas/email_rule_catchall-actions"}, "enabled": {"$ref": "#/components/schemas/email_rule_enabled"}, "matchers": {"$ref": "#/components/schemas/email_rule_catchall-matchers"}, "name": {"$ref": "#/components/schemas/email_rule_name"}, "owner_worker_tag": {"$ref": "#/components/schemas/email_rule_owner_worker_tag"}, "source": {"$ref": "#/components/schemas/email_rule_source"}}, "required": ["actions", "matchers"]}
```
