import { Component, OnInit } from '@angular/core';
import { ApiService } from '../../services/api.service';
import { LogModel } from '../../models/log.model';

@Component({
  selector: 'app-logs',
  standalone: true,
  templateUrl: './logs.html',
  styleUrl: './logs.scss'
})
export class Logs implements OnInit {

  logs: LogModel[] = [];

  async ngOnInit() {
    const res = await ApiService.getLogs();
    this.logs = res.data;
  }
}
