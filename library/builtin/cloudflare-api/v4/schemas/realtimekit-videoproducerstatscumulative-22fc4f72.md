---
title: realtimekit_VideoProducerStatsCumulative
page_id: schema-realtimekit-videoproducerstatscumulative-22fc4f72
path: schemas
description: Aggregated outbound (producer) video statistics for the session.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# realtimekit_VideoProducerStatsCumulative

Aggregated outbound (producer) video statistics for the session.

```yaml
{"description": "Aggregated outbound (producer) video statistics for the session.", "type": "object", "properties": {"frame_per_second": {"$ref": "#/components/schemas/realtimekit_PercentileStats"}, "frame_width": {"$ref": "#/components/schemas/realtimekit_PercentileStats"}, "high_negative_feedback_fraction": {"type": "number"}, "issues": {"type": "object", "properties": {"bandwidth_quality_limitation_fraction": {"type": "number"}, "cpu_quality_limitation_fraction": {"type": "number"}, "no_video_fraction": {"type": "number"}, "poor_resolution_fraction": {"type": "number"}, "quality_limitation_fraction": {"type": "number"}}}, "key_frames_encoded_fraction": {"type": "number"}, "packet_loss": {"$ref": "#/components/schemas/realtimekit_PacketLossCumulative"}, "quality_mos": {"$ref": "#/components/schemas/realtimekit_PercentileStats"}, "rtt": {"$ref": "#/components/schemas/realtimekit_LatencyCumulative"}}}
```
