---
title: magic_health_check_base
page_id: schema-magic-health-check-base-ce157592
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# magic_health_check_base

```yaml
{"type": "object", "properties": {"enabled": {"description": "Determines whether to run healthchecks for a tunnel.", "type": "boolean", "example": true, "default": true, "x-auditable": true}, "rate": {"description": "How frequent the health check is run. The default value is `mid`.", "type": "string", "example": "low", "default": "mid", "enum": ["low", "mid", "high"], "x-auditable": true}, "target": {"description": "The destination address in a request type health check. After the healthcheck is decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded to this address. This field defaults to `customer_gre_endpoint address`. This field is ignored for bidirectional healthchecks as the interface_address (not assigned to the Cloudflare side of the tunnel) is used as the target. Must be in object form if the x-magic-new-hc-target header is set to true and string form if x-magic-new-hc-target is absent or set to false.", "oneOf": [{"$ref": "#/components/schemas/magic_health_check_target"}, {"type": "string"}]}, "type": {"description": "The type of healthcheck to run, reply or request. The default value is `reply`.", "type": "string", "example": "request", "default": "reply", "enum": ["reply", "request"], "x-auditable": true}}}
```
