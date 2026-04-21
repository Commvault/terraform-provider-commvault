provider "commvault" {
  web_service_url = var.web_service_url
  user_name       = var.user_name
  password        = var.password
  ignore_cert     = true   # Set to false if CommServe has a valid SSL certificate
}
