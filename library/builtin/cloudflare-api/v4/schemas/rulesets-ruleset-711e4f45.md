---
title: rulesets_Ruleset
page_id: schema-rulesets-ruleset-711e4f45
path: schemas
description: A ruleset object.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# rulesets_Ruleset

A ruleset object.

```yaml
{"description": "A ruleset object.", "type": "object", "properties": {"description": {"description": "An informative description of the ruleset.", "type": "string", "example": "A description for my ruleset.", "default": "", "title": "Description"}, "id": {"allOf": [{"$ref": "#/components/schemas/rulesets_RulesetId"}, {"readOnly": true}]}, "last_updated": {"description": "The timestamp of when the ruleset was last modified.", "type": "string", "format": "date-time", "example": "2000-01-01T00:00:00.000000Z", "readOnly": true, "title": "Last Updated"}, "name": {"description": "The human-readable name of the ruleset.", "type": "string", "example": "My ruleset", "minLength": 1, "title": "Name"}, "version": {"allOf": [{"$ref": "#/components/schemas/rulesets_RulesetVersion"}, {"readOnly": true}]}}, "required": ["id", "version", "last_updated"], "title": "Ruleset"}
```
