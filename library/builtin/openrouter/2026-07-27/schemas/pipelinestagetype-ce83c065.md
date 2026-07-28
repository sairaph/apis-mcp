---
title: PipelineStageType
page_id: schema-pipelinestagetype-ce83c065
path: schemas
description: Categorical kind of a pipeline stage. Multiple plugins can share a type (e.g. all guardrail-level plugins emit `guardrail`); the `name` field disambiguates which plugin emitted it.
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# PipelineStageType

Categorical kind of a pipeline stage. Multiple plugins can share a type (e.g. all guardrail-level plugins emit `guardrail`); the `name` field disambiguates which plugin emitted it.

```yaml
{"description": "Categorical kind of a pipeline stage. Multiple plugins can share a type (e.g. all guardrail-level plugins emit `guardrail`); the `name` field disambiguates which plugin emitted it.", "enum": ["guardrail", "plugin", "server_tools", "response_healing", "context_compression"], "example": "guardrail", "type": "string", "x-speakeasy-unknown-values": "allow"}
```
