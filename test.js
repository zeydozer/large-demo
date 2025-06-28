import http from 'k6/http';

export let options = {
  stages: [
    { duration: '30s', target: 500 },
    { duration: '30s', target: 1000 },
    { duration: '30s', target: 1500 },
    { duration: '30s', target: 0 },
  ]
};

export default function () {
  const userId = Math.floor(Math.random() * 1000000) + 1;
  http.get(`http://localhost:8080/users/${userId}`);
}
