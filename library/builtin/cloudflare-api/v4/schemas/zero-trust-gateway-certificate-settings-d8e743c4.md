---
title: zero-trust-gateway_certificate-settings
page_id: schema-zero-trust-gateway-certificate-settings-d8e743c4
path: schemas
description: Specify certificate settings for Gateway TLS interception. If unset, the Cloudflare Root CA handles interception.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# zero-trust-gateway_certificate-settings

Specify certificate settings for Gateway TLS interception. If unset, the Cloudflare Root CA handles interception.

```yaml
{"description": "Specify certificate settings for Gateway TLS interception. If unset, the Cloudflare Root CA handles interception.", "type": "object", "properties": {"id": {"description": "Specify the UUID of the certificate used for interception. Ensure the certificate is available at the edge(previously called 'active'). A nil UUID directs Cloudflare to use the Root CA.", "type": "string", "example": "d1b364c5-1311-466e-a194-f0e943e0799f", "x-auditable": true}}, "nullable": true, "required": ["id"], "x-stainless-terraform-configurability": "optional"}
```
