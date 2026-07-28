---
title: magic-transit_endpoint_health_check
page_id: schema-magic-transit-endpoint-health-check-a65311b8
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# magic-transit_endpoint_health_check

```yaml
{"type": "object", "properties": {"check_type": {"$ref": "#/components/schemas/magic-transit_check_type"}, "endpoint": {"description": "the IP address of the host to perform checks against", "type": "string", "example": "203.0.113.1"}, "name": {"description": "Optional name associated with this check", "type": "string", "example": "My Endpoint"}}, "required": ["check_type", "endpoint"]}
```
