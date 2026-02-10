import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { ApiService } from '../../services/api.service';

@Component({
  selector: 'app-api-details',
  standalone: true,
  templateUrl: './api-details.html',
  styleUrl: './api-details.scss'
})
export class ApiDetails implements OnInit {

  apiId!: string;
  logs: any[] = [];

  constructor(private route: ActivatedRoute) {}

  async ngOnInit() {
    this.apiId = this.route.snapshot.paramMap.get('id')!;
    await this.loadLogs();
  }

  async loadLogs() {
    const res = await ApiService.getLogs(); // backend can filter by apiId later
    this.logs = res.data.filter((l: any) => l.apiId === this.apiId);
  }
}
