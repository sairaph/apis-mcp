---
title: mconn_recorded_event
page_id: schema-mconn-recorded-event-0424ab53
path: schemas
description: Recorded Event
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# mconn_recorded_event

Recorded Event

```yaml
{"description": "Recorded Event", "type": "object", "properties": {"e": {"$ref": "#/components/schemas/mconn_event"}, "n": {"description": "Sequence number, used to order events with the same timestamp", "type": "number"}, "t": {"description": "Time the Event was recorded (seconds since the Unix epoch)", "type": "number"}, "v": {"description": "Version", "type": "string"}}, "required": ["t", "n", "e"]}
```
