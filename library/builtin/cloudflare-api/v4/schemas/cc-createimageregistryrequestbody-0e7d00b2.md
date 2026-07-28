---
title: cc_CreateImageRegistryRequestBody
page_id: schema-cc-createimageregistryrequestbody-0e7d00b2
path: schemas
description: Request body for creating a new image registry configuration
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# cc_CreateImageRegistryRequestBody

Request body for creating a new image registry configuration

```yaml
{"description": "Request body for creating a new image registry configuration", "type": "object", "properties": {"auth": {"$ref": "#/components/schemas/cc_ImageRegistryAuth"}, "domain": {"$ref": "#/components/schemas/cc_Domain"}, "is_public": {"description": "If you own the registry and is private, this should be false or not defined. If it's a public registry like docker.io, you should set this to true", "type": "boolean"}, "kind": {"$ref": "#/components/schemas/cc_ExternalRegistryKind"}}, "required": ["domain"]}
```
