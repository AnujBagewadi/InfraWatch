import axios from 'axios';

const BASE_URL = 'http://localhost:8080';

export const ApiService = {
  getDashboard: () => axios.get(`${BASE_URL}/dashboard`),
  getApis: () => axios.get(`${BASE_URL}/apis`),
  addApi: (data: any) => axios.post(`${BASE_URL}/apis`, data),
  getLogs: () => axios.get(`${BASE_URL}/logs`),
  getAlerts: () => axios.get(`${BASE_URL}/alerts`),
  getApiDetails: (id: string) => axios.get(`${BASE_URL}/apis/${id}`),
};
