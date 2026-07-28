---
title: hyperdrive_hyperdrive-caching
page_id: schema-hyperdrive-hyperdrive-caching-eb2cf901
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# hyperdrive_hyperdrive-caching

```yaml
{"type": "object", "discriminator": {"mapping": {"false": "#/components/schemas/hyperdrive_hyperdrive-caching-enabled", "true": "#/components/schemas/hyperdrive_hyperdrive-caching-disabled"}, "propertyName": "disabled"}, "oneOf": [{"$ref": "#/components/schemas/hyperdrive_hyperdrive-caching-disabled"}, {"$ref": "#/components/schemas/hyperdrive_hyperdrive-caching-enabled"}]}
```
