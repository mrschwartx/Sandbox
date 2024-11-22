import http from 'k6/http';
import { check, sleep } from 'k6';

// Define the base URL of your API
const BASE_URL = 'http://localhost:8080';

// Define the test scenario options
export let options = {
  stages: [
    { duration: '30s', target: 20 },  // Ramp up to 20 users in 30 seconds
    { duration: '1m', target: 20 },   // Stay at 20 users for 1 minute
    { duration: '30s', target: 0 },   // Ramp down to 0 users in 30 seconds
  ],
};

// Default function which is called when running the test
export default function() {
  // Test all the endpoints sequentially or you can randomize them
  getAllTasks();
  createTask();
  getTaskByID();
  updateTask();
  deleteTask();
}

// Test GET /tasks (Fetch all tasks)
export function getAllTasks() {
  let response = http.get(`${BASE_URL}/tasks`);
  check(response, {
    'is status 200': (r) => r.status === 200,
    'is response time < 500ms': (r) => r.timings.duration < 500,
  });
  sleep(1);
}

// Test POST /tasks/create (Create a new task)
export function createTask() {
  let payload = JSON.stringify({
    title: 'Load Test Task',
    status: 'pending',
  });

  let params = {
    headers: {
      'Content-Type': 'application/json',
    },
  };

  let response = http.post(`${BASE_URL}/tasks/create`, payload, params);
  check(response, {
    'is status 200': (r) => r.status === 200,
    'is response time < 500ms': (r) => r.timings.duration < 500,
  });
  sleep(1);
}

// Test GET /tasks/get (Get task by ID)
export function getTaskByID() {
  let response = http.get(`${BASE_URL}/tasks/get?id=1`);  // Adjust the ID as necessary
  check(response, {
    'is status 200': (r) => r.status === 200,
    'is response time < 500ms': (r) => r.timings.duration < 500,
  });
  sleep(1);
}

// Test PUT /tasks/update (Update a task by ID)
export function updateTask() {
  let payload = JSON.stringify({
    title: 'Updated Load Test Task',
    status: 'completed',
  });

  let params = {
    headers: {
      'Content-Type': 'application/json',
    },
  };

  let response = http.put(`${BASE_URL}/tasks/update?id=1`, payload, params);  // Adjust ID as necessary
  check(response, {
    'is status 200': (r) => r.status === 200,
    'is response time < 500ms': (r) => r.timings.duration < 500,
  });
  sleep(1);
}

// Test DELETE /tasks/delete (Delete a task by ID)
export function deleteTask() {
  let response = http.del(`${BASE_URL}/tasks/delete?id=1`);  // Adjust ID as necessary
  check(response, {
    'is status 200': (r) => r.status === 200,
    'is response time < 500ms': (r) => r.timings.duration < 500,
  });
  sleep(1);
}
