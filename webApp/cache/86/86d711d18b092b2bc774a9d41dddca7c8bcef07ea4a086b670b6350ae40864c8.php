<?php

/* includes/scripts.twig */
class __TwigTemplate_dfca605066f70ccf8e7c1f2e4a036ba8decede39e73748f9fbcdd9d060f426a6 extends Twig_Template
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
        echo "<script class=\"cssdeck\" src=\"//cdnjs.cloudflare.com/ajax/libs/jquery/1.8.0/jquery.min.js\"></script>
<script>
var ctx = document.getElementById('myChart').getContext('2d');
var myChart = new Chart(ctx, {
  type: 'line',
  data: {
    labels: ['Apr.', 'May', 'Jun.', 'Jul.', 'Aug', 'Sep.', 'Oct.'],
    datasets: [{
      label: 'score',
      data: [12, 19, 3, 17, 6, 3, 7],
      backgroundColor: \"rgba(153,255,51,0.4)\"
    }]
  }
});
</script>";
    }

    public function getTemplateName()
    {
        return "includes/scripts.twig";
    }

    public function getDebugInfo()
    {
        return array (  23 => 1,);
    }

    public function getSourceContext()
    {
        return new Twig_Source("<script class=\"cssdeck\" src=\"//cdnjs.cloudflare.com/ajax/libs/jquery/1.8.0/jquery.min.js\"></script>
<script>
var ctx = document.getElementById('myChart').getContext('2d');
var myChart = new Chart(ctx, {
  type: 'line',
  data: {
    labels: ['Apr.', 'May', 'Jun.', 'Jul.', 'Aug', 'Sep.', 'Oct.'],
    datasets: [{
      label: 'score',
      data: [12, 19, 3, 17, 6, 3, 7],
      backgroundColor: \"rgba(153,255,51,0.4)\"
    }]
  }
});
</script>", "includes/scripts.twig", "/home/ubuntu/test/shimizu/KB_1802/webApp/templates/includes/scripts.twig");
    }
}
