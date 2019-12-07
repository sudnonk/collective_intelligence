const prev = document.getElementById("prev");
const start = document.getElementById("start");
const next = document.getElementById("next");
const top_button = document.getElementById("top");

let step = 0;

prev.onclick = (ev) => {
    step--;
    const json = get_json(step);
};

next.onclick = (ev) => {
    step++;
    const json = get_json(step);
};

top_button.onclick = (ev) => {
    step = 0;
    const json = get_json(step);
};

function get_json(step) {
    fetch('get.php?n=' + step)
        .then((response) => response.json())
        .then((json) => console.log(json))
        .catch((error) => console.error(error))
}