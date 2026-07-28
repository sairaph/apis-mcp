---
title: access_cf_resource_id
page_id: schema-access-cf-resource-id-2ab1338e
path: schemas
description: |-
    The unique Cloudflare-generated Id of the SCIM Group resource; also known as the "Id".
    Pass once for a single lookup (`?cf_resource_id=A`) or repeat the parameter
    (`?cf_resource_id=A&cf_resource_id=B`) to look up multiple groups in one request,
    up to 50 values. Mutually exclusive with `idp_resource_id`, `name`,
    `search_contains`, and `search_starts_with`.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_cf_resource_id

The unique Cloudflare-generated Id of the SCIM Group resource; also known as the "Id".
Pass once for a single lookup (`?cf_resource_id=A`) or repeat the parameter
(`?cf_resource_id=A&cf_resource_id=B`) to look up multiple groups in one request,
up to 50 values. Mutually exclusive with `idp_resource_id`, `name`,
`search_contains`, and `search_starts_with`.

```yaml
{"description": "The unique Cloudflare-generated Id of the SCIM Group resource; also known as the \"Id\".\nPass once for a single lookup (`?cf_resource_id=A`) or repeat the parameter\n(`?cf_resource_id=A&cf_resource_id=B`) to look up multiple groups in one request,\nup to 50 values. Mutually exclusive with `idp_resource_id`, `name`,\n`search_contains`, and `search_starts_with`.", "type": "array", "items": {"type": "string"}, "example": ["a2abeb50-59c9-4c01-8c5c-963d3bf5700f"], "maxItems": 50}
```
