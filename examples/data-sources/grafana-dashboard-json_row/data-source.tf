data "grafana-dashboard-json_row" "example" {
  title = "Example row"
  position = {
    top  = 0
    left = 0
  }
}

data "grafana-dashboard-json_row" "example2" {
  title    = "Example second row"
  position = data.grafana-dashboard-json_row.example.next_position.next_row
}

data "grafana-dashboard-json_dashboard" "example" {
  // other attributes

  panels = [
    data.grafana-dashboard-json_row.example.rendered_json,
    data.grafana-dashboard-json_row.example2.rendered_json,
  ]
}
