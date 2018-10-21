<?php

/* default.twig */
class __TwigTemplate_377f3ea56d0ae7e085e62f423a2aeba6f29c9c5b4f12b7b511bf6070050fdde5 extends Twig_Template
{
    private $source;

    public function __construct(Twig_Environment $env)
    {
        parent::__construct($env);

        $this->source = $this->getSourceContext();

        $this->parent = false;

        $this->blocks = array(
            'title' => array($this, 'block_title'),
            'content' => array($this, 'block_content'),
        );
    }

    protected function doDisplay(array $context, array $blocks = array())
    {
        // line 1
        echo "<!DOCTYPE html>
<html lang=\"ja\">
<head>
  <meta charset=\"UTF-8\">
  <meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">
  <meta http-equiv=\"X-UA-Compatible\" content=\"ie=edge\">
  <title>";
        // line 7
        $this->displayBlock('title', $context, $blocks);
        echo "</title>
  <link rel=\"stylesheet\" href=\"";
        // line 8
        echo twig_escape_filter($this->env, $this->extensions['Slim\Views\TwigExtension']->baseUrl(), "html", null, true);
        echo "/assets/css/style.css\">
  <script src=\"https://cdnjs.cloudflare.com/ajax/libs/Chart.js/2.7.2/Chart.bundle.min.js\"></script>
</head>
<body>
  ";
        // line 12
        $this->loadTemplate("includes/header.twig", "default.twig", 12)->display($context);
        // line 13
        echo "  <div class=\"content\">
  ";
        // line 14
        $this->displayBlock('content', $context, $blocks);
        // line 15
        echo "  </div>
  ";
        // line 16
        $this->loadTemplate("includes/footer.twig", "default.twig", 16)->display($context);
        // line 17
        echo "</body>
</html>";
    }

    // line 7
    public function block_title($context, array $blocks = array())
    {
    }

    // line 14
    public function block_content($context, array $blocks = array())
    {
    }

    public function getTemplateName()
    {
        return "default.twig";
    }

    public function isTraitable()
    {
        return false;
    }

    public function getDebugInfo()
    {
        return array (  66 => 14,  61 => 7,  56 => 17,  54 => 16,  51 => 15,  49 => 14,  46 => 13,  44 => 12,  37 => 8,  33 => 7,  25 => 1,);
    }

    public function getSourceContext()
    {
        return new Twig_Source("<!DOCTYPE html>
<html lang=\"ja\">
<head>
  <meta charset=\"UTF-8\">
  <meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">
  <meta http-equiv=\"X-UA-Compatible\" content=\"ie=edge\">
  <title>{% block title %}{% endblock %}</title>
  <link rel=\"stylesheet\" href=\"{{ base_url() }}/assets/css/style.css\">
  <script src=\"https://cdnjs.cloudflare.com/ajax/libs/Chart.js/2.7.2/Chart.bundle.min.js\"></script>
</head>
<body>
  {% include 'includes/header.twig' %}
  <div class=\"content\">
  {% block content %}{% endblock %}
  </div>
  {% include 'includes/footer.twig' %}
</body>
</html>", "default.twig", "/home/ubuntu/test/shimizu/KB_1802/webApp/templates/default.twig");
    }
}
