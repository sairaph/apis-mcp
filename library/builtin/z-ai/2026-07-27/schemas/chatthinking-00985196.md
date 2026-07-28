---
title: ChatThinking
page_id: schema-chatthinking-00985196
path: schemas
description: Only supported by GLM-4.5 series and higher models. This parameter is used to control whether the model enable the chain of thought.
source: https://docs.z.ai/openapi.json
source_type: openapi
imported_from: https://docs.z.ai/openapi.json
---

# ChatThinking

Only supported by GLM-4.5 series and higher models. This parameter is used to control whether the model enable the chain of thought.

```yaml
{"type": "object", "description": "Only supported by GLM-4.5 series and higher models. This parameter is used to control whether the model enable the chain of thought.", "properties": {"type": {"type": "string", "description": "Whether to enable the chain of thought(When enabled, GLM-5.2 GLM-5.1 GLM-5 GLM-5-Turbo GLM-5V-Turbo GLM-4.6 GLM-4.5 and others will automatically determine whether to think, while GLM-4.7 and GLM-4.5V will think compulsorily), default: enabled", "default": "enabled", "enum": ["enabled", "disabled"]}, "clear_thinking": {"type": "boolean", "description": "Default value is True. Controls whether to clear `reasoning_content` from previous conversation turns. View more in [Thinking Mode](/guides/capabilities/thinking-mode). \n - `true` (default): For this request, the system ignores/removes `reasoning_content` from prior turns, and only keeps non-reasoning context (e.g., user/assistant visible text, tool calls, and tool results). This is recommended for general chat or lightweight tasks to reduce context length and cost. \n - `false`: Retains `reasoning_content` from prior turns and includes it in the context sent to the model. To enable Preserved Thinking, you must forward the full, unmodified, and correctly ordered historical `reasoning_content` in `messages`. Missing, truncated, rewritten, or reordered blocks may degrade performance or prevent the feature from taking effect. \n - Notes: This parameter only affects cross-turn historical thinking blocks; it does not change whether the model generates/returns thinking in the current turn.", "default": true, "example": true}}}
```
