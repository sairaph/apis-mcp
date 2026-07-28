---
title: hyperdrive_hyperdrive-config-patch
page_id: schema-hyperdrive-hyperdrive-config-patch-93a73062
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# hyperdrive_hyperdrive-config-patch

```yaml
{"type": "object", "properties": {"caching": {"$ref": "#/components/schemas/hyperdrive_hyperdrive-caching"}, "mtls": {"$ref": "#/components/schemas/hyperdrive_hyperdrive-mtls"}, "name": {"$ref": "#/components/schemas/hyperdrive_hyperdrive-name"}, "origin": {"type": "object", "anyOf": [{"$ref": "#/components/schemas/hyperdrive_hyperdrive-database"}, {"oneOf": [{"$ref": "#/components/schemas/hyperdrive_internet-origin"}, {"$ref": "#/components/schemas/hyperdrive_over-access-origin"}, {"$ref": "#/components/schemas/hyperdrive_vpc-service-origin"}], "title": "Origin Database"}]}, "origin_connection_limit": {"$ref": "#/components/schemas/hyperdrive_hyperdrive-origin-connection-limit"}}}
```
