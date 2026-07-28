---
title: magic_psk_generation_response
page_id: schema-magic-psk-generation-response-b283f539
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# magic_psk_generation_response

```yaml
{"allOf": [{"$ref": "#/components/schemas/magic_api-response-single"}, {"properties": {"result": {"properties": {"ipsec_tunnel_id": {"$ref": "#/components/schemas/magic_identifier"}, "psk": {"$ref": "#/components/schemas/magic_psk"}, "psk_metadata": {"$ref": "#/components/schemas/magic_psk_metadata"}}}}}]}
```
