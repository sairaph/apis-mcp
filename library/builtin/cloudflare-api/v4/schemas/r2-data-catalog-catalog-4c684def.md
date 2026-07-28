---
title: r2-data-catalog_catalog
page_id: schema-r2-data-catalog-catalog-4c684def
path: schemas
description: Contains R2 Data Catalog information.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# r2-data-catalog_catalog

Contains R2 Data Catalog information.

```yaml
{"description": "Contains R2 Data Catalog information.", "type": "object", "properties": {"bucket": {"description": "Specifies the associated R2 bucket name.", "type": "string", "example": "my-data-bucket"}, "credential_status": {"description": "Shows the credential configuration status.", "type": "string", "allOf": [{"$ref": "#/components/schemas/r2-data-catalog_credential-status"}], "nullable": true}, "id": {"description": "Use this to uniquely identify the catalog.", "type": "string", "format": "uuid", "example": "550e8400-e29b-41d4-a716-446655440000"}, "maintenance_config": {"description": "Configures maintenance for the catalog.", "type": "object", "allOf": [{"$ref": "#/components/schemas/r2-data-catalog_catalog-maintenance-config"}], "nullable": true}, "name": {"description": "Specifies the catalog name (generated from account and bucket name).", "type": "string", "example": "account123_my-bucket"}, "status": {"$ref": "#/components/schemas/r2-data-catalog_catalog-status"}}, "required": ["id", "name", "bucket", "status"]}
```
