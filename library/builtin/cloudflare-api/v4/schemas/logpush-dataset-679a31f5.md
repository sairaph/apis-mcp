---
title: logpush_dataset
page_id: schema-logpush-dataset-679a31f5
path: schemas
description: Name of the dataset. A list of supported datasets can be found on the [Developer Docs](https://developers.cloudflare.com/logs/reference/log-fields/).
source: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
source_type: openapi
imported_from: https://raw.githubusercontent.com/cloudflare/api-schemas/c92b9b0fde23ae00fece2025662f96dc8e2d6283/openapi.json
---

# logpush_dataset

Name of the dataset. A list of supported datasets can be found on the [Developer Docs](https://developers.cloudflare.com/logs/reference/log-fields/).

```yaml
{"description": "Name of the dataset. A list of supported datasets can be found on the [Developer Docs](https://developers.cloudflare.com/logs/reference/log-fields/).", "type": "string", "example": "http_requests", "default": "http_requests", "enum": ["access_requests", "audit_logs", "audit_logs_v2", "biso_user_actions", "casb_findings", "device_posture_results", "dex_application_tests", "dex_device_state_events", "dlp_forensic_copies", "dns_firewall_logs", "dns_logs", "email_security_alerts", "email_security_post_delivery_events", "firewall_events", "gateway_dns", "gateway_http", "gateway_network", "http_requests", "ipsec_logs", "magic_ids_detections", "mcp_portal_logs", "mnm_flow_logs", "nel_reports", "network_analytics_logs", "page_shield_events", "sinkhole_http_logs", "spectrum_events", "ssh_logs", "turnstile_events", "warp_config_changes", "warp_toggle_changes", "websocket_analytics", "workers_trace_events", "zaraz_events", "zero_trust_network_sessions"], "nullable": true, "x-auditable": true}
```
