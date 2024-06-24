import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

interface Model {
  beta0: number;
  beta1: number;
}

interface TripData {
  distance: number;
  speed: number;
  duration: number;
}

@Injectable({
  providedIn: 'root'
})

export class TripService {
  private apiUrl = 'http://localhost:8080';

  constructor(private http: HttpClient) {}

  getModel(): Observable<Model> {
    return this.http.get<Model>(`${this.apiUrl}/model`);
  }

  getTrips(): Observable<TripData[]> {
    return this.http.get<TripData[]>(`${this.apiUrl}/trips`);
  }
}
