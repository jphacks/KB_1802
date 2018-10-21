<?php

/* home.twig */
class __TwigTemplate_226704c3f68adf0faddde78e9bdc472883b14fa3621ea3600fef9646db755516 extends Twig_Template
{
    private $source;

    public function __construct(Twig_Environment $env)
    {
        parent::__construct($env);

        $this->source = $this->getSourceContext();

        // line 1
        $this->parent = $this->loadTemplate("default.twig", "home.twig", 1);
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
        echo "Soisy
";
    }

    // line 7
    public function block_content($context, array $blocks = array())
    {
        // line 8
        echo "<div class=\"container\">
    <h1>現在の汚部屋</h1>
    <div class=\"columns\">
        <div class=\"column is-three-fifths\">
            <img src=\"";
        // line 12
        echo twig_escape_filter($this->env, ($context["raw_camera_data"] ?? null), "html", null, true);
        echo "\">
        </div>
        <div class=\"column\">
            <h2>
                評価
            </h2>
            <p class=\"title is-size-1 has-text-info\">";
        // line 18
        echo twig_escape_filter($this->env, ($context["evaluation"] ?? null), "html", null, true);
        echo "</p>
            <h2>
                スコア
            </h2>
            <p class=\"title\"><span class=\"clean_seq\">";
        // line 22
        echo twig_escape_filter($this->env, ($context["clean_seq"] ?? null), "html", null, true);
        echo "</span>/100</p>
        </div>
    </div>
    <h1>部屋ログ</h1>
    <div class=\"columns\">
        <div class=\"column is-three-fifths\">
            <canvas id=\"myChart\" style=\"width:100%;height:300px\"></canvas>
        </div>
        <div class=\"column\">
            <article class=\"message is-primary\">
            <div class=\"message-header\">
                <p>2018年10月20日</p>
            </div>
            <div class=\"message-body\">
                <img src=\"http://ec2-52-197-145-249.ap-northeast-1.compute.amazonaws.com/camData/header20181020-192741.jpg\" class=\"past-camera\">
            </div>
            </article>
        </div>
    </div>
</div>
";
        // line 42
        $this->loadTemplate("includes/scripts.twig", "home.twig", 42)->display($context);
    }

    public function getTemplateName()
    {
        return "home.twig";
    }

    public function isTraitable()
    {
        return false;
    }

    public function getDebugInfo()
    {
        return array (  89 => 42,  66 => 22,  59 => 18,  50 => 12,  44 => 8,  41 => 7,  36 => 4,  33 => 3,  15 => 1,);
    }

    public function getSourceContext()
    {
        return new Twig_Source("{% extends \"default.twig\" %}

{% block title %}
Soisy
{% endblock %}

{% block content %}
<div class=\"container\">
    <h1>現在の汚部屋</h1>
    <div class=\"columns\">
        <div class=\"column is-three-fifths\">
            <img src=\"{{ raw_camera_data }}\">
        </div>
        <div class=\"column\">
            <h2>
                評価
            </h2>
            <p class=\"title is-size-1 has-text-info\">{{ evaluation }}</p>
            <h2>
                スコア
            </h2>
            <p class=\"title\"><span class=\"clean_seq\">{{ clean_seq }}</span>/100</p>
        </div>
    </div>
    <h1>部屋ログ</h1>
    <div class=\"columns\">
        <div class=\"column is-three-fifths\">
            <canvas id=\"myChart\" style=\"width:100%;height:300px\"></canvas>
        </div>
        <div class=\"column\">
            <article class=\"message is-primary\">
            <div class=\"message-header\">
                <p>2018年10月20日</p>
            </div>
            <div class=\"message-body\">
                <img src=\"http://ec2-52-197-145-249.ap-northeast-1.compute.amazonaws.com/camData/header20181020-192741.jpg\" class=\"past-camera\">
            </div>
            </article>
        </div>
    </div>
</div>
{% include \"includes/scripts.twig\" %}
{% endblock %}", "home.twig", "/home/ubuntu/test/shimizu/KB_1802/webApp/templates/home.twig");
    }
}
