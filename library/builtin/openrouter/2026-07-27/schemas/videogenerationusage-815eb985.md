---
title: VideoGenerationUsage
page_id: schema-videogenerationusage-815eb985
path: schemas
description: Usage and cost information for the video generation. Available once the job has completed.
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# VideoGenerationUsage

Usage and cost information for the video generation. Available once the job has completed.

```yaml
{"description": "Usage and cost information for the video generation. Available once the job has completed.", "example": {"cost": 0.5, "is_byok": false}, "properties": {"cost": {"description": "The cost of the video generation in USD.", "format": "double", "type": ["number", "null"]}, "is_byok": {"description": "Whether the request was made using a Bring Your Own Key configuration.", "type": "boolean"}}, "type": "object"}
```
