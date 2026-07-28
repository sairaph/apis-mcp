---
title: MessagesOutputConfig
page_id: schema-messagesoutputconfig-ae403b90
path: schemas
description: Configuration for controlling output behavior. Supports the effort parameter and structured output format.
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# MessagesOutputConfig

Configuration for controlling output behavior. Supports the effort parameter and structured output format.

```yaml
{"description": "Configuration for controlling output behavior. Supports the effort parameter and structured output format.", "example": {"effort": "medium"}, "properties": {"effort": {"description": "How much effort the model should put into its response. Higher effort levels may result in more thorough analysis but take longer. Valid values are `low`, `medium`, `high`, `xhigh`, or `max`.", "enum": ["low", "medium", "high", "xhigh", "max", null], "example": "medium", "type": ["string", "null"], "x-speakeasy-unknown-values": "allow"}, "format": {"description": "A schema to specify Claude's output format in responses. See [structured outputs](https://platform.claude.com/docs/en/build-with-claude/structured-outputs).", "properties": {"schema": {"additionalProperties": {}, "type": "object"}, "type": {"enum": ["json_schema"], "type": "string"}}, "required": ["type", "schema"], "type": ["object", "null"]}, "task_budget": {"description": "Task budget for an agentic turn. The model sees a countdown of remaining tokens and uses it to prioritize work and wind down gracefully. Advisory — does not enforce a hard cap.", "example": {"total": 400000, "type": "tokens"}, "properties": {"remaining": {"minimum": 0, "type": ["integer", "null"]}, "total": {"minimum": 20000, "type": "integer"}, "type": {"enum": ["tokens"], "type": "string"}}, "required": ["type", "total"], "type": ["object", "null"]}}, "type": "object"}
```
