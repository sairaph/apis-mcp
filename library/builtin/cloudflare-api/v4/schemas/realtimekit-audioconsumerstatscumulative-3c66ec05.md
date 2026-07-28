---
title: realtimekit_AudioConsumerStatsCumulative
page_id: schema-realtimekit-audioconsumerstatscumulative-3c66ec05
path: schemas
description: Aggregated inbound (consumer) audio statistics for the session.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# realtimekit_AudioConsumerStatsCumulative

Aggregated inbound (consumer) audio statistics for the session.

```yaml
{"description": "Aggregated inbound (consumer) audio statistics for the session.", "type": "object", "properties": {"jitter_buffer_delay": {"$ref": "#/components/schemas/realtimekit_LatencyCumulative"}, "packet_loss": {"$ref": "#/components/schemas/realtimekit_PacketLossCumulative"}, "quality_mos": {"$ref": "#/components/schemas/realtimekit_PercentileStats"}}}
```
