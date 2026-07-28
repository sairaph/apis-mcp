---
title: OutputAdvisorServerToolItem
page_id: schema-outputadvisorservertoolitem-bd0e1daa
path: schemas
description: An openrouter:advisor server tool output item
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# OutputAdvisorServerToolItem

An openrouter:advisor server tool output item

```yaml
{"description": "An openrouter:advisor server tool output item", "example": {"id": "st_tmp_abc123", "status": "completed", "type": "openrouter:advisor"}, "properties": {"advice": {"description": "The advisor model's response (the advice text returned to the executor).", "type": "string"}, "error": {"description": "Error message when the advisor call did not produce advice.", "type": "string"}, "id": {"type": "string"}, "instance_name": {"description": "Provider-safe function name of the specific advisor instance that produced this item (e.g. `openrouter_advisor__1`). Present only when more than one advisor tool is configured; omitted for the default single advisor. Echo this field back unchanged so the advisor's cross-request memory stays namespaced to the correct instance. This identity is positional: it is derived from the index of the advisor entry in the request `tools` array, so clients must keep the order of advisor tool entries stable across requests in a conversation. Reordering or inserting advisor entries shifts these names and causes each advisor's cross-request memory to be attributed to the wrong instance.", "example": "openrouter_advisor__1", "type": "string"}, "model": {"description": "Slug of the advisor model that was consulted.", "type": "string"}, "prompt": {"description": "The prompt the executor sent to the advisor.", "type": "string"}, "status": {"$ref": "#/components/schemas/ToolCallStatus"}, "type": {"enum": ["openrouter:advisor"], "type": "string"}}, "required": ["status", "type"], "type": "object"}
```
