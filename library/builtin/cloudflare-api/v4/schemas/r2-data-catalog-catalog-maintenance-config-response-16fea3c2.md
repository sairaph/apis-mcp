---
title: r2-data-catalog_catalog-maintenance-config-response
page_id: schema-r2-data-catalog-catalog-maintenance-config-response-16fea3c2
path: schemas
description: Contains maintenance configuration and credential status.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# r2-data-catalog_catalog-maintenance-config-response

Contains maintenance configuration and credential status.

```yaml
{"description": "Contains maintenance configuration and credential status.", "type": "object", "properties": {"credential_status": {"$ref": "#/components/schemas/r2-data-catalog_credential-status"}, "maintenance_config": {"$ref": "#/components/schemas/r2-data-catalog_catalog-maintenance-config"}}, "required": ["maintenance_config", "credential_status"]}
```
