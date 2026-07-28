---
title: FusionSource
page_id: schema-fusionsource-21a3887e
path: schemas
description: A web page retrieved via web search during a fusion run.
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# FusionSource

A web page retrieved via web search during a fusion run.

```yaml
{"description": "A web page retrieved via web search during a fusion run.", "example": {"title": "Example article title", "url": "https://example.com/article"}, "properties": {"title": {"description": "Title of the retrieved web page.", "type": "string"}, "url": {"description": "URL of the web page a panel or the judge retrieved during the run.", "type": "string"}}, "required": ["url", "title"], "type": "object"}
```
