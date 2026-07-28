---
title: digital-experience-monitoring_rtt_stats
page_id: schema-digital-experience-monitoring-rtt-stats-ebfd6502
path: schemas
description: Round-trip time statistics for the WARP tunnel.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# digital-experience-monitoring_rtt_stats

Round-trip time statistics for the WARP tunnel.

```yaml
{"description": "Round-trip time statistics for the WARP tunnel.", "type": "object", "properties": {"minRttUs": {"description": "Minimum round-trip time in microseconds.", "type": "object", "allOf": [{"$ref": "#/components/schemas/digital-experience-monitoring_directional_stat"}], "nullable": true}, "rttUs": {"description": "Round-trip time in microseconds.", "type": "object", "allOf": [{"$ref": "#/components/schemas/digital-experience-monitoring_directional_stat"}], "nullable": true}, "rttVarUs": {"description": "Round-trip time variance in microseconds.", "type": "object", "allOf": [{"$ref": "#/components/schemas/digital-experience-monitoring_directional_stat"}], "nullable": true}}}
```
