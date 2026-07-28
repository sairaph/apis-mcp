---
title: hyperdrive_hyperdrive-config
page_id: schema-hyperdrive-hyperdrive-config-ecb7cfd7
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# hyperdrive_hyperdrive-config

```yaml
{"type": "object", "properties": {"caching": {"$ref": "#/components/schemas/hyperdrive_hyperdrive-caching"}, "created_on": {"description": "Defines the creation time of the Hyperdrive configuration.", "type": "string", "format": "date-time", "example": "2017-01-01T00:00:00Z", "readOnly": true, "x-auditable": true}, "id": {"$ref": "#/components/schemas/hyperdrive_identifier"}, "modified_on": {"description": "Defines the last modified time of the Hyperdrive configuration.", "type": "string", "format": "date-time", "example": "2017-01-01T00:00:00Z", "readOnly": true, "x-auditable": true}, "mtls": {"$ref": "#/components/schemas/hyperdrive_hyperdrive-mtls"}, "name": {"$ref": "#/components/schemas/hyperdrive_hyperdrive-name"}, "origin": {"type": "object", "oneOf": [{"allOf": [{"$ref": "#/components/schemas/hyperdrive_hyperdrive-database-full"}, {"$ref": "#/components/schemas/hyperdrive_internet-origin"}], "title": "Public Database"}, {"allOf": [{"$ref": "#/components/schemas/hyperdrive_hyperdrive-database-full"}, {"$ref": "#/components/schemas/hyperdrive_over-access-origin"}], "title": "Access-protected Database behind Cloudflare Tunnel"}, {"allOf": [{"$ref": "#/components/schemas/hyperdrive_hyperdrive-database-full"}, {"$ref": "#/components/schemas/hyperdrive_vpc-service-origin"}], "title": "Database reachable through a Workers VPC"}]}, "origin_connection_limit": {"$ref": "#/components/schemas/hyperdrive_hyperdrive-origin-connection-limit"}, "restarted_on": {"description": "Defines the last time the Hyperdrive connection pool was explicitly restarted via the restart endpoint. Omitted if the pool has never been explicitly restarted.", "type": "string", "format": "date-time", "example": "2017-01-01T00:00:00Z", "nullable": true, "readOnly": true, "x-auditable": true}}, "required": ["id", "name", "origin"]}
```
