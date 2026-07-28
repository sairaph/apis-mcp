---
title: realtimekit_VideoConsumerStatsCumulative
page_id: schema-realtimekit-videoconsumerstatscumulative-24e7a5f6
path: schemas
description: Aggregated inbound (consumer) video statistics for the session.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# realtimekit_VideoConsumerStatsCumulative

Aggregated inbound (consumer) video statistics for the session.

```yaml
{"description": "Aggregated inbound (consumer) video statistics for the session.", "type": "object", "properties": {"frame_per_second": {"$ref": "#/components/schemas/realtimekit_PercentileStats"}, "frame_width": {"$ref": "#/components/schemas/realtimekit_PercentileStats"}, "issues": {"type": "object", "properties": {"lag_fraction": {"type": "number"}, "no_video_fraction": {"type": "number"}, "poor_resolution_fraction": {"type": "number"}}}, "jitter_buffer_delay": {"$ref": "#/components/schemas/realtimekit_LatencyCumulative"}, "key_frames_decoded_fraction": {"type": "number"}, "packet_loss": {"$ref": "#/components/schemas/realtimekit_PacketLossCumulative"}, "quality_mos": {"$ref": "#/components/schemas/realtimekit_PercentileStats"}}}
```
