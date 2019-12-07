<?php

    if ($_SERVER["REQUEST_METHOD"] === "GET") {
        $num = filter_input(INPUT_GET, "n", FILTER_VALIDATE_INT);
        if (!is_int($num)) {
            header($_SERVER["SERVER_PROTOCOL"] . " 400 Bad Request.", true, 400);
            exit();
        }

        $file_name = realpath(sprintf("%s/../json/%d.json", __DIR__, $num));
        if (file_exists($file_name) && mime_content_type($file_name) === "application/json") {
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