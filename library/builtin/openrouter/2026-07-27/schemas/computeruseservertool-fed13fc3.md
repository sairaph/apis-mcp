---
title: ComputerUseServerTool
page_id: schema-computeruseservertool-fed13fc3
path: schemas
description: Computer use preview tool configuration
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ComputerUseServerTool

Computer use preview tool configuration

```yaml
{"description": "Computer use preview tool configuration", "example": {"display_height": 768, "display_width": 1024, "environment": "linux", "type": "computer_use_preview"}, "properties": {"display_height": {"type": "integer"}, "display_width": {"type": "integer"}, "environment": {"enum": ["windows", "mac", "linux", "ubuntu", "browser"], "type": "string", "x-speakeasy-unknown-values": "allow"}, "type": {"enum": ["computer_use_preview"], "type": "string"}}, "required": ["type", "display_height", "display_width", "environment"], "type": "object"}
```
