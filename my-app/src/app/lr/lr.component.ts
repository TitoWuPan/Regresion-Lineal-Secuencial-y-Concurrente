import { Component, OnInit } from '@angular/core';
import { TripService } from '../service/api.service';
import { interval, Subscription } from 'rxjs';

@Component({
  selector: 'app-lr',
  templateUrl: './lr.component.html',
  styleUrl: './lr.component.scss'
})

export class LRComponent implements OnInit  {
  model: any = {};
  trips: any[] = [];

  xValue: number = 0;
  calculatedY: any = 0;

  private updateSubscription: Subscription = new Subscription();;

  constructor(private TripService: TripService) {}

  ngOnInit(): void {
    this.Update();
    this.updateSubscription = interval(60000).subscribe(() => {
      this.Update();
    });
  }

  ngOnDestroy(): void {
    if (this.updateSubscription) {
      this.updateSubscription.unsubscribe();
    }
  }

  calculateY(): void {
    if (this.model.beta1 !== undefined && this.model.beta0 !== undefined) {
      this.calculatedY = this.model.beta1 * this.xValue + this.model.beta0;
    }
  }

  Update() {
    this.TripService.getModel().subscribe(data => {
      this.model = data;
      console.log(this.model)
    });
    this.TripService.getTrips().subscribe(data => {
      this.trips = data;
      console.log(this.trips)
    });
  }
}
