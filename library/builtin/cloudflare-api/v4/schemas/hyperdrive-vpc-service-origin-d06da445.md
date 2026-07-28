---
title: hyperdrive_vpc-service-origin
page_id: schema-hyperdrive-vpc-service-origin-d06da445
path: schemas
description: Connect to a database through a Workers VPC Service. TLS settings (mTLS, sslmode) cannot be configured on the Hyperdrive when using a VPC Service origin; TLS must be managed on the VPC Service itself.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# hyperdrive_vpc-service-origin

Connect to a database through a Workers VPC Service. TLS settings (mTLS, sslmode) cannot be configured on the Hyperdrive when using a VPC Service origin; TLS must be managed on the VPC Service itself.

```yaml
{"description": "Connect to a database through a Workers VPC Service. TLS settings (mTLS, sslmode) cannot be configured on the Hyperdrive when using a VPC Service origin; TLS must be managed on the VPC Service itself.\n", "type": "object", "properties": {"service_id": {"description": "The identifier of the Workers VPC Service to connect through. Hyperdrive will egress through the specified VPC Service to reach the origin database.", "type": "string", "example": "0123456789abcdef0123456789abcdef", "x-auditable": true}}, "required": ["service_id"], "title": "Database reachable through a Workers VPC"}
```
