---
title: resource-tagging_delete_tags_request_zone_level_access_application_policy
page_id: schema-resource-tagging-delete-tags-request-zone-level-access-application-polic-93217792
path: schemas
description: Request body schema for deleting tags from access_application_policy resources.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# resource-tagging_delete_tags_request_zone_level_access_application_policy

Request body schema for deleting tags from access_application_policy resources.

```yaml
{"description": "Request body schema for deleting tags from access_application_policy resources.", "allOf": [{"$ref": "#/components/schemas/resource-tagging_delete_tags_request_zone_level_base"}, {"properties": {"access_application_id": {"$ref": "#/components/schemas/resource-tagging_access_application_id"}, "resource_type": {"$ref": "#/components/schemas/resource-tagging_zone_resource_type_access_application_policy_enum"}}, "required": ["access_application_id"]}]}
```
