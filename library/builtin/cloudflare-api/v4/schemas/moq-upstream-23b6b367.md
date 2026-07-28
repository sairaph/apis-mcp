---
title: moq_upstream
page_id: schema-moq-upstream-23b6b367
path: schemas
description: A single upstream MOQT server publisher.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# moq_upstream

A single upstream MOQT server publisher.

```yaml
{"description": "A single upstream MOQT server publisher.", "type": "object", "properties": {"url": {"description": "Upstream MOQT server publisher URL. Must be an absolute URL with a\nhost and a scheme crique can dial: moqt:// (raw QUIC) or https://\n(WebTransport). Validated on update (PUT); rejected with 21013.\n", "type": "string", "format": "uri"}}, "required": ["url"]}
```
