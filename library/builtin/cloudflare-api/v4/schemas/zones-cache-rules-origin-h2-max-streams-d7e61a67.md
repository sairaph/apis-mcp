---
title: zones_cache-rules_origin_h2_max_streams
page_id: schema-zones-cache-rules-origin-h2-max-streams-d7e61a67
path: schemas
description: Origin H2 Max Streams configures the max number of concurrent requests that Cloudflare will send within the same connection when communicating with the origin server, if the origin supports it. Note that if your origin does not support H2 multiplexing, 5xx errors may be observed, particularly 520s. Also note that the default value is `100` for all plan types except Enterprise where it is `1`. `1` means that H2 multiplexing is disabled.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zones_cache-rules_origin_h2_max_streams

Origin H2 Max Streams configures the max number of concurrent requests that Cloudflare will send within the same connection when communicating with the origin server, if the origin supports it. Note that if your origin does not support H2 multiplexing, 5xx errors may be observed, particularly 520s. Also note that the default value is `100` for all plan types except Enterprise where it is `1`. `1` means that H2 multiplexing is disabled.

```yaml
{"description": "Origin H2 Max Streams configures the max number of concurrent requests that Cloudflare will send within the same connection when communicating with the origin server, if the origin supports it. Note that if your origin does not support H2 multiplexing, 5xx errors may be observed, particularly 520s. Also note that the default value is `100` for all plan types except Enterprise where it is `1`. `1` means that H2 multiplexing is disabled.", "type": "object", "allOf": [{"$ref": "#/components/schemas/zones_cache-rules_base"}, {"properties": {"id": {"description": "Value of the zone setting.", "type": "string", "example": "origin_h2_max_streams", "enum": ["origin_h2_max_streams"], "x-auditable": true}, "value": {"$ref": "#/components/schemas/zones_cache-rules_origin_h2_max_streams_value"}}, "type": "object"}], "title": "Origin H2 Max Streams"}
```
