---
title: magic-visibility-pcaps_pcaps_request_full
page_id: schema-magic-visibility-pcaps-pcaps-request-full-527c1c22
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# magic-visibility-pcaps_pcaps_request_full

```yaml
{"type": "object", "properties": {"byte_limit": {"$ref": "#/components/schemas/magic-visibility-pcaps_pcaps_byte_limit"}, "colo_name": {"$ref": "#/components/schemas/magic-visibility-pcaps_pcaps_colo_name"}, "destination_conf": {"$ref": "#/components/schemas/magic-visibility-pcaps_pcaps_destination_conf"}, "filter_v1": {"$ref": "#/components/schemas/magic-visibility-pcaps_pcaps_filter_v1"}, "packet_limit": {"$ref": "#/components/schemas/magic-visibility-pcaps_pcaps_packet_limit"}, "system": {"$ref": "#/components/schemas/magic-visibility-pcaps_pcaps_system"}, "time_limit": {"$ref": "#/components/schemas/magic-visibility-pcaps_pcaps_time_limit_full"}, "type": {"$ref": "#/components/schemas/magic-visibility-pcaps_pcaps_type"}}, "required": ["time_limit", "type", "system", "colo_name", "destination_conf"]}
```
