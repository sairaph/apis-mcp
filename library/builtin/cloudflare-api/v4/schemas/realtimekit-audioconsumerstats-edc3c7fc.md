---
title: realtimekit_AudioConsumerStats
page_id: schema-realtimekit-audioconsumerstats-edc3c7fc
path: schemas
description: Per-sample inbound (consumer) audio statistics.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# realtimekit_AudioConsumerStats

Per-sample inbound (consumer) audio statistics.

```yaml
{"description": "Per-sample inbound (consumer) audio statistics.", "type": "object", "properties": {"bytes_received": {"type": "number"}, "concealment_events": {"type": "number"}, "consumer_id": {"type": "string"}, "jitter": {"type": "number"}, "jitter_buffer_delay": {"type": "number"}, "jitter_buffer_emitted_count": {"type": "number"}, "mid": {"type": "string"}, "mos_quality": {"type": "number"}, "packets_lost": {"type": "number"}, "packets_received": {"type": "number"}, "peer_id": {"type": "string"}, "producer_id": {"type": "string"}, "ssrc": {"type": "number"}, "timestamp": {"type": "string"}}}
```
