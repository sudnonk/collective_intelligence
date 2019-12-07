const prev = document.getElementById("prev");
const start = document.getElementById("start");
const next = document.getElementById("next");
const top_button = document.getElementById("top");

const info_step = document.getElementById("step");
const info_cells = document.getElementById("cells");

const canvas = document.getElementById("canvas");

let step = 0;

prev.onclick = (ev) => {
    step--;
    if (step < 0) {
        step = 0;
    }
    show();
};

next.onclick = (ev) => {
    step++;
    show();
};

top_button.onclick = (ev) => {
    step = 0;
    show();
};

function show() {
    get_json(step)
        .then((json) => {
            console.log(json);
            show_info(json);
            visualize(json);
        })
}

async function get_json(step) {
    return fetch('get.php?n=' + step)
        .then((response) => response.json())
        .catch((error) => console.error(error))
}

/**
 * @param json Object
 */
function show_info(json) {
    const num = json.Cells.length;
    console.log(num);

    info_step.innerText = step;
    info_cells.innerText = num;
}

function visualize(json) {

}