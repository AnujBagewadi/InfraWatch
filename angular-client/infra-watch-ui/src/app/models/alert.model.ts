export interface AlertModel {
  id?: string;          // Mongo _id (optional for UI)
  apiId: string;
  message: string;
  responseTime: number;
  timestamp: string;
}
