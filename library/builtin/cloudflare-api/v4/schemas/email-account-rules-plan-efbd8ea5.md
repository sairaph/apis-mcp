---
title: email_account_rules_plan
page_id: schema-email-account-rules-plan-efbd8ea5
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# email_account_rules_plan

```yaml
{"type": "object", "properties": {"zones": {"type": "array", "items": {"$ref": "#/components/schemas/email_account_rules_plan_zone"}}}, "example": {"zones": [{"changes": [{"target": "support@example.com", "type": "added"}, {"remote": {"actions": [{"type": "worker", "value": ["my-worker"]}], "enabled": false, "id": "f1e2d3c4b5a6978869708192a3b4c5d6", "matchers": [{"field": "to", "type": "literal", "value": "info@example.com"}], "source": "wrangler"}, "target": "info@example.com", "type": "updated"}, {"remote": {"actions": [{"type": "worker", "value": ["my-worker"]}], "enabled": true, "id": "0a1b2c3d4e5f60718293a4b5c6d7e8f9", "matchers": [{"field": "to", "type": "literal", "value": "old@example.com"}], "source": "wrangler"}, "target": "old@example.com", "type": "deleted"}, {"remote": {"actions": [{"type": "forward", "value": ["billing-team@example.net"]}], "enabled": true, "id": "9f8e7d6c5b4a39281706f5e4d3c2b1a0", "matchers": [{"field": "to", "type": "literal", "value": "billing@example.com"}], "source": "api"}, "target": "billing@example.com", "type": "conflict"}, {"remote": {"actions": [{"type": "worker", "value": ["other-worker"]}], "enabled": true, "id": "3c4d5e6f70819203a4b5c6d7e8f90a1b", "matchers": [{"type": "all"}], "owner_worker_name": "other-worker", "source": "wrangler"}, "target": "*@example.com", "type": "conflict"}], "zone_id": "023e105f4ecef8ad9ca31a8372d0c353", "zone_name": "example.com"}]}}
```
