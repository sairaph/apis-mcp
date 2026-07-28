---
title: mconn_event_metadata
page_id: schema-mconn-event-metadata-f03e14a0
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# mconn_event_metadata

```yaml
{"type": "object", "properties": {"a": {"description": "Time the Event was collected (seconds since the Unix epoch)", "type": "number"}, "k": {"description": "Kind", "type": "string"}, "n": {"description": "Sequence number, used to order events with the same timestamp", "type": "number"}, "t": {"description": "Time the Event was recorded (seconds since the Unix epoch)", "type": "number"}}, "required": ["a", "t", "n", "k"]}
```
