---
title: DatetimeServerTool
page_id: schema-datetimeservertool-8be47695
path: schemas
description: 'OpenRouter built-in server tool: returns the current date and time'
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# DatetimeServerTool

OpenRouter built-in server tool: returns the current date and time

```yaml
{"description": "OpenRouter built-in server tool: returns the current date and time", "example": {"parameters": {"timezone": "America/New_York"}, "type": "openrouter:datetime"}, "properties": {"parameters": {"$ref": "#/components/schemas/DatetimeServerToolConfig"}, "type": {"enum": ["openrouter:datetime"], "type": "string"}}, "required": ["type"], "type": "object"}
```
