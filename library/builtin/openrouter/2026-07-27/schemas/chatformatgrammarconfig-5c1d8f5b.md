---
title: ChatFormatGrammarConfig
page_id: schema-chatformatgrammarconfig-5c1d8f5b
path: schemas
description: Custom grammar response format
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ChatFormatGrammarConfig

Custom grammar response format

```yaml
{"description": "Custom grammar response format", "example": {"grammar": "root ::= \"yes\" | \"no\"", "type": "grammar"}, "properties": {"grammar": {"description": "Custom grammar for text generation", "example": "root ::= \"yes\" | \"no\"", "type": "string"}, "type": {"enum": ["grammar"], "type": "string"}}, "required": ["type", "grammar"], "type": "object"}
```
