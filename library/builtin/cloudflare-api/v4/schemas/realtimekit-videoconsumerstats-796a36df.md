---
title: realtimekit_VideoConsumerStats
page_id: schema-realtimekit-videoconsumerstats-796a36df
path: schemas
description: Per-sample inbound (consumer) video statistics.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# realtimekit_VideoConsumerStats

Per-sample inbound (consumer) video statistics.

```yaml
{"description": "Per-sample inbound (consumer) video statistics.", "type": "object", "properties": {"bytes_received": {"type": "number"}, "consumer_id": {"type": "string"}, "fir_count": {"type": "number"}, "frame_height": {"type": "number"}, "frame_width": {"type": "number"}, "frames_decoded": {"type": "number"}, "frames_dropped": {"type": "number"}, "frames_per_second": {"type": "number"}, "jitter": {"type": "number"}, "jitter_buffer_delay": {"type": "number"}, "jitter_buffer_emitted_count": {"type": "number"}, "key_frames_decoded": {"type": "number"}, "mid": {"type": "string"}, "mos_quality": {"type": "number"}, "packets_lost": {"type": "number"}, "packets_received": {"type": "number"}, "peer_id": {"type": "string"}, "producer_id": {"type": "string"}, "ssrc": {"type": "number"}, "timestamp": {"type": "string"}}}
```
