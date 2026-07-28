---
title: tls-certificates-and-hostnames_custom_cert_bundle
page_id: schema-tls-certificates-and-hostnames-custom-cert-bundle-89e56b23
path: schemas
description: Array of custom certificate and key pairs (1 or 2 pairs allowed).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# tls-certificates-and-hostnames_custom_cert_bundle

Array of custom certificate and key pairs (1 or 2 pairs allowed).

```yaml
{"description": "Array of custom certificate and key pairs (1 or 2 pairs allowed).", "type": "array", "items": {"$ref": "#/components/schemas/tls-certificates-and-hostnames_custom_cert_and_key"}, "maxItems": 2, "minItems": 1}
```
