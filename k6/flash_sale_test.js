import http from 'k6/http';
import { check, sleep } from 'k6';

export const options = {
	stages: [
				{ duration: '5s', target: 50 },
				{ duration: '10s', target: 100 },
				{ duration: '5s', target: 0 },
	],
	thresholds: {
			http_req_duration: ['p(95)<500'],
			http_req_failed: ['rate<0.01'],
	},
};

export default function() {
	
	const userID = Math.floor(Math.random() * 1000);
	const itemID = 'X345';

	const res = http.get(`http://order-service:8700/place-order?user_id=u${userID}&item_id=${itemID}`)

	check(res, {
			'status is 200 or 400': (r) => r.status === 200 || r.status === 409,
	});

	sleep(0.1)

}
