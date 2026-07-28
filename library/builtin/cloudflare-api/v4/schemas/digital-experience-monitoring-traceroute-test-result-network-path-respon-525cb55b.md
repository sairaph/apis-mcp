---
title: digital-experience-monitoring_traceroute_test_result_network_path_response
page_id: schema-digital-experience-monitoring-traceroute-test-result-network-path-respon-525cb55b
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# digital-experience-monitoring_traceroute_test_result_network_path_response

```yaml
{"type": "object", "properties": {"colo": {"$ref": "#/components/schemas/digital-experience-monitoring_colo"}, "deviceName": {"description": "Name of the device associated with this network path response.", "type": "string"}, "execution_context": {"description": "Whether the test was run inside or outside of the WARP tunnel.", "allOf": [{"$ref": "#/components/schemas/digital-experience-monitoring_execution_context"}]}, "hops": {"description": "An array of the hops taken by the device to reach the end destination.", "type": "array", "items": {"properties": {"asn": {"type": "integer", "nullable": true}, "aso": {"type": "string", "nullable": true}, "ipAddress": {"type": "string", "nullable": true}, "location": {"type": "object", "nullable": true, "properties": {"city": {"type": "string", "nullable": true}, "state": {"type": "string", "nullable": true}, "zip": {"type": "string", "nullable": true}}}, "mile": {"type": "string", "enum": ["client-to-app", "client-to-cf-egress", "client-to-cf-ingress", "client-to-isp"], "nullable": true}, "name": {"type": "string", "nullable": true}, "packetLossPct": {"type": "number", "format": "float", "nullable": true}, "rttMs": {"type": "integer", "nullable": true}, "ttl": {"type": "integer"}}, "required": ["ttl"], "type": "object"}}, "resultId": {"$ref": "#/components/schemas/digital-experience-monitoring_uuid"}, "testId": {"$ref": "#/components/schemas/digital-experience-monitoring_uuid"}, "testName": {"description": "Name of the traceroute test.", "type": "string"}, "time_start": {"description": "Timestamp indicating when the traceroute test execution began.", "type": "string", "example": "2023-07-16 15:00:00+00"}, "tunnel_type": {"type": "string", "nullable": true}}, "required": ["resultId", "hops"]}
```
