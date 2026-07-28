---
title: dns-records_priority
page_id: schema-dns-records-priority-04cc1942
path: schemas
description: Required for MX and URI records; ignored for other record types (but may still be returned by the API). Records with lower priorities are preferred. This field is to be deprecated in favor of the priority field within the data map.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# dns-records_priority

Required for MX and URI records; ignored for other record types (but may still be returned by the API). Records with lower priorities are preferred. This field is to be deprecated in favor of the priority field within the data map.

```yaml
{"description": "Required for MX and URI records; ignored for other record types (but may still be returned by the API). Records with lower priorities are preferred. This field is to be deprecated in favor of the priority field within the data map.", "type": "number", "example": 10, "maximum": 65535, "minimum": 0, "x-auditable": true}
```
