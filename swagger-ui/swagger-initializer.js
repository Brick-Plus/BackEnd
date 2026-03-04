window.onload = function() {
  window.ui = SwaggerUIBundle({
    url: "/config.openapi.yaml",
    dom_id: '#swagger-ui',
  });
};
