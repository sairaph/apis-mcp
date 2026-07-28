---
title: resource-tagging_set_tags_request_account_level_worker_version
page_id: schema-resource-tagging-set-tags-request-account-level-worker-version-929d7b5f
path: schemas
description: Request body schema for setting tags on worker_version resources.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# resource-tagging_set_tags_request_account_level_worker_version

Request body schema for setting tags on worker_version resources.

```yaml
{"description": "Request body schema for setting tags on worker_version resources.", "allOf": [{"$ref": "#/components/schemas/resource-tagging_delete_tags_request_account_level_worker_version"}, {"properties": {"tags": {"$ref": "#/components/schemas/resource-tagging_tags"}}}]}
```
