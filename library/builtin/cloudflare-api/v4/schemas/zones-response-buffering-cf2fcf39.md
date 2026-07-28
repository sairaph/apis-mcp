---
title: zones_response_buffering
page_id: schema-zones-response-buffering-cf2fcf39
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zones_response_buffering

```yaml
{"type": "object", "properties": {"id": {"description": "Turn on or off whether Cloudflare should wait for an entire file\nfrom the origin server before forwarding it to the site visitor. By\ndefault, Cloudflare sends packets to the client as they arrive from\nthe origin server.\n", "type": "string", "enum": ["response_buffering"], "x-auditable": true}, "value": {"description": "The status of Response Buffering\n", "type": "string", "example": "on", "enum": ["on", "off"], "x-auditable": true}}, "deprecated": true, "title": "Response Buffering", "x-stainless-deprecation-message": "This page rule is deprecated. This functionality is no longer supported.", "x-stainless-skip": ["terraform"]}
```
