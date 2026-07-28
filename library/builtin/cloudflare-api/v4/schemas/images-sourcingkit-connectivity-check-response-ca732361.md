---
title: images_sourcingkit_connectivity_check_response
page_id: schema-images-sourcingkit-connectivity-check-response-ca732361
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# images_sourcingkit_connectivity_check_response

```yaml
{"allOf": [{"$ref": "#/components/schemas/images_api-response-single"}, {"properties": {"result": {"type": "object", "properties": {"code": {"description": "Machine-readable error code if connectivity failed.", "type": "string", "nullable": true}, "connectivityStatus": {"description": "Whether the connectivity check succeeded.", "type": "string", "enum": ["ok", "error"]}, "reason": {"description": "Human-readable error description if connectivity failed.", "type": "string", "nullable": true}}}}}]}
```
