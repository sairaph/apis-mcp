---
title: zones_cache-rules_origin_max_http_version
page_id: schema-zones-cache-rules-origin-max-http-version-26f4648b
path: schemas
description: Origin Max HTTP Setting Version sets the highest HTTP version Cloudflare will attempt to use with your origin. This setting allows Cloudflare to make HTTP/2 requests to your origin. (Refer to [Enable HTTP/2 to Origin](https://developers.cloudflare.com/cache/how-to/enable-http2-to-origin/), for more information.). The default value is "2" for all plan types except Enterprise where it is "1".
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zones_cache-rules_origin_max_http_version

Origin Max HTTP Setting Version sets the highest HTTP version Cloudflare will attempt to use with your origin. This setting allows Cloudflare to make HTTP/2 requests to your origin. (Refer to [Enable HTTP/2 to Origin](https://developers.cloudflare.com/cache/how-to/enable-http2-to-origin/), for more information.). The default value is "2" for all plan types except Enterprise where it is "1".

```yaml
{"description": "Origin Max HTTP Setting Version sets the highest HTTP version Cloudflare will attempt to use with your origin. This setting allows Cloudflare to make HTTP/2 requests to your origin. (Refer to [Enable HTTP/2 to Origin](https://developers.cloudflare.com/cache/how-to/enable-http2-to-origin/), for more information.). The default value is \"2\" for all plan types except Enterprise where it is \"1\".", "type": "object", "allOf": [{"$ref": "#/components/schemas/zones_cache-rules_base"}, {"properties": {"id": {"description": "Value of the zone setting.", "type": "string", "example": "origin_max_http_version", "enum": ["origin_max_http_version"], "x-auditable": true}, "value": {"$ref": "#/components/schemas/zones_cache-rules_origin_max_http_version_value"}}, "type": "object"}], "title": "Origin Max HTTP Version"}
```
