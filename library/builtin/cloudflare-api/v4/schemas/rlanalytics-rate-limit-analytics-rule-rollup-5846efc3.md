---
title: rlanalytics_rate_limit_analytics_rule_rollup
page_id: schema-rlanalytics-rate-limit-analytics-rule-rollup-5846efc3
path: schemas
description: Action totals and per-colo action counts for a particular rule.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# rlanalytics_rate_limit_analytics_rule_rollup

Action totals and per-colo action counts for a particular rule.

```yaml
{"description": "Action totals and per-colo action counts for a particular rule.", "type": "object", "properties": {"actions": {"$ref": "#/components/schemas/rlanalytics_rate_limit_analytics_action_counters"}, "colos": {"description": "Maps each colo name to its action counters.", "type": "object", "additionalProperties": {"$ref": "#/components/schemas/rlanalytics_rate_limit_analytics_colo_entry"}}, "versions": {"description": "Number of active versions of the rate limit rule.", "type": "integer", "example": 1}}, "required": ["versions", "actions", "colos"]}
```
