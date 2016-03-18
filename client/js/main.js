$(() => {
	const ws = new WebSocket("ws://localhost:3000/reset");
	const message = JSON.stringify({reset: true})

	ws.onmessage = msg => {
		const data = JSON.parse(msg.data);
		if (data.reset) {
			clock.setTime(clockSeconds, function () {
				clock.start();
			});
		}
	};

	const $button = $("button");
	const clockSeconds = 600;
	const resetTimer = () => {
		console.log(clock.getTime().time);
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
