import BaseComponent from "../BaseComponent.js";
import Field from "../Field/Field.js";

export default class Board extends BaseComponent {
    constructor(playerNum) {
        super();
        this.root.classList.add('board-root');

        const img = document.createElement('img');
        this.root.appendChild(img);
        img.src = `/static/img/board.jpg`;
        img.classList.add('board-img');
        img.classList.add('main-back_img');

        this.fields = Array(playerNum);
        for(let i = 0; i < playerNum; ++i) {
            this.fields[i] = new Field();
            this.fields[i].render(this.root).classList.add('board-field');
        }
    }
}