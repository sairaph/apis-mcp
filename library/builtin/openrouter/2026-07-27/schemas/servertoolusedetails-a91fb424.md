---
title: ServerToolUseDetails
page_id: schema-servertoolusedetails-a91fb424
path: schemas
description: Usage for server-side tool execution (e.g., web search)
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# ServerToolUseDetails

Usage for server-side tool execution (e.g., web search)

```yaml
{"description": "Usage for server-side tool execution (e.g., web search)", "example": {"tool_calls_executed": 2, "tool_calls_requested": 2, "web_search_requests": 2}, "properties": {"tool_calls_executed": {"description": "Number of OpenRouter server tool calls that executed and produced a result.", "type": ["integer", "null"]}, "tool_calls_requested": {"description": "Total number of OpenRouter server-orchestrated tool calls the model requested, across all tool types. Provider-native tools (e.g. native web search) are not counted here.", "type": ["integer", "null"]}, "web_search_requests": {"description": "Number of web searches performed by server-side tools. For server-orchestrated tool calls a web search is also counted in tool_calls_requested; provider-native web search may report web_search_requests only. Do not sum the two.", "type": ["integer", "null"]}}, "type": ["object", "null"]}
```
