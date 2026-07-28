---
title: realtimekit_TrackConfigLayer
page_id: schema-realtimekit-trackconfiglayer-d592d8d3
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# realtimekit_TrackConfigLayer

```yaml
{"type": "object", "properties": {"file_name_prefix": {"description": "A file name prefix to apply for files generated from this layer", "type": "string", "pattern": "^[-\\w\\s]+$"}, "media_kind": {"description": "Media kind to record. Track recording currently supports audio only.", "type": "string", "default": "audio", "enum": ["audio"]}}, "title": "TrackLayerConfig"}
```
