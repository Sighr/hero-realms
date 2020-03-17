import BaseComponent from "../BaseComponent.js";
import Field from "../Field/Field.js";

/**
 * class for data that is shared between players:
 * gems, market, market deck, sacrifice pull
 */
export default class Board extends BaseComponent {
    /**
     *
     * @param playerNum - number of players
     */
    constructor(playerNum) {
        super();
        this.root.classList.add('board-root');

        const img = document.createElement('img');
        this.root.appendChild(img);
        img.src = `/public/img/board.jpg`;
        img.classList.add('board-img');
        img.classList.add('main-back_img');

        this.fields = Array(playerNum);
        for(let i = 0; i < playerNum; ++i) {
            this.fields[i] = new Field();
            this.fields[i].render(this.root).classList.add('board-field');
        }
    }
}