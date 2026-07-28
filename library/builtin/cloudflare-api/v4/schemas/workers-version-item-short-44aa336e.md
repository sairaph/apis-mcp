---
title: workers_version-item-short
page_id: schema-workers-version-item-short-44aa336e
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers_version-item-short

```yaml
{"type": "object", "properties": {"id": {"description": "Unique identifier for the version.", "type": "string", "example": "18f97339-c287-4872-9bdd-e2135c07ec12", "readOnly": true, "x-auditable": true}, "metadata": {"type": "object", "example": {"author_email": "user@example.com", "author_id": "408cbcdfd4dda4617efef40b04d168a1", "created_on": "2022-11-08T17:19:29.176266Z", "modified_on": "2022-11-08T17:19:29.176266Z", "source": "api"}, "properties": {"author_email": {"description": "Email of the user who created the version.", "type": "string", "example": "user@example.com", "readOnly": true, "x-auditable": true}, "author_id": {"description": "Identifier of the user who created the version.", "type": "string", "example": "408cbcdfd4dda4617efef40b04d168a1", "readOnly": true, "x-auditable": true}, "created_on": {"description": "When the version was created.", "type": "string", "example": "2022-11-08T17:19:29.176266Z", "readOnly": true, "x-auditable": true}, "hasPreview": {"description": "Whether the version can be previewed.", "type": "boolean", "readOnly": true}, "modified_on": {"description": "When the version was last modified.", "type": "string", "example": "2022-11-08T17:19:29.176266Z", "readOnly": true, "x-auditable": true}, "source": {"description": "The source of the version upload.", "type": "string", "example": "api", "enum": ["unknown", "api", "wrangler", "terraform", "dash", "cf_cli", "dash_template", "integration", "quick_editor", "playground", "workersci"], "readOnly": true, "x-auditable": true}}, "readOnly": true}, "number": {"description": "Sequential version number.", "type": "number", "example": 1, "readOnly": true, "x-auditable": true}}}
```
