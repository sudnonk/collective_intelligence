<?php

    if ($_SERVER["REQUEST_METHOD"] === "GET") {
        $num = filter_input(INPUT_GET, "n", FILTER_VALIDATE_INT);
        if (!is_int($num) || $num < 0 || $num > PHP_INT_MAX) {
            header($_SERVER["SERVER_PROTOCOL"] . " 400 Bad Request.", true, 400);
            exit();
        }

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
    } else {
        header($_SERVER["SERVER_PROTOCOL"] . " 405 Method Not Allowed.", true, 405);
        header("Allow: GET");
        exit();
    }