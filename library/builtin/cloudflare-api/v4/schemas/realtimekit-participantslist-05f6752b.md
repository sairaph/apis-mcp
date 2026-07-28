---
title: realtimekit_ParticipantsList
page_id: schema-realtimekit-participantslist-05f6752b
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# realtimekit_ParticipantsList

```yaml
{"type": "object", "properties": {"created_at": {"description": "timestamp when this participant was created.", "type": "string"}, "custom_participant_id": {"description": "ID passed by client to create this participant.", "type": "string"}, "display_name": {"description": "Display name of participant when joining the session.", "type": "string"}, "duration": {"description": "number of minutes for which the participant was in the session.", "type": "number"}, "id": {"description": "Participant ID. This maps to the corresponding peerId.", "type": "string"}, "joined_at": {"description": "timestamp at which participant joined the session.", "type": "string"}, "left_at": {"description": "timestamp at which participant left the session.", "type": "string"}, "peer_events": {"description": "Connection lifecycle events for the participant's peer. Only included when `include_peer_events` is true.", "type": "array", "items": {"$ref": "#/components/schemas/realtimekit_PeerEvent"}}, "preset_name": {"description": "Name of the preset associated with the participant.", "type": "string"}, "updated_at": {"description": "timestamp when this participant's data was last updated.", "type": "string"}, "user_id": {"description": "User id for this participant.", "type": "string"}}}
```
