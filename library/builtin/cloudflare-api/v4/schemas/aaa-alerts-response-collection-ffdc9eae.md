---
title: aaa_alerts-response_collection
page_id: schema-aaa-alerts-response-collection-ffdc9eae
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# aaa_alerts-response_collection

```yaml
{"allOf": [{"$ref": "#/components/schemas/aaa_api-response-common-2"}, {"properties": {"result": {"type": "object", "example": {"Origin Monitoring": [{"description": "High levels of 5xx HTTP errors at your origin.", "display_name": "Origin Error Rate Alert", "filter_options": [{"AvailableValues": null, "ComparisonOperator": "==", "Key": "zones", "Range": "1-n"}, {"AvailableValues": [{"Description": "Service-Level Objective of 99.7", "ID": "99.7"}, {"Description": "Service-Level Objective of 99.8", "ID": "99.8"}], "ComparisonOperator": ">=", "Key": "slo", "Range": "0-1"}], "type": "http_alert_origin_error"}]}, "additionalProperties": {"items": {"$ref": "#/components/schemas/aaa_alert-types"}, "type": "array"}}}}]}
```
