---
title: access_certificates
page_id: schema-access-certificates-904ebc32
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_certificates

```yaml
{"type": "object", "properties": {"associated_hostnames": {"$ref": "#/components/schemas/access_associated_hostnames"}, "created_at": {"$ref": "#/components/schemas/access_created_at"}, "expires_on": {"$ref": "#/components/schemas/access_timestamp"}, "fingerprint": {"$ref": "#/components/schemas/access_fingerprint"}, "id": {"description": "The ID of the application that will use this certificate.", "type": "string", "x-auditable": true}, "name": {"$ref": "#/components/schemas/access_name-7"}, "updated_at": {"$ref": "#/components/schemas/access_updated_at"}}}
```
