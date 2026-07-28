---
title: CodeInterpreterServerTool
page_id: schema-codeinterpreterservertool-53b1f3d0
path: schemas
description: Code interpreter tool configuration
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# CodeInterpreterServerTool

Code interpreter tool configuration

```yaml
{"description": "Code interpreter tool configuration", "example": {"container": "auto", "type": "code_interpreter"}, "properties": {"container": {"anyOf": [{"type": "string"}, {"properties": {"file_ids": {"items": {"type": "string"}, "type": "array"}, "memory_limit": {"enum": ["1g", "4g", "16g", "64g", null], "type": ["string", "null"], "x-speakeasy-unknown-values": "allow"}, "type": {"enum": ["auto"], "type": "string"}}, "required": ["type"], "type": "object"}]}, "type": {"enum": ["code_interpreter"], "type": "string"}}, "required": ["type", "container"], "type": "object"}
```
