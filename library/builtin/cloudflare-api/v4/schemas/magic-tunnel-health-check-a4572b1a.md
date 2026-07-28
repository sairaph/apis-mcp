---
title: magic_tunnel_health_check
page_id: schema-magic-tunnel-health-check-a4572b1a
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# magic_tunnel_health_check

```yaml
{"allOf": [{"$ref": "#/components/schemas/magic_health_check_base"}, {"properties": {"direction": {"description": "The direction of the flow of the healthcheck. Either unidirectional, where the probe comes to you via the tunnel and the result comes back to Cloudflare via the open Internet, or bidirectional where both the probe and result come and go via the tunnel.", "type": "string", "example": "bidirectional", "default": "unidirectional", "enum": ["unidirectional", "bidirectional"]}}}]}
```
