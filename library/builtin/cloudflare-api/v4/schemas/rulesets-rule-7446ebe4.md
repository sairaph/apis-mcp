---
title: rulesets_Rule
page_id: schema-rulesets-rule-7446ebe4
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# rulesets_Rule

```yaml
{"type": "object", "properties": {"action": {"$ref": "#/components/schemas/rulesets_RuleAction"}, "action_parameters": {"description": "The parameters configuring the rule's action.", "type": "object", "default": {}, "title": "Action Parameters"}, "categories": {"$ref": "#/components/schemas/rulesets_RuleCategories"}, "description": {"description": "An informative description of the rule.", "type": "string", "default": "", "title": "Description"}, "enabled": {"allOf": [{"$ref": "#/components/schemas/rulesets_RuleEnabled"}, {"default": true}]}, "exposed_credential_check": {"$ref": "#/components/schemas/rulesets_RuleExposedCredentialCheck"}, "expression": {"description": "The expression defining which traffic will match the rule.", "type": "string", "example": "ip.src eq 1.1.1.1", "minLength": 1, "title": "Expression"}, "id": {"$ref": "#/components/schemas/rulesets_RuleId"}, "last_updated": {"description": "The timestamp of when the rule was last modified.", "type": "string", "format": "date-time", "example": "2000-01-01T00:00:00.000000Z", "readOnly": true, "title": "Last Updated", "x-stainless-skip": ["terraform"]}, "logging": {"$ref": "#/components/schemas/rulesets_RuleLogging"}, "ratelimit": {"$ref": "#/components/schemas/rulesets_RuleRatelimit"}, "ref": {"description": "The reference of the rule (the rule's ID by default).", "type": "string", "example": "my_ref", "minLength": 1, "title": "Ref"}, "version": {"description": "The version of the rule.", "type": "string", "example": "1", "pattern": "^[0-9]+$", "readOnly": true, "title": "Version", "x-stainless-skip": ["terraform"]}}, "required": ["version", "last_updated"], "title": "Rule"}
```
