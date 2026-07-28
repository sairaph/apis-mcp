---
title: tls-certificates-and-hostnames_status-11
page_id: schema-tls-certificates-and-hostnames-status-11-80f50e86
path: schemas
description: Client Certificates may be active or revoked, and the pending_reactivation or pending_revocation represent in-progress asynchronous transitions.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# tls-certificates-and-hostnames_status-11

Client Certificates may be active or revoked, and the pending_reactivation or pending_revocation represent in-progress asynchronous transitions.

```yaml
{"description": "Client Certificates may be active or revoked, and the pending_reactivation or pending_revocation represent in-progress asynchronous transitions.", "type": "string", "example": "active", "enum": ["active", "pending_reactivation", "pending_revocation", "revoked"], "x-auditable": true}
```
