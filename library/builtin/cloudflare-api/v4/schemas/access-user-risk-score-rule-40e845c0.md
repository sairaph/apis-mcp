---
title: access_user_risk_score_rule
page_id: schema-access-user-risk-score-rule-40e845c0
path: schemas
description: Matches a user's risk score.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_user_risk_score_rule

Matches a user's risk score.

```yaml
{"description": "Matches a user's risk score.", "type": "object", "properties": {"user_risk_score": {"type": "object", "properties": {"user_risk_score": {"description": "A list of risk score levels to match. Values can be low, medium, high, or unscored.", "type": "array", "items": {"enum": ["low", "medium", "high", "unscored"], "type": "string"}, "example": ["low", "medium"], "minItems": 1}}, "required": ["user_risk_score"]}}, "required": ["user_risk_score"], "title": "User Risk Score"}
```
