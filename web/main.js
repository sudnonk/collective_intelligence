const prev = $("#prev");
const start = $("#start");
const next = $("#next");
const top_button = $("#top");

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

    circles.on("click", () => {
        console.log(this, $(this));
        const id = $(this).attr("id").slice(2);
        console.log(id);
        console.log(json.Cells);
        const cell = json.Cells[id];
        console.log(cell);
    });
    paths.on("click", () => {
        const id = $(this).attr("id").slice(2);
        const path = json.Paths[id];
        console.log(path);
    })
}

function show_svg(svg) {
    $(canvas).html(svg);
}