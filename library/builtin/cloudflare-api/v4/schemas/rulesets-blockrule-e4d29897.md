---
title: rulesets_BlockRule
page_id: schema-rulesets-blockrule-e4d29897
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# rulesets_BlockRule

```yaml
{"allOf": [{"$ref": "#/components/schemas/rulesets_Rule"}, {"properties": {"action": {"enum": ["block"]}, "action_parameters": {"properties": {"response": {"description": "The response to show when the block is applied.", "type": "object", "properties": {"content": {"description": "The content to return.", "example": "{\n  \"success\": false,\n  \"error\": \"you have been blocked\"\n}", "minLength": 1, "title": "Content", "type": "string"}, "content_type": {"description": "The type of the content to return.", "type": "string", "example": "application/json", "minLength": 1, "title": "Content Type"}, "status_code": {"description": "The status code to return.", "type": "integer", "maximum": 499, "minimum": 400, "title": "Status Code"}}, "required": ["status_code", "content", "content_type"], "title": "Response"}}}, "description": {"example": "Block the request."}}, "title": "Block Rule"}]}
```
