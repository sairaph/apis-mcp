---
title: r2-data-catalog_namespace-list-response
page_id: schema-r2-data-catalog-namespace-list-response-d2d85d9b
path: schemas
description: Contains the list of namespaces with optional pagination.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# r2-data-catalog_namespace-list-response

Contains the list of namespaces with optional pagination.

```yaml
{"description": "Contains the list of namespaces with optional pagination.", "type": "object", "properties": {"details": {"description": "Contains detailed metadata for each namespace when return_details is true.\nEach object includes the namespace, UUID, and timestamps.\n", "type": "array", "items": {"$ref": "#/components/schemas/r2-data-catalog_namespace-details"}, "nullable": true}, "namespace_uuids": {"description": "Contains UUIDs for each namespace when return_uuids is true.\nThe order corresponds to the namespaces array.\n", "type": "array", "items": {"format": "uuid", "type": "string"}, "example": ["0199b999-6869-7383-bb1f-d30e059d5326", "0199b99b-2c88-73b3-8dbb-421e0e8f2757"], "nullable": true}, "namespaces": {"description": "Lists namespaces in the catalog.", "type": "array", "items": {"$ref": "#/components/schemas/r2-data-catalog_namespace-identifier"}}, "next_page_token": {"description": "Use this opaque token to fetch the next page of results.\nA null or absent value indicates the last page.\n", "type": "string", "example": "MSYxNzU5NzU1NTc4NTA0MTk0JjAxOTliOTliLTJjODgtNzNiMy04ZGJiLTQyMWUwZThmMjc1Nw", "nullable": true}}, "required": ["namespaces"]}
```
