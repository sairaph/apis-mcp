---
title: realtimekit_PeerReportEvent
page_id: schema-realtimekit-peerreportevent-11a85f1b
path: schemas
description: A timestamped event recorded during the participant's connection.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# realtimekit_PeerReportEvent

A timestamped event recorded during the participant's connection.

```yaml
{"description": "A timestamped event recorded during the participant's connection.", "type": "object", "properties": {"metadata": {"description": "Event-specific metadata. Keys vary per event; values are primitive scalars (string, number, boolean, or null).", "type": "object", "additionalProperties": {"nullable": true, "oneOf": [{"type": "string"}, {"type": "number"}, {"type": "boolean"}]}}, "name": {"description": "Name of the event.", "type": "string"}, "timestamp": {"description": "Timestamp when the event occurred.", "type": "string"}}}
```
