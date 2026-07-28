---
title: realtimekit_AudioProducerStats
page_id: schema-realtimekit-audioproducerstats-529c07d3
path: schemas
description: Per-sample outbound (producer) audio statistics.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# realtimekit_AudioProducerStats

Per-sample outbound (producer) audio statistics.

```yaml
{"description": "Per-sample outbound (producer) audio statistics.", "type": "object", "properties": {"bytes_sent": {"type": "number"}, "jitter": {"type": "number"}, "mid": {"type": "string"}, "mos_quality": {"type": "number"}, "packets_lost": {"type": "number"}, "packets_sent": {"type": "number"}, "producer_id": {"type": "string"}, "rtt": {"type": "number"}, "ssrc": {"type": "number"}, "timestamp": {"type": "string"}}}
```
