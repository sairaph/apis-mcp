---
title: access_idp_resource_id-2
page_id: schema-access-idp-resource-id-2-d371b074
path: schemas
description: |-
    The IdP-generated Id of the SCIM User resource; also known as the "external Id".
    Pass once for a single lookup (`?idp_resource_id=A`) or repeat the parameter
    (`?idp_resource_id=A&idp_resource_id=B`) to look up multiple users in one request,
    up to 50 values. Mutually exclusive with `cf_resource_id`, `username`, `email`,
    `name`, `search_contains`, and `search_starts_with`.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_idp_resource_id-2

The IdP-generated Id of the SCIM User resource; also known as the "external Id".
Pass once for a single lookup (`?idp_resource_id=A`) or repeat the parameter
(`?idp_resource_id=A&idp_resource_id=B`) to look up multiple users in one request,
up to 50 values. Mutually exclusive with `cf_resource_id`, `username`, `email`,
`name`, `search_contains`, and `search_starts_with`.

```yaml
{"description": "The IdP-generated Id of the SCIM User resource; also known as the \"external Id\".\nPass once for a single lookup (`?idp_resource_id=A`) or repeat the parameter\n(`?idp_resource_id=A&idp_resource_id=B`) to look up multiple users in one request,\nup to 50 values. Mutually exclusive with `cf_resource_id`, `username`, `email`,\n`name`, `search_contains`, and `search_starts_with`.", "type": "array", "items": {"type": "string"}, "example": ["john_smith_01"], "maxItems": 50}
```
