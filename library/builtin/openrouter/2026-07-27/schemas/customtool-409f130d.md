---
title: CustomTool
page_id: schema-customtool-409f130d
path: schemas
description: Custom tool configuration
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# CustomTool

Custom tool configuration

```yaml
{"description": "Custom tool configuration", "example": {"name": "my_tool", "type": "custom"}, "properties": {"description": {"type": "string"}, "format": {"anyOf": [{"properties": {"type": {"enum": ["text"], "type": "string"}}, "required": ["type"], "type": "object"}, {"properties": {"definition": {"type": "string"}, "syntax": {"enum": ["lark", "regex"], "type": "string", "x-speakeasy-unknown-values": "allow"}, "type": {"enum": ["grammar"], "type": "string"}}, "required": ["type", "definition", "syntax"], "type": "object"}]}, "name": {"type": "string"}, "type": {"enum": ["custom"], "type": "string"}}, "required": ["type", "name"], "type": "object"}
```
