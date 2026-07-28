---
title: access_authenticator_device_aaguid
page_id: schema-access-authenticator-device-aaguid-a7022944
path: schemas
description: A FIDO2 authenticator device AAGUID entry
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# access_authenticator_device_aaguid

A FIDO2 authenticator device AAGUID entry

```yaml
{"description": "A FIDO2 authenticator device AAGUID entry", "type": "object", "properties": {"aaguid": {"$ref": "#/components/schemas/access_aaguid"}, "name": {"$ref": "#/components/schemas/access_name-14"}}, "required": ["aaguid", "name"]}
```
