---
title: realtimekit_CandidatePairStats
page_id: schema-realtimekit-candidatepairstats-0066ed86
path: schemas
description: ICE candidate pair statistics for a transport.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# realtimekit_CandidatePairStats

ICE candidate pair statistics for a transport.

```yaml
{"description": "ICE candidate pair statistics for a transport.", "type": "object", "properties": {"available_incoming_bitrate": {"type": "number"}, "available_outgoing_bitrate": {"type": "number"}, "bytes_discarded_on_send": {"type": "number"}, "bytes_received": {"type": "number"}, "bytes_sent": {"type": "number"}, "current_round_trip_time": {"type": "number"}, "last_packet_received_timestamp": {"description": "Epoch milliseconds when the last packet was received.", "type": "number"}, "last_packet_sent_timestamp": {"description": "Epoch milliseconds when the last packet was sent.", "type": "number"}, "local_candidate_address": {"type": "string"}, "local_candidate_id": {"type": "string"}, "local_candidate_network_type": {"type": "string"}, "local_candidate_port": {"type": "number"}, "local_candidate_protocol": {"type": "string"}, "local_candidate_related_address": {"type": "string"}, "local_candidate_related_port": {"type": "number"}, "local_candidate_type": {"type": "string"}, "local_candidate_url": {"type": "string"}, "nominated": {"type": "boolean"}, "packets_discarded_on_send": {"type": "number"}, "packets_received": {"type": "number"}, "packets_sent": {"type": "number"}, "remote_candidate_address": {"type": "string"}, "remote_candidate_id": {"type": "string"}, "remote_candidate_port": {"type": "number"}, "remote_candidate_protocol": {"type": "string"}, "remote_candidate_type": {"type": "string"}, "remote_candidate_url": {"type": "string"}, "total_round_trip_time": {"type": "number"}}}
```
