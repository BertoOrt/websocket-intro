$(() => {
	const ws = new WebSocket("ws://localhost:3000/reset");
	const message = JSON.stringify({reset: true})

	ws.onmessage = msg => {
		const data = JSON.parse(msg.data);
		console.log("reset", data.reset);
		if (data.reset) {
			// resetTimer()
		}
	};

	const $button = $("button");
	const clockSeconds = 600;
	const resetTimer = () => {
		ws.send(message);
		clock.setTime(clockSeconds, function () {
			clock.start();
		});
	};

	// add reset function
	$button.on("click", resetTimer)

	// create clock
	const clock = new FlipClock($('.your-clock'), {
		countdown: true,
		stop: () => {
			$('.over').css('display', 'inline-block');
			this.stop();
			$button.off();
		}
	});

	// start timer
	clock.setTime(clockSeconds);
	clock.start();
})
