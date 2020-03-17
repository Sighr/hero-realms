

export default class BaseComponent {
    constructor() {
        this.root = document.createElement('div');
    }

    render(parent) {
        if(!parent) {
            return this.root;
        }
        parent.appendChild(this.root);
        return this.root;
    }
}