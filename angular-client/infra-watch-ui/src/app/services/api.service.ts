import axios from 'axios';
import { ApiModel } from '../models/api.model';
import { LogModel } from '../models/log.model';
import { AlertModel } from '../models/alert.model';

const BASE_URL = 'http://localhost:8080';

export const ApiService = {
  getDashboard: () => axios.get(`${BASE_URL}/dashboard`),

  getApis: () => axios.get<ApiModel[]>(`${BASE_URL}/apis`),

  addApi: (data: ApiModel) => axios.post(`${BASE_URL}/apis`, data),

  getLogs: () => axios.get<LogModel[]>(`${BASE_URL}/logs`),

  getAlerts: () => axios.get<AlertModel[]>(`${BASE_URL}/alerts`),

  getApiDetails: (id: string) => axios.get(`${BASE_URL}/apis/${id}`),
};
