<?php

use Slim\Http\Request;
use Slim\Http\Response;

// Routes

$app->get('/',function (Request $request, Response $response, array $args)
{
    // For Demo
    $rand = rand(0,1);
    // Render index view
    $view = $this->get('view');
    $view->render($response, 'home.twig', array(
        "raw_camera_data" => "http://ec2-52-197-145-249.ap-northeast-1.compute.amazonaws.com/resultData/result.jpg",
        "evaluation" => "Good！",
        "clean_seq" => "80",
    ));
    return $response;
});

$app->get('/dump',function (Request $request, Response $response, array $args)
{
    // Sample log message

    // Render index view
    $view = $this->get('view');
    $view->render($response, 'dump.twig', []);
    return $response;
});