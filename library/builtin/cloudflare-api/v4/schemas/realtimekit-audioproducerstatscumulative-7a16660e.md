---
title: realtimekit_AudioProducerStatsCumulative
page_id: schema-realtimekit-audioproducerstatscumulative-7a16660e
path: schemas
description: Aggregated outbound (producer) audio statistics for the session.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# realtimekit_AudioProducerStatsCumulative

Aggregated outbound (producer) audio statistics for the session.

```yaml
{"description": "Aggregated outbound (producer) audio statistics for the session.", "type": "object", "properties": {"packet_loss": {"$ref": "#/components/schemas/realtimekit_PacketLossCumulative"}, "quality_mos": {"$ref": "#/components/schemas/realtimekit_PercentileStats"}, "rtt": {"$ref": "#/components/schemas/realtimekit_LatencyCumulative"}}}
```
