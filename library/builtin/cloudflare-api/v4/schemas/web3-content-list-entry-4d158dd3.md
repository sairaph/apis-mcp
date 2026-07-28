---
title: web3_content_list_entry
page_id: schema-web3-content-list-entry-4d158dd3
path: schemas
description: Specify a content list entry to block.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# web3_content_list_entry

Specify a content list entry to block.

```yaml
{"description": "Specify a content list entry to block.", "type": "object", "properties": {"content": {"$ref": "#/components/schemas/web3_content_list_entry_content"}, "created_on": {"$ref": "#/components/schemas/web3_timestamp"}, "description": {"$ref": "#/components/schemas/web3_content_list_entry_description"}, "id": {"$ref": "#/components/schemas/web3_identifier"}, "modified_on": {"$ref": "#/components/schemas/web3_timestamp"}, "type": {"$ref": "#/components/schemas/web3_content_list_entry_type"}}}
```
