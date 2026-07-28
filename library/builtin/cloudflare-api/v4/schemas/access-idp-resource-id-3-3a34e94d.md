---
title: access_idp_resource_id-3
page_id: schema-access-idp-resource-id-3-3a34e94d
path: schemas
description: |-
    The IdP-generated Id of the SCIM resource. Pass once for a single
    lookup (`?idp_resource_id=A`) or repeat the parameter
    (`?idp_resource_id=A&idp_resource_id=B`) to filter by multiple
    resources in one request.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_idp_resource_id-3

The IdP-generated Id of the SCIM resource. Pass once for a single
lookup (`?idp_resource_id=A`) or repeat the parameter
(`?idp_resource_id=A&idp_resource_id=B`) to filter by multiple
resources in one request.

```yaml
{"description": "The IdP-generated Id of the SCIM resource. Pass once for a single\nlookup (`?idp_resource_id=A`) or repeat the parameter\n(`?idp_resource_id=A&idp_resource_id=B`) to filter by multiple\nresources in one request.", "type": "array", "items": {"type": "string"}, "example": ["all_employees"]}
```
