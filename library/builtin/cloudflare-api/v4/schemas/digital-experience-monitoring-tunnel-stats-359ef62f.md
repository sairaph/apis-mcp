---
title: digital-experience-monitoring_tunnel_stats
page_id: schema-digital-experience-monitoring-tunnel-stats-359ef62f
path: schemas
description: WARP tunnel packet and byte counters.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# digital-experience-monitoring_tunnel_stats

WARP tunnel packet and byte counters.

```yaml
{"description": "WARP tunnel packet and byte counters.", "type": "object", "properties": {"bytesLost": {"description": "Number of bytes lost, split by direction.", "type": "object", "allOf": [{"$ref": "#/components/schemas/digital-experience-monitoring_directional_stat"}], "nullable": true}, "bytesReceived": {"description": "Number of bytes received, split by direction.", "type": "object", "allOf": [{"$ref": "#/components/schemas/digital-experience-monitoring_directional_stat"}], "nullable": true}, "bytesRetransmitted": {"description": "Number of bytes retransmitted, split by direction.", "type": "object", "allOf": [{"$ref": "#/components/schemas/digital-experience-monitoring_directional_stat"}], "nullable": true}, "bytesSent": {"description": "Number of bytes sent, split by direction.", "type": "object", "allOf": [{"$ref": "#/components/schemas/digital-experience-monitoring_directional_stat"}], "nullable": true}, "packetsLost": {"description": "Number of packets lost, split by direction.", "type": "object", "allOf": [{"$ref": "#/components/schemas/digital-experience-monitoring_directional_stat"}], "nullable": true}, "packetsReceived": {"description": "Number of packets received, split by direction.", "type": "object", "allOf": [{"$ref": "#/components/schemas/digital-experience-monitoring_directional_stat"}], "nullable": true}, "packetsRetransmitted": {"description": "Number of packets retransmitted, split by direction.", "type": "object", "allOf": [{"$ref": "#/components/schemas/digital-experience-monitoring_directional_stat"}], "nullable": true}, "packetsSent": {"description": "Number of packets sent, split by direction.", "type": "object", "allOf": [{"$ref": "#/components/schemas/digital-experience-monitoring_directional_stat"}], "nullable": true}, "statsWindowMs": {"description": "The measurement window duration in milliseconds.", "type": "integer", "nullable": true}}}
```
