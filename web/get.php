<?php

    if ($_SERVER["REQUEST_METHOD"] !== "GET") {
        header($_SERVER["SERVER_PROTOCOL"] . " 405 Method Not Allowed.", true, 405);
        header("Allow: GET");
        exit();
    }

    $num = filter_input(INPUT_GET, "n", FILTER_VALIDATE_INT);
    if (!is_int($num) || $num < 0 || $num > PHP_INT_MAX) {
        header($_SERVER["SERVER_PROTOCOL"] . " 400 Bad Request.", true, 400);
        exit();
    }
    $type = filter_input(INPUT_GET, "t");
    if ($type === "j") {
        $file_name = realpath(sprintf("%s/../json/%d.json", __DIR__, $num));
        if (file_exists($file_name) && pathinfo($file_name, PATHINFO_EXTENSION) === "json") {
            header("Content-Type: application/json");
            header("Content-Length: " . filesize($file_name));
            readfile($file_name);
            exit();
        } else {
            header($_SERVER["SERVER_PROTOCOL"] . " 404 Not Found.", true, 404);
            exit();
        }
    } elseif ($type === "s") {
        $file_name = realpath(sprintf("%s/../svgs/%d.svg", __DIR__, $num));
        if (file_exists($file_name) && pathinfo($file_name, PATHINFO_EXTENSION) === "svg") {
            header("Content-Type: image/svg+xml");
            header("Content-Length: " . filesize($file_name));
            readfile($file_name);
            exit();
        } else {
            header($_SERVER["SERVER_PROTOCOL"] . " 404 Not Found.", true, 404);
            exit();
        }
    } else {
        header($_SERVER["SERVER_PROTOCOL"] . " 400 Bad Request.", true, 400);
        exit();
    }