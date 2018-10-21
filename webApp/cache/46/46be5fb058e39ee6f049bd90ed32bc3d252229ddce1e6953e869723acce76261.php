<?php

/* includes/header.twig */
class __TwigTemplate_7e73e17842f30ae7092f41d06dc724ab8112aadc3fc8b842f750820c23f6f172 extends Twig_Template
{
    private $source;

    public function __construct(Twig_Environment $env)
    {
        parent::__construct($env);

        $this->source = $this->getSourceContext();

        $this->parent = false;

        $this->blocks = array(
        );
    }

    protected function doDisplay(array $context, array $blocks = array())
    {
        // line 1
        echo "<section class=\"hero is-primary is-medium\">
  <!-- Hero content: will be in the middle -->
  <div class=\"hero-body\">
    <div class=\"container\">
    </div>
  </div>

  <!-- Hero footer: will stick at the bottom -->
  <div class=\"hero-foot\">
    <nav class=\"tabs\">
      <div class=\"container\">
        <ul>
          <li><a>今週のベストショット</a></li>
          <li><a>友人の反応</a></li>
          <li><a>通知設定</a></li>
          <li><a>カメラ設定</a></li>
        </ul>
      </div>
    </nav>
  </div>
</section>
";
    }

    public function getTemplateName()
    {
        return "includes/header.twig";
    }

    public function getDebugInfo()
    {
        return array (  23 => 1,);
    }

    public function getSourceContext()
    {
        return new Twig_Source("<section class=\"hero is-primary is-medium\">
  <!-- Hero content: will be in the middle -->
  <div class=\"hero-body\">
    <div class=\"container\">
    </div>
  </div>

  <!-- Hero footer: will stick at the bottom -->
  <div class=\"hero-foot\">
    <nav class=\"tabs\">
      <div class=\"container\">
        <ul>
          <li><a>今週のベストショット</a></li>
          <li><a>友人の反応</a></li>
          <li><a>通知設定</a></li>
          <li><a>カメラ設定</a></li>
        </ul>
      </div>
    </nav>
  </div>
</section>
", "includes/header.twig", "/home/ubuntu/test/shimizu/KB_1802/webApp/templates/includes/header.twig");
    }
}
