---
title: custom-indicator-feeds_indicator_feed_item
page_id: schema-custom-indicator-feeds-indicator-feed-item-8ea01646
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# custom-indicator-feeds_indicator_feed_item

```yaml
{"properties": {"created_on": {"description": "The date and time when the data entry was created", "type": "string", "format": "date-time", "x-auditable": true}, "description": {"$ref": "#/components/schemas/custom-indicator-feeds_description"}, "id": {"$ref": "#/components/schemas/custom-indicator-feeds_id"}, "is_attributable": {"$ref": "#/components/schemas/custom-indicator-feeds_is_attributable"}, "is_downloadable": {"$ref": "#/components/schemas/custom-indicator-feeds_is_downloadable"}, "is_public": {"$ref": "#/components/schemas/custom-indicator-feeds_is_public"}, "modified_on": {"description": "The date and time when the data entry was last modified", "type": "string", "format": "date-time", "x-auditable": true}, "name": {"$ref": "#/components/schemas/custom-indicator-feeds_name"}}, "example": {"created_on": "2023-05-12T12:21:56.777653Z", "description": "example feed description", "id": 1, "is_attributable": false, "is_downloadable": false, "is_public": false, "modified_on": "2023-06-18T03:13:34.123321Z", "name": "example_feed_1"}}
```
