import { Component, OnInit } from '@angular/core';
import { BaseChartDirective } from 'ng2-charts';
import { ApiService } from '../../services/api.service';
import { Chart, registerables } from 'chart.js';

Chart.register(...registerables);

@Component({
  selector: 'app-dashboard',
  standalone: true,
  imports: [BaseChartDirective],
  templateUrl: './dashboard.html',
  styleUrl: './dashboard.scss'
})
export class Dashboard implements OnInit {

  stats: any = {};

  lineChartData: any = {
    labels: [],
    datasets: [
      {
        data: [],
        label: 'Response Time (ms)',
        fill: false,
      }
    ]
  };

  constructor() {}

  async ngOnInit() {
    const res = await ApiService.getDashboard();
    this.stats = res.data;

    // assume backend sends last 10 response times
    this.lineChartData.labels = this.stats.timestamps;
    this.lineChartData.datasets[0].data = this.stats.responseTimes;
  }
}
