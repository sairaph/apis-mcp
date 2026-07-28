---
title: api-shield_Rule
page_id: schema-api-shield-rule-90f69562
path: schemas
description: A Token Validation rule that can enforce security policies using JWT Tokens.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# api-shield_Rule

A Token Validation rule that can enforce security policies using JWT Tokens.

```yaml
{"description": "A Token Validation rule that can enforce security policies using JWT Tokens.", "type": "object", "properties": {"action": {"$ref": "#/components/schemas/api-shield_action"}, "created_at": {"$ref": "#/components/schemas/api-shield_timestamp-2"}, "description": {"$ref": "#/components/schemas/api-shield_description-2"}, "enabled": {"$ref": "#/components/schemas/api-shield_enabled"}, "expression": {"$ref": "#/components/schemas/api-shield_expression"}, "id": {"$ref": "#/components/schemas/api-shield_uuid-2"}, "last_updated": {"$ref": "#/components/schemas/api-shield_timestamp-2"}, "selector": {"$ref": "#/components/schemas/api-shield_selector"}, "title": {"$ref": "#/components/schemas/api-shield_title-2"}}, "required": ["title", "description", "action", "enabled", "expression", "selector"]}
```
