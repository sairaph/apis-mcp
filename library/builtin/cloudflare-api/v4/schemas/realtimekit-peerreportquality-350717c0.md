---
title: realtimekit_PeerReportQuality
page_id: schema-realtimekit-peerreportquality-350717c0
path: schemas
description: Media quality statistics for the participant.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# realtimekit_PeerReportQuality

Media quality statistics for the participant.

```yaml
{"description": "Media quality statistics for the participant.", "type": "object", "properties": {"audio_consumer": {"type": "array", "items": {"$ref": "#/components/schemas/realtimekit_AudioConsumerStats"}}, "audio_consumer_cumulative": {"$ref": "#/components/schemas/realtimekit_AudioConsumerStatsCumulative"}, "audio_producer": {"type": "array", "items": {"$ref": "#/components/schemas/realtimekit_AudioProducerStats"}}, "audio_producer_cumulative": {"$ref": "#/components/schemas/realtimekit_AudioProducerStatsCumulative"}, "screenshare_audio_consumer": {"type": "array", "items": {"$ref": "#/components/schemas/realtimekit_AudioConsumerStats"}}, "screenshare_audio_consumer_cumulative": {"$ref": "#/components/schemas/realtimekit_AudioConsumerStatsCumulative"}, "screenshare_audio_producer": {"type": "array", "items": {"$ref": "#/components/schemas/realtimekit_AudioProducerStats"}}, "screenshare_audio_producer_cumulative": {"$ref": "#/components/schemas/realtimekit_AudioProducerStatsCumulative"}, "screenshare_video_consumer": {"type": "array", "items": {"$ref": "#/components/schemas/realtimekit_VideoConsumerStats"}}, "screenshare_video_consumer_cumulative": {"$ref": "#/components/schemas/realtimekit_VideoConsumerStatsCumulative"}, "screenshare_video_producer": {"type": "array", "items": {"$ref": "#/components/schemas/realtimekit_VideoProducerStats"}}, "screenshare_video_producer_cumulative": {"$ref": "#/components/schemas/realtimekit_VideoProducerStatsCumulative"}, "video_consumer": {"type": "array", "items": {"$ref": "#/components/schemas/realtimekit_VideoConsumerStats"}}, "video_consumer_cumulative": {"$ref": "#/components/schemas/realtimekit_VideoConsumerStatsCumulative"}, "video_producer": {"type": "array", "items": {"$ref": "#/components/schemas/realtimekit_VideoProducerStats"}}, "video_producer_cumulative": {"$ref": "#/components/schemas/realtimekit_VideoProducerStatsCumulative"}}}
```
