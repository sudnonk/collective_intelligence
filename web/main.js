const prev = $("#prev");
const start = $("#start");
const next = $("#next");
const top_button = $("#top");

const info_step = $("#step");
const info_cells = $("#cells");

const canvas = $("#canvas");

let step = 0;

$(prev).on("click", () => {
    step--;
    if (step < 0) {
        step = 0;
    }
    show();
});

$(next).on("click", () => {
    step++;
    show();
});

$(top_button).on("click", () => {
    step = 0;
    show();
});

function show() {
    Promise.all([get_json(step), get_svg(step)])
        .then((result) => {
            console.log(result);
            const json = result[0];
            const svg = result[1];
            show_info(json);
            show_svg(svg);
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
    const num = Object.keys(json.Cells).length;

    $(info_step).text(step.toString());
    $(info_cells).text(num.toString());
}

function show_svg(svg) {
    console.log(svg);
    $(canvas).html(svg);
}