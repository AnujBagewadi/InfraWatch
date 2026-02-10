import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ApiDetails } from './api-details';

describe('ApiDetails', () => {
  let component: ApiDetails;
  let fixture: ComponentFixture<ApiDetails>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [ApiDetails]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ApiDetails);
    component = fixture.componentInstance;
    await fixture.whenStable();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
