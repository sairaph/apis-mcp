---
title: magic-visibility-pcaps_pcaps_request_simple
page_id: schema-magic-visibility-pcaps-pcaps-request-simple-12f32e79
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# magic-visibility-pcaps_pcaps_request_simple

```yaml
{"type": "object", "properties": {"filter_v1": {"$ref": "#/components/schemas/magic-visibility-pcaps_pcaps_filter_v1"}, "offset_time": {"$ref": "#/components/schemas/magic-visibility-pcaps_pcaps_offset_time"}, "packet_limit": {"$ref": "#/components/schemas/magic-visibility-pcaps_pcaps_packet_limit"}, "system": {"$ref": "#/components/schemas/magic-visibility-pcaps_pcaps_system"}, "time_limit": {"$ref": "#/components/schemas/magic-visibility-pcaps_pcaps_time_limit_sampled"}, "type": {"$ref": "#/components/schemas/magic-visibility-pcaps_pcaps_type"}}, "required": ["time_limit", "packet_limit", "type", "system"]}
```
