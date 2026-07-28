---
title: resource-tagging_set_tags_request_zone_level_access_application_policy
page_id: schema-resource-tagging-set-tags-request-zone-level-access-application-policy-fe94038b
path: schemas
description: Request body schema for setting tags on access_application_policy resources.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# resource-tagging_set_tags_request_zone_level_access_application_policy

Request body schema for setting tags on access_application_policy resources.

```yaml
{"description": "Request body schema for setting tags on access_application_policy resources.", "allOf": [{"$ref": "#/components/schemas/resource-tagging_delete_tags_request_zone_level_access_application_policy"}, {"properties": {"tags": {"$ref": "#/components/schemas/resource-tagging_tags"}}}]}
```
