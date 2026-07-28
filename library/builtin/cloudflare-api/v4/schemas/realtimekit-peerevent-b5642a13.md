---
title: realtimekit_PeerEvent
page_id: schema-realtimekit-peerevent-b5642a13
path: schemas
description: A connection lifecycle event recorded for a participant's peer.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# realtimekit_PeerEvent

A connection lifecycle event recorded for a participant's peer.

```yaml
{"description": "A connection lifecycle event recorded for a participant's peer.", "type": "object", "properties": {"created_at": {"description": "Timestamp when this peer event was created.", "type": "string"}, "event_name": {"description": "Name of the peer event.", "type": "string", "enum": ["PEER_CREATED", "PEER_JOINING", "PEER_LEAVING"]}, "id": {"description": "ID of the peer event.", "type": "string"}, "minutes_consumed": {"description": "Minutes consumed attributed to this event.", "type": "number"}, "participant_id": {"description": "ID of the participant this event belongs to.", "type": "string", "nullable": true}, "peer_id": {"description": "Peer ID this event belongs to.", "type": "string"}, "preset_view_type": {"description": "View type of the preset associated with the peer.", "type": "string", "enum": ["GROUP_CALL", "WEBINAR", "AUDIO_ROOM", "LIVESTREAM", "CHAT"], "nullable": true}, "session_id": {"description": "ID of the session this event belongs to.", "type": "string", "nullable": true}, "socket_session_id": {"description": "ID of the socket session associated with this event.", "type": "string", "nullable": true}, "updated_at": {"description": "Timestamp when this peer event was last updated.", "type": "string"}}}
```
