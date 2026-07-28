---
title: workers_export_storage
page_id: schema-workers-export-storage-edfe0128
path: schemas
description: |-
    Durable Object storage backend. `sqlite` is the recommended (and
    only) backend for new namespaces. `legacy-kv` is accepted only for
    a class whose namespace already exists as KV-backed; the `exports`
    flow never provisions a new `legacy-kv` namespace.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_export_storage

Durable Object storage backend. `sqlite` is the recommended (and
only) backend for new namespaces. `legacy-kv` is accepted only for
a class whose namespace already exists as KV-backed; the `exports`
flow never provisions a new `legacy-kv` namespace.

```yaml
{"description": "Durable Object storage backend. `sqlite` is the recommended (and\nonly) backend for new namespaces. `legacy-kv` is accepted only for\na class whose namespace already exists as KV-backed; the `exports`\nflow never provisions a new `legacy-kv` namespace.\n", "type": "string", "enum": ["sqlite", "legacy-kv"]}
```
