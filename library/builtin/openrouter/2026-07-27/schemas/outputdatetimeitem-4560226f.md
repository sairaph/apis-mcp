---
title: OutputDatetimeItem
page_id: schema-outputdatetimeitem-4560226f
path: schemas
description: An openrouter:datetime server tool output item
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# OutputDatetimeItem

An openrouter:datetime server tool output item

```yaml
{"description": "An openrouter:datetime server tool output item", "example": {"datetime": "2026-03-12T14:30:00.000Z", "id": "dt_tmp_abc123", "status": "completed", "timezone": "UTC", "type": "openrouter:datetime"}, "properties": {"datetime": {"description": "ISO 8601 datetime string", "type": "string"}, "id": {"type": "string"}, "status": {"$ref": "#/components/schemas/ToolCallStatus"}, "timezone": {"description": "IANA timezone name", "type": "string"}, "type": {"enum": ["openrouter:datetime"], "type": "string"}}, "required": ["status", "type", "datetime", "timezone"], "type": "object"}
```
