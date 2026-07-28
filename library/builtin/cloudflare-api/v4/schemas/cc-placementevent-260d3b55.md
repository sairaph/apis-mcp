---
title: cc_PlacementEvent
page_id: schema-cc-placementevent-260d3b55
path: schemas
description: An event within a Placement or a Job
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cc_PlacementEvent

An event within a Placement or a Job

```yaml
{"description": "An event within a Placement or a Job", "type": "object", "properties": {"details": {"type": "object", "additionalProperties": true}, "id": {"type": "string"}, "message": {"type": "string"}, "name": {"$ref": "#/components/schemas/cc_EventName"}, "statusChange": {"type": "object", "additionalProperties": true}, "time": {"$ref": "#/components/schemas/cc_ISO8601Timestamp"}, "type": {"$ref": "#/components/schemas/cc_EventType"}}, "required": ["id", "time", "type", "name", "message", "details", "statusChange"]}
```
