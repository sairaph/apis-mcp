---
title: spectrum-config_app_config
page_id: schema-spectrum-config-app-config-a2e42f38
path: schemas
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# spectrum-config_app_config

```yaml
{"allOf": [{"$ref": "#/components/schemas/spectrum-config_base_app_config"}, {"properties": {"argo_smart_routing": {"$ref": "#/components/schemas/spectrum-config_argo_smart_routing"}, "dns": {"$ref": "#/components/schemas/spectrum-config_dns"}, "edge_ips": {"$ref": "#/components/schemas/spectrum-config_edge_ips"}, "ip_firewall": {"$ref": "#/components/schemas/spectrum-config_ip_firewall"}, "origin_direct": {"$ref": "#/components/schemas/spectrum-config_origin_direct"}, "origin_dns": {"$ref": "#/components/schemas/spectrum-config_origin_dns"}, "origin_port": {"$ref": "#/components/schemas/spectrum-config_origin_port"}, "protocol": {"$ref": "#/components/schemas/spectrum-config_protocol"}, "proxy_protocol": {"$ref": "#/components/schemas/spectrum-config_proxy_protocol"}, "tls": {"$ref": "#/components/schemas/spectrum-config_tls"}, "traffic_type": {"$ref": "#/components/schemas/spectrum-config_traffic_type"}, "virtual_network_id": {"$ref": "#/components/schemas/spectrum-config_virtual_network_id"}}, "required": ["protocol", "dns", "traffic_type"], "type": "object"}]}
```
