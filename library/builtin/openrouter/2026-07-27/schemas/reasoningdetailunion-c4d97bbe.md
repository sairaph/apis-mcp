---
title: ReasoningDetailUnion
page_id: schema-reasoningdetailunion-c4d97bbe
path: schemas
description: Reasoning detail union schema
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ReasoningDetailUnion

Reasoning detail union schema

```yaml
{"description": "Reasoning detail union schema", "discriminator": {"mapping": {"reasoning.encrypted": "#/components/schemas/ReasoningDetailEncrypted", "reasoning.server_tool_call": "#/components/schemas/ReasoningDetailServerToolCall", "reasoning.summary": "#/components/schemas/ReasoningDetailSummary", "reasoning.text": "#/components/schemas/ReasoningDetailText"}, "propertyName": "type"}, "example": {"summary": "The model analyzed the problem by first identifying key constraints, then evaluating possible solutions...", "type": "reasoning.summary"}, "oneOf": [{"$ref": "#/components/schemas/ReasoningDetailSummary"}, {"$ref": "#/components/schemas/ReasoningDetailEncrypted"}, {"$ref": "#/components/schemas/ReasoningDetailText"}, {"$ref": "#/components/schemas/ReasoningDetailServerToolCall"}]}
```
