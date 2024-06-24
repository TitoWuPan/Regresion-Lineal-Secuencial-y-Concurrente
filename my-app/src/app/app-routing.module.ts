import { NgModule, Component } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LRComponent} from './lr/lr.component';

const routes: Routes = [
  {path: '', component: LRComponent},
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
