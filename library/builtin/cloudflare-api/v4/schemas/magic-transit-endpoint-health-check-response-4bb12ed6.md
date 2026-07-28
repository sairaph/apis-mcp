---
title: magic-transit_endpoint_health_check_response
page_id: schema-magic-transit-endpoint-health-check-response-4bb12ed6
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# magic-transit_endpoint_health_check_response

```yaml
{"type": "object", "allOf": [{"$ref": "#/components/schemas/magic-transit_endpoint_health_check"}, {"properties": {"id": {"$ref": "#/components/schemas/magic-transit_uuid"}}}], "required": ["id", "check_type", "endpoint"]}
```
