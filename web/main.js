const prev = $("prev");
const start = $("start");
const next = $("next");
const top_button = $("top");

const info_step = $("step");
const info_cells = $("cells");

const canvas = $("canvas");

let step = 0;

prev.on("click", () => {
    console.log($(this).innerText);
    step--;
    if (step < 0) {
        step = 0;
    }
    show();
});

next.on("click", () => {
    step++;
    show();
});

top_button.on("click", () => {
    step = 0;
    show();
});

function show() {
    Promise.all([get_json(step), get_svg(step)])
        .then((json, svg) => {
            show_info(json);
            show_svg(svg);
        });
}

async function get_json(step) {
    return fetch('get.php?t=j&n=' + step)
        .then((response) => response.json())
        .catch((error) => console.error(error))
}

async function get_svg(step) {
    return fetch("svg.php?t=s&n=" + step)
        .then((response) => response.text())
}

function show_info(json) {
    const num = Object.keys(json.Cells).length;

    info_step.innerText = step.toString();
    info_cells.innerText = num.toString();
}

function show_svg(svg) {
    canvas.innerHTML = svg;
}