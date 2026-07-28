---
title: realtimekit_PeerReportMetadata
page_id: schema-realtimekit-peerreportmetadata-b0784fd0
path: schemas
description: Connection and device metadata for the participant.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# realtimekit_PeerReportMetadata

Connection and device metadata for the participant.

```yaml
{"description": "Connection and device metadata for the participant.", "type": "object", "properties": {"audio_devices_updates": {"type": "array", "items": {"$ref": "#/components/schemas/realtimekit_PeerReportDeviceUpdate"}}, "browser_metadata": {"type": "object", "properties": {"browser": {"type": "string"}, "browser_version": {"type": "string"}, "engine": {"type": "string"}, "user_agent": {"type": "string"}, "webgl_support": {"type": "boolean", "nullable": true}}}, "candidate_pairs": {"type": "object", "properties": {"consuming_transport": {"type": "array", "items": {"$ref": "#/components/schemas/realtimekit_CandidatePairStats"}}, "producing_transport": {"type": "array", "items": {"$ref": "#/components/schemas/realtimekit_CandidatePairStats"}}}}, "device_info": {"type": "object", "properties": {"cpus": {"type": "number"}, "is_mobile": {"type": "boolean"}, "os": {"type": "string"}, "os_version": {"type": "string"}}}, "events": {"type": "array", "items": {"$ref": "#/components/schemas/realtimekit_PeerReportEvent"}}, "ip_information": {"type": "object", "properties": {"asn": {"type": "object", "properties": {"asn": {"type": "string"}, "domain": {"type": "string"}, "name": {"type": "string"}, "route": {"type": "string"}, "type": {"type": "string"}}}, "city": {"type": "string"}, "country": {"type": "string"}, "ipv4": {"type": "string"}, "org": {"type": "string"}, "region": {"type": "string"}, "timezone": {"type": "string"}}}, "native_metadata": {"type": "object", "properties": {"audio_encoder": {"type": "string"}, "video_encoder": {"type": "string"}}}, "pc_metadata": {"type": "array", "items": {"properties": {"effective_network_type": {"type": "string"}, "reflexive_connectivity": {"type": "boolean"}, "relay_connectivity": {"type": "boolean"}, "sdp": {"type": "array", "items": {"type": "string"}}, "timestamp": {"type": "string"}, "turn_connectivity": {"type": "boolean"}}, "type": "object"}}, "room_view_type": {"type": "string"}, "sdk_name": {"type": "string"}, "sdk_type": {"type": "string"}, "sdk_version": {"type": "string"}, "selected_device_updates": {"type": "array", "items": {"properties": {"device": {"$ref": "#/components/schemas/realtimekit_PeerReportDevice"}, "timestamp": {"type": "string"}}, "type": "object"}}, "speaker_devices_updates": {"type": "array", "items": {"$ref": "#/components/schemas/realtimekit_PeerReportDeviceUpdate"}}, "video_devices_updates": {"type": "array", "items": {"$ref": "#/components/schemas/realtimekit_PeerReportDeviceUpdate"}}}}
```
