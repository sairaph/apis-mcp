---
title: aaa_filter_options
page_id: schema-aaa-filter-options-f12795e9
path: schemas
description: 'Format of additional configuration options (filters) for the alert type. Data type of filters during policy creation: Array of strings.'
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# aaa_filter_options

Format of additional configuration options (filters) for the alert type. Data type of filters during policy creation: Array of strings.

```yaml
{"description": "Format of additional configuration options (filters) for the alert type. Data type of filters during policy creation: Array of strings.", "type": "array", "items": {}, "example": [{"AvailableValues": null, "ComparisonOperator": "==", "Key": "zones", "Range": "1-n"}, {"AvailableValues": [{"Description": "Service-Level Objective of 99.7", "ID": "99.7"}, {"Description": "Service-Level Objective of 99.8", "ID": "99.8"}], "ComparisonOperator": ">=", "Key": "slo", "Range": "0-1"}]}
```
