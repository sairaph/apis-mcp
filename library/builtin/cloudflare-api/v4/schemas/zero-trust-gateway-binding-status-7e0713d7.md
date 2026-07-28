---
title: zero-trust-gateway_binding_status
page_id: schema-zero-trust-gateway-binding-status-7e0713d7
path: schemas
description: Indicate the read-only deployment status of the certificate on Cloudflare's edge. Gateway TLS interception can use certificates in the 'available' (previously called 'active') state.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zero-trust-gateway_binding_status

Indicate the read-only deployment status of the certificate on Cloudflare's edge. Gateway TLS interception can use certificates in the 'available' (previously called 'active') state.

```yaml
{"description": "Indicate the read-only deployment status of the certificate on Cloudflare's edge. Gateway TLS interception can use certificates in the 'available' (previously called 'active') state.", "type": "string", "enum": ["pending_deployment", "available", "pending_deletion", "inactive"], "readOnly": true, "x-auditable": true}
```
