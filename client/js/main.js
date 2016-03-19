'use strict'

$(() => {
	const ws = new WebSocket('ws://localhost:3000/reset');
	const $button = $('button');
	const clockSeconds = 1200;

	// create clock
	const clock = new FlipClock($('.your-clock'), {
		countdown: true,
		stop: () => {
			// end game and send message that it's over
			endGame();
			let message = JSON.stringify({gameOver: true, time: 0})
			ws.send(message);
		}
	});

	// displays message and disables button when the game is over
	const endGame = () => {
		$button.off();
		$button.text('Game Over')
		$('.over').css('display', 'inline-block');
		this.stop();
	}

	// get time and start it
	ws.onmessage = msg => {
		let data = JSON.parse(msg.data);
		if (data.gameOver) {
			endGame();
		} else {
			let startTime = new Date(data.time).getTime();
			let currentTime = new Date().getTime();
			let timeLeft = clockSeconds - Math.floor((currentTime - startTime) / 1000)
			clock.setTime(timeLeft);
			clock.start();
		}
	};

	// send reset message
	const resetTimer = () => {
		let message = JSON.stringify({gameOver: false, time: clock.getTime().time})
		ws.send(message);
		clock.setTime(clockSeconds, () => {
			clock.start();
		});
	};

	// add eventListener
	$button.on('click', resetTimer)
})
