---
title: UnifiedBenchmarksResponse
page_id: schema-unifiedbenchmarksresponse-65132901
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# UnifiedBenchmarksResponse

```yaml
{"example": {"data": [{"agentic_index": 58.3, "coding_index": 65.8, "display_name": "GPT-4o", "intelligence_index": 71.2, "model_permaslug": "openai/gpt-4o", "pricing": {"completion": "0.00001", "prompt": "0.0000025"}, "source": "artificial-analysis"}], "meta": {"as_of": "2026-06-03T12:00:00Z", "citation": null, "model_count": 1, "source": null, "source_url": null, "task_type": null, "version": "v1"}}, "properties": {"data": {"items": {"discriminator": {"mapping": {"artificial-analysis": "#/components/schemas/UnifiedBenchmarksAAItem", "design-arena": "#/components/schemas/UnifiedBenchmarksDAItem"}, "propertyName": "source"}, "oneOf": [{"$ref": "#/components/schemas/UnifiedBenchmarksAAItem"}, {"$ref": "#/components/schemas/UnifiedBenchmarksDAItem"}]}, "type": "array"}, "meta": {"$ref": "#/components/schemas/UnifiedBenchmarksMeta"}}, "required": ["data", "meta"], "type": "object"}
```
