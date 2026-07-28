---
title: VideoGenerationResponse
page_id: schema-videogenerationresponse-1a49fa9a
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# VideoGenerationResponse

```yaml
{"example": {"generation_id": "gen-xyz789", "id": "job-abc123", "polling_url": "/api/v1/videos/job-abc123", "status": "pending"}, "properties": {"error": {"type": "string"}, "generation_id": {"description": "The generation ID associated with this video generation job. Available once the job has been processed.", "type": "string"}, "id": {"type": "string"}, "polling_url": {"type": "string"}, "status": {"enum": ["pending", "in_progress", "completed", "failed", "cancelled", "expired"], "type": "string", "x-speakeasy-unknown-values": "allow"}, "unsigned_urls": {"items": {"type": "string"}, "type": "array"}, "usage": {"$ref": "#/components/schemas/VideoGenerationUsage"}}, "required": ["id", "polling_url", "status"], "type": "object"}
```
