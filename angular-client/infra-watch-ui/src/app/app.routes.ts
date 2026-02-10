import { Routes } from '@angular/router';
import { Dashboard } from './pages/dashboard/dashboard';
import { Apis } from './pages/apis/apis';
import { Logs } from './pages/logs/logs';
import { Alerts } from './pages/alerts/alerts';
import { ApiDetails } from './pages/api-details/api-details';

export const routes: Routes = [
  { path: '', component: Dashboard },
  { path: 'apis', component: Apis },
  { path: 'logs', component: Logs },
  { path: 'alerts', component: Alerts },
  { path: 'apis/:id', component: ApiDetails},
];
