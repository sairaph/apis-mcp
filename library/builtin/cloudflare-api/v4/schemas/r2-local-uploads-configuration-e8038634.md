---
title: r2_local_uploads_configuration
page_id: schema-r2-local-uploads-configuration-e8038634
path: schemas
description: Configuration for local uploads on a bucket.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# r2_local_uploads_configuration

Configuration for local uploads on a bucket.

```yaml
{"description": "Configuration for local uploads on a bucket.", "type": "object", "properties": {"enabled": {"description": "Whether local uploads is enabled for this bucket. When enabled, object's data is written to the nearest region first, then asynchronously replicated to the bucket's primary region.", "type": "boolean", "x-auditable": true}}}
```
