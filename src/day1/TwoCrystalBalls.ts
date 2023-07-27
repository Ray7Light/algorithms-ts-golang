export default function two_crystal_balls(breaks: boolean[]): number {
	const jmpAmount = Math.floor(Math.sqrt(breaks.length));

	let i = jmpAmount;

	for (; i < breaks.length; i += jmpAmount) {
		if (breaks[i]) {
			break;
		}
	}

	if (i > breaks.length) {
		return -1;
	}

	i -= jmpAmount;

	for (; i < breaks.length; i++) {
		if (breaks[i]) {
			break;
		}
	}

	return i;
}
