<?php

/* dump.twig */
class __TwigTemplate_8c3e87ca49c1090ea8074cd7bca0403981b63e4affa18ce0621fcc64b8ea1929 extends Twig_Template
{
    private $source;

    public function __construct(Twig_Environment $env)
    {
        parent::__construct($env);

        $this->source = $this->getSourceContext();

        // line 1
        $this->parent = $this->loadTemplate("default.twig", "dump.twig", 1);
        $this->blocks = array(
            'title' => array($this, 'block_title'),
            'content' => array($this, 'block_content'),
        );
    }

    protected function doGetParent(array $context)
    {
        return "default.twig";
    }

    protected function doDisplay(array $context, array $blocks = array())
    {
        $this->parent->display($context, array_merge($this->blocks, $blocks));
    }

    // line 3
    public function block_title($context, array $blocks = array())
    {
        // line 4
        echo "data dump!
";
    }

    // line 7
    public function block_content($context, array $blocks = array())
    {
        // line 8
        echo "<div class=\"container\">
  ";
        // line 9
        echo twig_escape_filter($this->env, twig_var_dump($this->env, $context), "html", null, true);
        echo "
</div>
";
    }

    public function getTemplateName()
    {
        return "dump.twig";
    }

    public function isTraitable()
    {
        return false;
    }

    public function getDebugInfo()
    {
        return array (  47 => 9,  44 => 8,  41 => 7,  36 => 4,  33 => 3,  15 => 1,);
    }

    public function getSourceContext()
    {
        return new Twig_Source("{% extends \"default.twig\" %}

{% block title %}
data dump!
{% endblock %}

{% block content %}
<div class=\"container\">
  {{ dump() }}
</div>
{% endblock %}", "dump.twig", "/home/ubuntu/test/shimizu/KB_1802/webApp/templates/dump.twig");
    }
}
