---
title: access_resource_group_name
page_id: schema-access-resource-group-name-9b88c83f
path: schemas
description: |-
    The display name of the SCIM Group resource. Pass once for a single
    lookup (`?resource_group_name=A`) or repeat the parameter
    (`?resource_group_name=A&resource_group_name=B`) to filter by multiple
    group names in one request.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_resource_group_name

The display name of the SCIM Group resource. Pass once for a single
lookup (`?resource_group_name=A`) or repeat the parameter
(`?resource_group_name=A&resource_group_name=B`) to filter by multiple
group names in one request.

```yaml
{"description": "The display name of the SCIM Group resource. Pass once for a single\nlookup (`?resource_group_name=A`) or repeat the parameter\n(`?resource_group_name=A&resource_group_name=B`) to filter by multiple\ngroup names in one request.", "type": "array", "items": {"type": "string"}, "example": ["ALL_EMPLOYEES"]}
```
