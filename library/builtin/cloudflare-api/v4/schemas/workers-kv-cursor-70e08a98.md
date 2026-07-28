---
title: workers-kv_cursor
page_id: schema-workers-kv-cursor-70e08a98
path: schemas
description: Opaque token indicating the position from which to continue when requesting the next set of records if the amount of list results was limited by the limit parameter. A valid value for the cursor can be obtained from the cursors object in the result_info structure.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# workers-kv_cursor

Opaque token indicating the position from which to continue when requesting the next set of records if the amount of list results was limited by the limit parameter. A valid value for the cursor can be obtained from the cursors object in the result_info structure.

```yaml
{"description": "Opaque token indicating the position from which to continue when requesting the next set of records if the amount of list results was limited by the limit parameter. A valid value for the cursor can be obtained from the cursors object in the result_info structure.", "type": "string", "example": "6Ck1la0VxJ0djhidm1MdX2FyDGxLKVeeHZZmORS_8XeSuhz9SjIJRaSa2lnsF01tQOHrfTGAP3R5X1Kv5iVUuMbNKhWNAXHOl6ePB0TUL8nw", "x-stainless-pagination-property": {"purpose": "next_cursor_field"}}
```
