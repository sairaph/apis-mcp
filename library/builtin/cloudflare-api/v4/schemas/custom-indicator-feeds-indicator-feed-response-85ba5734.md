---
title: custom-indicator-feeds_indicator_feed_response
page_id: schema-custom-indicator-feeds-indicator-feed-response-85ba5734
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# custom-indicator-feeds_indicator_feed_response

```yaml
{"allOf": [{"$ref": "#/components/schemas/custom-indicator-feeds_api-response-common"}, {"properties": {"result": {"type": "array", "items": {"$ref": "#/components/schemas/custom-indicator-feeds_indicator_feed_item"}, "example": [{"created_on": "2023-05-12T12:21:56.777653Z", "description": "user specified description 1", "id": 1, "modified_on": "2023-06-18T03:13:34.123321Z", "name": "user_specified_name_1"}, {"created_on": "2023-05-21T21:43:52.867525Z", "description": "User specified description 2", "id": 2, "modified_on": "2023-06-28T18:46:18.764425Z", "name": "user_specified_name_2"}]}}}]}
```
