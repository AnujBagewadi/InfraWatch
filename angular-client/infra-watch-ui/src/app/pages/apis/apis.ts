import { Component, OnInit } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { RouterLink } from '@angular/router';
import { ApiService } from '../../services/api.service';

@Component({
  selector: 'app-apis',
  standalone: true,
  imports: [FormsModule, RouterLink],
  templateUrl: './apis.html',
  styleUrl: './apis.scss'
})
export class Apis implements OnInit {

  apis: any[] = [];

  newApi = {
    name: '',
    url: '',
    method: 'GET',
    threshold: 1000
  };

  async ngOnInit() {
    await this.loadApis();
  }

  async loadApis() {
    const res = await ApiService.getApis();
    this.apis = res.data;
  }

  async addApi() {
    await ApiService.addApi(this.newApi);
    this.newApi = { name: '', url: '', method: 'GET', threshold: 1000 };
    await this.loadApis();
  }
}
