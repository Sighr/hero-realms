import Board from "../Components/Board/Board.js";

const game = document.getElementById('game');

const card = new Board(2);
card.render(game);
