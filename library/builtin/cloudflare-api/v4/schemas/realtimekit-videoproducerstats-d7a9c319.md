---
title: realtimekit_VideoProducerStats
page_id: schema-realtimekit-videoproducerstats-d7a9c319
path: schemas
description: Per-sample outbound (producer) video statistics.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# realtimekit_VideoProducerStats

Per-sample outbound (producer) video statistics.

```yaml
{"description": "Per-sample outbound (producer) video statistics.", "type": "object", "properties": {"bytes_sent": {"type": "number"}, "fir_count": {"type": "number"}, "frame_height": {"type": "number"}, "frame_width": {"type": "number"}, "frames_encoded": {"type": "number"}, "frames_per_second": {"type": "number"}, "jitter": {"type": "number"}, "key_frames_encoded": {"type": "number"}, "mid": {"type": "string"}, "mos_quality": {"type": "number"}, "packets_lost": {"type": "number"}, "packets_sent": {"type": "number"}, "pli_count": {"type": "number"}, "producer_id": {"type": "string"}, "quality_limitation_durations": {"type": "object", "properties": {"bandwidth": {"type": "number"}, "cpu": {"type": "number"}, "none": {"type": "number"}, "other": {"type": "number"}}}, "quality_limitation_reason": {"type": "string", "enum": ["cpu", "bandwidth", "none", "other"]}, "quality_limitation_resolution_changes": {"type": "number"}, "rtt": {"type": "number"}, "ssrc": {"type": "number"}, "timestamp": {"type": "string"}}}
```
