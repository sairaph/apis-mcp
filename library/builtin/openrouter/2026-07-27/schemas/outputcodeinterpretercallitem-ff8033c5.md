---
title: OutputCodeInterpreterCallItem
page_id: schema-outputcodeinterpretercallitem-ff8033c5
path: schemas
description: A code interpreter execution call with outputs
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# OutputCodeInterpreterCallItem

A code interpreter execution call with outputs

```yaml
{"allOf": [{"$ref": "#/components/schemas/CodeInterpreterCallItem"}, {"properties": {}, "type": "object"}], "description": "A code interpreter execution call with outputs", "example": {"code": "print(\"hello\")", "container_id": "ctr-xyz789", "id": "ci-abc123", "outputs": [{"logs": "hello\n", "type": "logs"}], "status": "completed", "type": "code_interpreter_call"}}
```
