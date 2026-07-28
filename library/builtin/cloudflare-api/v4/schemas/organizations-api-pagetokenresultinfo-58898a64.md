---
title: organizations-api_PageTokenResultInfo
page_id: schema-organizations-api-pagetokenresultinfo-58898a64
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# organizations-api_PageTokenResultInfo

```yaml
{"type": "object", "properties": {"next_page_token": {"description": "Use this opaque token in the next request to retrieve the\nnext page.\n\nParameters used to filter the retrieved list must remain in subsequent\nrequests with a page token.", "type": "string"}, "total_size": {"description": "Counts the total amount of items in a list with the applied filters. The API omits next_page_token to indicate no more items in a particular list.", "type": "integer"}}}
```
