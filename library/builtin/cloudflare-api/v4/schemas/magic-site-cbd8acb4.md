---
title: magic_site
page_id: schema-magic-site-cbd8acb4
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# magic_site

```yaml
{"type": "object", "properties": {"connector_id": {"$ref": "#/components/schemas/magic_connector-id"}, "description": {"type": "string"}, "ha_mode": {"description": "Site high availability mode. If set to true, the site can have two connectors and runs in high availability mode.", "type": "boolean", "example": true}, "id": {"$ref": "#/components/schemas/magic_identifier"}, "location": {"$ref": "#/components/schemas/magic_site-location"}, "name": {"$ref": "#/components/schemas/magic_site-name"}, "secondary_connector_id": {"$ref": "#/components/schemas/magic_secondary-connector-id"}}}
```
