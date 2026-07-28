---
title: VideoModelsListResponse
page_id: schema-videomodelslistresponse-01918413
path: schemas
source: https://openrouter.ai/openapi.json
source_type: openapi
imported_from: https://openrouter.ai/openapi.json
---

# VideoModelsListResponse

```yaml
{"example": {"data": [{"allowed_passthrough_parameters": [], "canonical_slug": "google/veo-3.1", "created": 1700000000, "description": "Google video generation model", "generate_audio": true, "id": "google/veo-3.1", "name": "Veo 3.1", "pricing_skus": {"generate": "0.50"}, "seed": null, "supported_aspect_ratios": ["16:9"], "supported_durations": [5, 8], "supported_frame_images": ["first_frame", "last_frame"], "supported_resolutions": ["720p"], "supported_sizes": null}]}, "properties": {"data": {"items": {"$ref": "#/components/schemas/VideoModel"}, "type": "array"}}, "required": ["data"], "type": "object"}
```
