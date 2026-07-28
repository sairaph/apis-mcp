---
title: usage-analytics_billing_usage_data_point
page_id: schema-usage-analytics-billing-usage-data-point-263084f2
path: schemas
description: A single billing usage data point for a time period, containing metrics for all billable products.
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# usage-analytics_billing_usage_data_point

A single billing usage data point for a time period, containing metrics for all billable products.

```yaml
{"description": "A single billing usage data point for a time period, containing metrics for all billable products.\n", "type": "object", "properties": {"argoAcceleratedBytes": {"description": "Number of Argo accelerated bytes in this time period.", "type": "integer", "format": "int64", "example": 5000000}, "imageResizingRequests": {"description": "Number of Image Resizing requests in this time period.", "type": "integer", "format": "int64", "example": 15000}, "loadBalancingQueries": {"description": "Number of Load Balancing DNS queries in this time period.", "type": "integer", "format": "int64", "example": 10000}, "mediaUniqueTransformations": {"description": "Number of Media unique image transformations in this time period.", "type": "integer", "format": "int64", "example": 45000}, "rateLimitingRequestsAllowed": {"description": "Number of Rate Limiting requests allowed in this time period.", "type": "integer", "format": "int64", "example": 50000}, "spectrumBytesTransferred": {"description": "Number of Spectrum bytes transferred in this time period.", "type": "integer", "format": "int64", "example": 8000000}, "streamMinutesViewed": {"description": "Number of Stream billable minutes viewed in this time period.", "type": "integer", "format": "int64", "example": 125000}, "ts": {"description": "Unix timestamp (epoch seconds) for the start of this time period.", "type": "integer", "format": "int64", "example": 1693526400}, "workersKVReads": {"description": "Number of Workers KV reads in this time period.", "type": "integer", "format": "int64", "example": 30000}, "workersRequests": {"description": "Number of Workers requests in this time period.", "type": "integer", "format": "int64", "example": 200000}}}
```
