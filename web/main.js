const prev = $("#prev");
const start = $("#start");
const next = $("#next");
const top_button = $("#top");

const canvas = $("#canvas");

const cell_id = $("#cell_id");
const fearness = $("#fearness");
const kindness = $("#kindness");
const resource = $("#resource");
const width = $("#width");

let step = 0;
let fps = 5;
let selected_id;

$(prev).on("click", () => {
    animation_end();
    step--;
    if (step < 0) {
        step = 0;
    }
    show();
});

let is_on = false;

function toggle_button() {
    if (is_on) {
        $(start).text("停止");
    } else {
        $(start).text("開始");
    }
}

$(start).on("click", () => {
    if (is_on) {
        animation_end();
    } else {
        animation_start();
    }
});

$(next).on("click", () => {
    animation_end();
    step++;
    show();
});

$(top_button).on("click", () => {
    animation_end();
    animation_reset();
    step = 0;
    show();
});

$("[id^=f-]").on("click", function () {
    fps = $(this).data("fps");
});

function show() {
    Promise.all([get_json(step), get_svg(step)])
        .then((result) => {
            const json = result[0];
            const svg = result[1];
            show_svg(svg);
            show_info(json);
        });
}

function get_json(step) {
    return fetch('get.php?t=j&n=' + step)
        .then((response) => response.json())
        .catch((error) => console.error(error))
}

function get_svg(step) {
    return fetch("get.php?t=s&n=" + step)
        .then((response) => response.text())
}

function show_info(json) {
    const circles = $("[id^=c-]");
    const paths = $("[id^=p-]");

    if (selected_id !== null) {
        show_cell(json, selected_id);
    }

    circles.on("click", function () {
        const id = $(this).attr("id").slice(2);
        show_cell(json, id);
        selected_id = id;
    });
    paths.on("click", function () {
        const id = $(this).attr("id").slice(2);
        const path = json.Paths[id];

        $(width).text(path.width);
    });
}

function show_cell(json, id) {
    const cell = json.Cells[id];
    if (cell === undefined) {
        $(cell_id).text("消滅")
    } else {
        $(cell_id).text(cell.id);
        $(fearness).text((Math.round(cell.persona.fear * 100) / 100).toString());
        $(kindness).text((Math.round(cell.persona.kindness * 100) / 100).toString());
        $(resource).text(cell.resource.toString());
    }
}

function show_svg(svg) {
    $(canvas).html(svg);
}

let frame = 0;
let anime;

function animation_start() {

    frame++;
    if (frame % fps === 0) {
        if (step > 999) {
            animation_end();
        }
        show(frame / fps);
        step++;
    }
    anime = window.requestAnimationFrame(animation_start);
    is_on = true;
    toggle_button();
}

function animation_end() {
    window.cancelAnimationFrame(anime);
    is_on = false;
    toggle_button();
}

function animation_reset() {
    frame = 0;
    is_on = false;
    toggle_button();
}