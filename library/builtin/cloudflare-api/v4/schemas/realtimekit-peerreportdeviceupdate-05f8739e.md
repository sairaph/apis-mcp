---
title: realtimekit_PeerReportDeviceUpdate
page_id: schema-realtimekit-peerreportdeviceupdate-05f8739e
path: schemas
description: A change to the set of available devices at a point in time.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# realtimekit_PeerReportDeviceUpdate

A change to the set of available devices at a point in time.

```yaml
{"description": "A change to the set of available devices at a point in time.", "type": "object", "properties": {"added": {"description": "Devices that became available.", "type": "array", "items": {"$ref": "#/components/schemas/realtimekit_PeerReportDevice"}}, "removed": {"description": "Devices that became unavailable.", "type": "array", "items": {"$ref": "#/components/schemas/realtimekit_PeerReportDevice"}}, "timestamp": {"description": "Timestamp of the device update.", "type": "string"}}}
```
