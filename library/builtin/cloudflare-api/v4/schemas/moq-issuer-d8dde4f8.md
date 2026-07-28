---
title: moq_issuer
page_id: schema-moq-issuer-d8dde4f8
path: schemas
description: One arm of the discriminated-union token collection.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# moq_issuer

One arm of the discriminated-union token collection.

```yaml
{"description": "One arm of the discriminated-union token collection.", "discriminator": {"mapping": {"cloudflare_jwt": "#/components/schemas/moq_cloudflare_jwt_issuer"}, "propertyName": "type"}, "oneOf": [{"$ref": "#/components/schemas/moq_cloudflare_jwt_issuer"}]}
```
