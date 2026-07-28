---
title: access_cf_resource_id-3
page_id: schema-access-cf-resource-id-3-3470837d
path: schemas
description: |-
    The unique Cloudflare-generated Id of the SCIM resource. Pass once for
    a single lookup (`?cf_resource_id=A`) or repeat the parameter
    (`?cf_resource_id=A&cf_resource_id=B`) to filter by multiple resources
    in one request.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_cf_resource_id-3

The unique Cloudflare-generated Id of the SCIM resource. Pass once for
a single lookup (`?cf_resource_id=A`) or repeat the parameter
(`?cf_resource_id=A&cf_resource_id=B`) to filter by multiple resources
in one request.

```yaml
{"description": "The unique Cloudflare-generated Id of the SCIM resource. Pass once for\na single lookup (`?cf_resource_id=A`) or repeat the parameter\n(`?cf_resource_id=A&cf_resource_id=B`) to filter by multiple resources\nin one request.", "type": "array", "items": {"type": "string"}, "example": ["bd97ef8d-7986-43e3-9ee0-c25dda33e4b0"]}
```
