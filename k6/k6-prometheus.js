import { textSummary } from "https://jslib.k6.io/k6-summary/0.0.1/index.js";
import { htmlReport } from "https://raw.githubusercontent.com/benc-uk/k6-reporter/2.3.0/dist/bundle.js";
import { check, group, sleep } from 'k6';
import http from 'k6/http';
import { Rate, Trend } from 'k6/metrics';

const errorRate = new Rate('errors');
const apiResponseTime = new Trend('api_response_time');

const BASE_URL = 'http://localhost:8080/api/v1';

export const options = {
  stages: [
    { duration: '30s', target: 10 },
    { duration: '1m', target: 10 },
    { duration: '10s', target: 0 },
  ],
  summaryTrendStats: ['avg', 'min', 'med', 'max', 'p(90)', 'p(95)', 'p(99)'],
  thresholds: {
    http_req_duration: ['p(95)<500'],
    http_req_failed: ['rate<0.1'],
    errors: ['rate<0.1'],
  },
};

function uuidv4() {
  return 'xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx'.replace(/[xy]/g, c => {
    var r = Math.random() * 16 | 0,
      v = c === 'x' ? r : (r & 0x3 | 0x8);
    return v.toString(16);
  });
}

function getHeaders() {
  return {
    'Content-Type': 'application/json',
  };
}

const testReciters = [
  { style: 'Style 1' },
  { style: 'Style 2' },
];

export default function () {
  group('API Health Checks', () => {
    healthChecks();
  });

  group('Quran Data Retrieval', () => {
    quranDataTests();
  });

  group('Reciter Endpoint Tests', () => {
    reciterTests();
  });

  sleep(Math.random() * 2 + 1);
}

function healthChecks() {
  const responses = http.batch([
    ['GET', `${BASE_URL}/health`, null, { tags: { name: 'get_health_detailed', expected_status: '503' }, expected_statuses: [200, 503] }],
    ['GET', `${BASE_URL}/health/simple`, null, { tags: { name: 'get_health_simple', expected_status: '200' } }]
  ]);

  if (![200, 503].includes(responses[0].status)) {
    console.log(`[FAILED] GET /health failed. Status: ${responses[0].status}, Body: ${responses[0].body}, Request: ${responses[0].request.url}`);
  }
  if (responses[1].status !== 200) {
    console.log(`[FAILED] GET /health/simple failed. Status: ${responses[1].status}, Body: ${responses[1].body}, Request: ${responses[1].request.url}`);
  }

  const healthCheck = check(responses[0], {
    'GET /health responds correctly (200 or 503)': (r) => [200, 503].includes(r.status),
    'GET /health has data': (r) => r.body && JSON.parse(r.body).data !== null,
  });

  const simpleHealthCheck = check(responses[1], {
    'GET /health/simple status is 200': (r) => r.status === 200,
  });

  errorRate.add(!healthCheck);
  apiResponseTime.add(responses[0].timings.duration);
  errorRate.add(!simpleHealthCheck);
  apiResponseTime.add(responses[1].timings.duration);
}

function quranDataTests() {
  const juzUrl = `${BASE_URL}/juz`;
  let response = http.get(juzUrl, {
    headers: getHeaders(),
    tags: { name: 'get_all_juz', expected_status: '200' }
  });

  if (response.status !== 200) {
    console.log(`[FAILED] GET ${juzUrl} failed with status: ${response.status}. Body: ${response.body}`);
  }

  const juzCheck = check(response, {
    'GET /juz status is 200': (r) => r.status === 200,
    'GET /juz has data': (r) => r.body && JSON.parse(r.body).data !== null,
  });
  errorRate.add(!juzCheck);
  apiResponseTime.add(response.timings.duration);

  const surahsUrl = `${BASE_URL}/surahs`;
  response = http.get(surahsUrl, {
    headers: getHeaders(),
    tags: { name: 'get_all_surahs', expected_status: '200' }
  });

  if (response.status !== 200) {
    console.log(`[FAILED] GET ${surahsUrl} failed with status: ${response.status}. Body: ${response.body}`);
  }

  const surahCheck = check(response, {
    'GET /surahs status is 200': (r) => r.status === 200,
    'GET /surahs has data': (r) => r.body && JSON.parse(r.body).data !== null,
  });
  errorRate.add(!surahCheck);
  apiResponseTime.add(response.timings.duration);

  const surahId = Math.floor(Math.random() * 114) + 1;
  const ayahId = Math.floor(Math.random() * 7) + 1;
  const ayahUrl = `${BASE_URL}/ayahs/${surahId}/${ayahId}`;
  response = http.get(ayahUrl, {
    headers: getHeaders(),
    tags: { name: 'get_ayah_by_id', expected_status: '200/404' },
    expected_statuses: [200, 404]
  });

  if (![200, 404].includes(response.status)) {
    console.log(`[FAILED] GET ${ayahUrl} failed with status: ${response.status}. Body: ${response.body}`);
  }

  const ayahCheck = check(response, {
    'GET /ayahs/{surahId}/{ayahId} responds with 200 or 404': (r) => [200, 404].includes(r.status),
  });
  errorRate.add(!ayahCheck);
  apiResponseTime.add(response.timings.duration);
}

function reciterTests() {
  const recitersUrl = `${BASE_URL}/reciters`;
  let response = http.get(recitersUrl, {
    headers: getHeaders(),
    tags: { name: 'get_all_reciters', expected_status: '200' }
  });

  if (response.status !== 200) {
    console.log(`[FAILED] GET ${recitersUrl} failed with status: ${response.status}. Body: ${response.body}`);
  }

  const getRecitersCheck = check(response, {
    'GET /reciters status is 200': (r) => r.status === 200,
    'GET /reciters has data': (r) => r.body && JSON.parse(r.body).data !== null,
  });
  errorRate.add(!getRecitersCheck);
  apiResponseTime.add(response.timings.duration);

  if (Math.random() < 0.1) {
    const reciterData = JSON.parse(JSON.stringify(testReciters[Math.floor(Math.random() * testReciters.length)]));
    reciterData.id = `test-reciter-${uuidv4()}`;
    reciterData.name = `Test Reciter ${Date.now()}`;

    const reciterPostUrl = `${BASE_URL}/reciters`;
    response = http.post(reciterPostUrl, JSON.stringify(reciterData), {
      headers: getHeaders(),
      tags: { name: 'create_reciter', expected_status: '201' }
    });

    if (response.status !== 201) {
      console.log(`[FAILED] POST ${reciterPostUrl} failed. Status: ${response.status}, Request Body: ${JSON.stringify(reciterData)}, Response Body: ${response.body}`);
    }

    const createCheck = check(response, {
      'POST /reciters status is 201': (r) => r.status === 201,
      'POST /reciters returns created data': (r) => r.body && JSON.parse(r.body).data !== null,
    });
    errorRate.add(!createCheck);
    apiResponseTime.add(response.timings.duration);

    if (response.status === 201) {
      const createdReciter = JSON.parse(response.body).data;
      
      const getReciterUrl = `${BASE_URL}/reciters/${createdReciter.id}`;
      response = http.get(getReciterUrl, {
        headers: getHeaders(),
        tags: { name: 'get_reciter_by_id', expected_status: '200' }
      });

      if (response.status !== 200) {
        console.log(`[FAILED] GET ${getReciterUrl} failed with status: ${response.status}. Response body: ${response.body}`);
      }

      const getByIdCheck = check(response, {
        'GET /reciters/{id} status is 200': (r) => r.status === 200,
        'GET /reciters/{id} returns correct reciter': (r) => r.body && JSON.parse(r.body).data.id === createdReciter.id,
      });
      errorRate.add(!getByIdCheck);
      apiResponseTime.add(response.timings.duration);
    }
  }
}

export function setup() {
  console.log('Starting performance test for Al-Quran API');
  console.log(`Base URL: ${BASE_URL}`);
  
  const response = http.get(`${BASE_URL}/health/simple`, { tags: { name: 'setup_check', expected_status: '200' } });
  
  if (response.status !== 200) {
    console.error(`Setup failed: Unable to connect to API. Status: ${response.status}`);
    console.error(`Response: ${response.body}`);
    return false;
  }
  
  console.log('API connection successful. Starting tests...');
  return { baseUrl: BASE_URL };
}

export function teardown(data) {
  if (data.baseUrl) {
    console.log('Performance test completed');
    console.log(`Base URL tested: ${data.baseUrl}`);
  } else {
    console.log('Performance test failed during setup.');
  }
}

export function handleSummary(data) {
    return {
        "results/result.html": htmlReport(data),
        stdout: textSummary(data, { indent: " ", enableColors: true }),
    };
}