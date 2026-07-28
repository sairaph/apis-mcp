---
title: one_UseCaseDetail
page_id: schema-one-usecasedetail-70a2e775
path: schemas
description: Full use case with scopes and features for detail endpoint.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# one_UseCaseDetail

Full use case with scopes and features for detail endpoint.

```yaml
{"description": "Full use case with scopes and features for detail endpoint.", "type": "object", "properties": {"base_scopes": {"description": "Scopes always required for this use case.", "type": "array", "items": {"$ref": "#/components/schemas/one_Permission"}}, "description": {"description": "Use case description.", "type": "string"}, "display_name": {"description": "Human-readable use case name.", "type": "string"}, "features": {"description": "Optional features with extra scopes.", "type": "array", "items": {"$ref": "#/components/schemas/one_FeatureScope"}}, "id": {"description": "Use case identifier.", "type": "string"}}, "required": ["base_scopes", "description", "display_name", "features", "id"]}
```
