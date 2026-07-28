---
title: intel-sinkholes_sinkhole_item
page_id: schema-intel-sinkholes-sinkhole-item-04d80f00
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# intel-sinkholes_sinkhole_item

```yaml
{"type": "object", "properties": {"account_tag": {"description": "The account tag that owns this sinkhole.", "type": "string"}, "created_on": {"description": "The date and time when the sinkhole was created.", "type": "string", "format": "date-time"}, "id": {"description": "The unique identifier for the sinkhole.", "type": "string"}, "modified_on": {"description": "The date and time when the sinkhole was last modified.", "type": "string", "format": "date-time"}, "name": {"description": "The name of the sinkhole.", "type": "string"}, "r2_bucket": {"description": "The name of the R2 bucket to store results.", "type": "string"}, "r2_id": {"description": "The id of the R2 instance.", "type": "string"}}, "example": {"account_tag": "233f45e61fd1f7e21e1e154ede4q2859", "created_on": "2023-05-12T12:21:56.777653Z", "id": "93defa6e909e464e8c89a85859f36d3c", "modified_on": "2023-06-18T03:13:34.123321Z", "name": "my_sinkhole", "r2_bucket": "my_bucket", "r2_id": "example_r2_id"}}
```
