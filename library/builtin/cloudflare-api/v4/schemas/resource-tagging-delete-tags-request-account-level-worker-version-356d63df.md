---
title: resource-tagging_delete_tags_request_account_level_worker_version
page_id: schema-resource-tagging-delete-tags-request-account-level-worker-version-356d63df
path: schemas
description: Request body schema for deleting tags from worker_version resources.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# resource-tagging_delete_tags_request_account_level_worker_version

Request body schema for deleting tags from worker_version resources.

```yaml
{"description": "Request body schema for deleting tags from worker_version resources.", "allOf": [{"$ref": "#/components/schemas/resource-tagging_delete_tags_request_account_level_base"}, {"properties": {"resource_type": {"$ref": "#/components/schemas/resource-tagging_account_resource_type_worker_version_enum"}, "worker_id": {"$ref": "#/components/schemas/resource-tagging_worker_id"}}, "required": ["worker_id"]}]}
```
