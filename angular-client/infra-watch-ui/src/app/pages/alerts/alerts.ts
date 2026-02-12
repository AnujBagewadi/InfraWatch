import { Component, OnInit } from '@angular/core';
import { ApiService } from '../../services/api.service';
import { AlertModel } from '../../models/alert.model';

@Component({
  selector: 'app-alerts',
  standalone: true,
  templateUrl: './alerts.html',
  styleUrl: './alerts.scss'
})
export class Alerts implements OnInit {

  alerts: AlertModel[] = [];

  async ngOnInit() {
    const res = await ApiService.getAlerts();
    this.alerts = res.data;
  }
}
