import BaseComponent from "../BaseComponent.js";

/**
 * class for data that is unique to specific player
 */
export default class Field extends BaseComponent {
    constructor() {
        // here should be drawing of all player's field:
        // deck, discard, hand, field where cards are placed after they are played, etc.
        super();
        this.root.classList.add('field-root');
        const img = document.createElement('img');
        this.root.appendChild(img);
        img.src = `/static/img/field.jpg`;
        img.classList.add('field-img');
        img.classList.add('main-back_img');
        // should implement hierarchy elements
        // this.deck = new Deck();
        // this.discard = new Discard();
        // this.hand = new Hand();
    }
}