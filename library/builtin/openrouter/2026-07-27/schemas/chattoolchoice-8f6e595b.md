---
title: ChatToolChoice
page_id: schema-chattoolchoice-8f6e595b
path: schemas
description: Tool choice configuration
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ChatToolChoice

Tool choice configuration

```yaml
{"anyOf": [{"enum": ["none"], "type": "string"}, {"enum": ["auto"], "type": "string"}, {"enum": ["required"], "type": "string"}, {"$ref": "#/components/schemas/ChatNamedToolChoice"}, {"$ref": "#/components/schemas/ChatServerToolChoice"}], "description": "Tool choice configuration", "example": "auto"}
```
