---
title: digital-experience-monitoring_device_isp
page_id: schema-digital-experience-monitoring-device-isp-91763b5b
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# digital-experience-monitoring_device_isp

```yaml
{"type": "object", "properties": {"ip": {"$ref": "#/components/schemas/digital-experience-monitoring_ip_info"}, "test_id": {"description": "The test that generated this result.", "type": "string", "example": "f174e90a-fafe-4643-bbbc-4a0ed4fc8415"}, "test_result_id": {"description": "The specific test result.", "type": "string", "example": "a1b2c3d4-e5f6-7890-abcd-ef1234567890"}, "time_start": {"description": "Timestamp of when the ISP was observed.", "type": "string", "format": "date-time", "example": "2024-06-01T12:00:00Z"}}, "required": ["time_start", "test_id", "test_result_id"]}
```
