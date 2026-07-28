---
title: digital-experience-monitoring_traceroute_test_network_path_response
page_id: schema-digital-experience-monitoring-traceroute-test-network-path-response-f4f359fb
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# digital-experience-monitoring_traceroute_test_network_path_response

```yaml
{"type": "object", "properties": {"deviceName": {"description": "Name of the device that ran the test.", "type": "string"}, "id": {"$ref": "#/components/schemas/digital-experience-monitoring_uuid"}, "interval": {"description": "The interval at which the Traceroute synthetic application test is set to run.", "type": "string", "example": "0h5m0s"}, "kind": {"enum": ["traceroute"]}, "name": {"type": "string"}, "networkPath": {"type": "object", "nullable": true, "properties": {"sampling": {"description": "Specifies the sampling applied, if any, to the slots response. When sampled, results shown represent the first test run to the start of each sampling interval.", "type": "object", "nullable": true, "properties": {"unit": {"enum": ["hours"]}, "value": {"type": "integer"}}, "required": ["value", "unit"]}, "slots": {"type": "array", "items": {"properties": {"clientToAppRttMs": {"description": "Round trip time in ms of the client to app mile", "type": "integer", "nullable": true}, "clientToCfEgressRttMs": {"description": "Round trip time in ms of the client to Cloudflare egress mile", "type": "integer", "nullable": true}, "clientToCfIngressRttMs": {"description": "Round trip time in ms of the client to Cloudflare ingress mile", "type": "integer", "nullable": true}, "clientToIspRttMs": {"description": "Round trip time in ms of the client to ISP mile", "type": "integer", "nullable": true}, "id": {"$ref": "#/components/schemas/digital-experience-monitoring_uuid"}, "timestamp": {"type": "string", "example": "2023-07-16 15:00:00+00"}}, "required": ["id", "timestamp", "clientToAppRttMs", "clientToCfIngressRttMs", "clientToCfEgressRttMs"], "type": "object"}}}, "required": ["slots"]}, "url": {"description": "The host of the Traceroute synthetic application test.", "type": "string", "example": "1.1.1.1"}}, "required": ["id"]}
```
