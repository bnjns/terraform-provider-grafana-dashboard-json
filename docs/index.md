page_title: "grafana-dashboard-json Provider"
description: |-
  This provider allows practitioners to programmatically generate the JSON for your Grafana dashboards, using native Terraform functionality.
---

# Grafana Dashboard JSON Provider

Managing dashboards with the [Grafana Terraform provider](https://registry.terraform.io/providers/grafana/grafana)
requires writing the JSON manually; this makes dashboard configuration verbose and difficult to maintain.

This provider allows you to programmatically generate your dashboard JSON using "native" Terraform functionality, such
as blocks and arguments. This provider does not actually create your dashboards, but is intended to be used in
conjunction with the [Grafana Terraform provider](https://registry.terraform.io/providers/grafana/grafana).

->This provider is designed to be a proof-of-concept; there is an [open issue](https://github.com/grafana/terraform-provider-grafana/issues/299)
which hopes that this functionality becomes available in the official `grafana/grafana` provider.
<br><br>While this provider should be safe for practitioners to use in their production systems, development may lag
behind the Grafana API and official Terraform provider.

## Example Usage

```terraform
provider "grafana-dashboard-json" {
  // No provider configuration is necessary
}
```
