import http from 'k6/http';
import { check, sleep } from 'k6';
import { Counter } from 'k6/metrics';

// Define custom counters
export const status200 = new Counter('http_200');
export const status429 = new Counter('http_429');
export const statusOther = new Counter('http_other');

export const options = {
    vus: 150,
    duration: '30s',
};

export default function () {
    const res = http.get('http://localhost:8080/products');

    if (res.status === 200) {
        status200.add(1);
    } else if (res.status === 429) {
        status429.add(1);
    } else {
        statusOther.add(1);
    }

    check(res, {
        'status is 200 or 429': (r) => r.status === 200 || r.status === 429,
    });

    sleep(0.1);
}
