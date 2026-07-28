---
title: registrar-api-sandbox_domain_name
page_id: schema-registrar-api-sandbox-domain-name-f9d3f737
path: schemas
description: |-
    Fully qualified domain name (FQDN) including the extension
    (e.g., `example.com`, `mybrand.app`). The domain name uniquely
    identifies a registration — the same domain cannot be registered
    twice, making it a natural idempotency key for registration requests.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# registrar-api-sandbox_domain_name

Fully qualified domain name (FQDN) including the extension
(e.g., `example.com`, `mybrand.app`). The domain name uniquely
identifies a registration — the same domain cannot be registered
twice, making it a natural idempotency key for registration requests.

```yaml
{"description": "Fully qualified domain name (FQDN) including the extension\n(e.g., `example.com`, `mybrand.app`). The domain name uniquely\nidentifies a registration — the same domain cannot be registered\ntwice, making it a natural idempotency key for registration requests.\n", "type": "string", "example": "example.com", "x-auditable": true}
```
