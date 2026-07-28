---
title: hyperdrive_hyperdrive-mtls
page_id: schema-hyperdrive-hyperdrive-mtls-801c4e4a
path: schemas
description: mTLS configuration for the origin connection. Cannot be used with VPC Service origins; TLS must be managed on the VPC Service.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# hyperdrive_hyperdrive-mtls

mTLS configuration for the origin connection. Cannot be used with VPC Service origins; TLS must be managed on the VPC Service.

```yaml
{"description": "mTLS configuration for the origin connection. Cannot be used with VPC Service origins; TLS must be managed on the VPC Service.", "type": "object", "properties": {"ca_certificate_id": {"description": "Define CA certificate ID obtained after uploading CA cert.", "type": "string", "example": "00000000-0000-0000-0000-0000000000"}, "mtls_certificate_id": {"description": "Define mTLS certificate ID obtained after uploading client cert.", "type": "string", "example": "00000000-0000-0000-0000-0000000000"}, "sslmode": {"description": "Set SSL mode to 'require', 'verify-ca', or 'verify-full' to verify the CA.", "type": "string", "example": "verify-full"}}, "title": "mTLS"}
```
