---
title: resource-tagging_tagged_resource_object_access_application_policy_base
page_id: schema-resource-tagging-tagged-resource-object-access-application-policy-base-9338c3ee
path: schemas
description: Base schema for access_application_policy resources (without type discriminator)
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# resource-tagging_tagged_resource_object_access_application_policy_base

Base schema for access_application_policy resources (without type discriminator)

```yaml
{"description": "Base schema for access_application_policy resources (without type discriminator)", "type": "object", "properties": {"access_application_id": {"$ref": "#/components/schemas/resource-tagging_access_application_id"}, "etag": {"$ref": "#/components/schemas/resource-tagging_etag"}, "id": {"$ref": "#/components/schemas/resource-tagging_resource_id"}, "name": {"$ref": "#/components/schemas/resource-tagging_resource_name"}, "tags": {"$ref": "#/components/schemas/resource-tagging_tags"}, "zone_id": {"$ref": "#/components/schemas/resource-tagging_zone_id"}}, "required": ["id", "tags", "name", "etag", "zone_id", "access_application_id"]}
```
