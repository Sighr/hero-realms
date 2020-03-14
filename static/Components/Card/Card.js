import BaseComponent from "../BaseComponent.js";

export default class Card extends BaseComponent {
    constructor(cardName) {
        super();
        this.root.classList.add('card-root');
        const img = document.createElement('img');
        this.root.appendChild(img);
        img.src = `/static/img/${cardName}`;
        img.classList.add('card-img');
    }
}